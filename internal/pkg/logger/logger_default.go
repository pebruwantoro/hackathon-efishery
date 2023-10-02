package logger

import (
	"context"
	"encoding/json"
	"os"

	"github.com/gogo/protobuf/proto"
	"github.com/spf13/cast"
	"go.elastic.co/apm/module/apmzap/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type defaultLogger struct {
	zapLogger *zap.Logger
	level     Level
}

func NewLogger(opt Option) Logger {
	logger := &defaultLogger{}

	if opt.IsEnable {
		logger.zapLogger = NewZapLogger(logger.level, zapcore.AddSync(os.Stdout))
	} else {
		logger.zapLogger = NewZapLogger(logger.level, nil)
	}

	Log = logger

	return logger
}

func (d *defaultLogger) Debug(ctx context.Context, message string, details ...interface{}) {
	zapLogs := []zap.Field{
		zap.String("level", "debug"),
	}

	traceContextFields := apmzap.TraceContext(ctx)

	fields := formatToField(details...)
	zapLogs = append(zapLogs, formatLogs(ctx, message, fields...)...)
	d.zapLogger.With(traceContextFields...).Debug(message, zapLogs...)
}

func (d *defaultLogger) Info(ctx context.Context, message string, details ...interface{}) {
	zapLogs := []zap.Field{
		zap.String("level", "info"),
	}

	traceContextFields := apmzap.TraceContext(ctx)

	fields := formatToField(details...)
	zapLogs = append(zapLogs, formatLogs(ctx, message, fields...)...)
	d.zapLogger.With(traceContextFields...).Info(message, zapLogs...)
}

func (d *defaultLogger) Warn(ctx context.Context, message string, details ...interface{}) {
	zapLogs := []zap.Field{
		zap.String("level", "warn"),
	}

	traceContextFields := apmzap.TraceContext(ctx)

	fields := formatToField(details...)
	zapLogs = append(zapLogs, formatLogs(ctx, message, fields...)...)
	d.zapLogger.With(traceContextFields...).Warn(message, zapLogs...)
}

func (d *defaultLogger) Error(ctx context.Context, message string, details ...interface{}) {
	zapLogs := []zap.Field{
		zap.String("level", "error"),
	}

	traceContextFields := apmzap.TraceContext(ctx)

	fields := formatToField(details...)
	zapLogs = append(zapLogs, formatLogs(ctx, message, fields...)...)
	d.zapLogger.With(traceContextFields...).Error(message, zapLogs...)
}

func (d *defaultLogger) Fatal(ctx context.Context, message string, details ...interface{}) {
	zapLogs := []zap.Field{
		zap.String("level", "fatal"),
	}

	traceContextFields := apmzap.TraceContext(ctx)

	fields := formatToField(details...)
	zapLogs = append(zapLogs, formatLogs(ctx, message, fields...)...)
	d.zapLogger.With(traceContextFields...).Fatal(message, zapLogs...)
}

func (d *defaultLogger) Panic(ctx context.Context, message string, details ...interface{}) {
	zapLogs := []zap.Field{
		zap.String("level", "panic"),
	}

	traceContextFields := apmzap.TraceContext(ctx)

	fields := formatToField(details...)
	zapLogs = append(zapLogs, formatLogs(ctx, message, fields...)...)
	d.zapLogger.With(traceContextFields...).Panic(message, zapLogs...)
}

func formatLogs(ctx context.Context, msg string, fields ...Field) (logRecord []zap.Field) {
	ctxVal := ExtractCtx(ctx)

	// add global value from context that must be exist on all logs!
	logRecord = append(logRecord, zap.String("message", msg))

	logRecord = append(logRecord, zap.String("_app_name", ctxVal.ServiceName))
	logRecord = append(logRecord, zap.String("_app_version", ctxVal.ServiceVersion))
	logRecord = append(logRecord, zap.Int("_app_port", ctxVal.ServicePort))
	logRecord = append(logRecord, zap.String("_app_tag", ctxVal.Tag))
	logRecord = append(logRecord, zap.String("_app_method", ctxVal.ReqMethod))
	logRecord = append(logRecord, zap.String("_app_uri", ctxVal.ReqURI))

	// add additional data that available across all log, such as user_id
	if ctxVal.AdditionalData != nil {
		logRecord = append(logRecord, zap.Any("_app_data", ctxVal.AdditionalData))
	}

	for _, field := range fields {
		logRecord = append(logRecord, formatLog(field.Key, field.Val))
	}

	return
}

func formatToField(details ...interface{}) (logRecord []Field) {
	for index, msg := range details {
		logRecord = append(logRecord, Field{
			Key: "_message_" + cast.ToString(index),
			Val: msg,
		},
		)
	}

	return
}

func formatLog(key string, msg interface{}) (logRecord zap.Field) {
	if msg == nil {
		logRecord = zap.Any(key, struct{}{})
		return
	}

	// handle proto message
	p, ok := msg.(proto.Message)
	if ok {
		b, _err := json.Marshal(p)
		if _err != nil {
			logRecord = zap.Any(key, p.String())
			return
		}

		var data interface{}
		if _err = json.Unmarshal(b, &data); _err != nil {
			// string cannot be masked, so only try to marshal as json object
			logRecord = zap.Any(key, p.String())
			return
		}

		// use object json
		logRecord = zap.Any(key, data)
		return
	}

	// handle string, string is cannot be masked, just write it
	// but try to parse as json object if possible
	if str, ok := msg.(string); ok {
		var data interface{}
		if _err := json.Unmarshal([]byte(str), &data); _err != nil {
			logRecord = zap.String(key, str)
			return
		}

		logRecord = zap.Any(key, data)
		return
	}

	// not masked since it failed to convert to reflect.Value above
	logRecord = zap.Any(key, msg)
	return
}
