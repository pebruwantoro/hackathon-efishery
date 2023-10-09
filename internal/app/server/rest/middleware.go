package rest

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.elastic.co/apm/module/apmechov4/v2"

	"github.com/pebruwantoro/hackathon-efishery/internal/app/container"
	"github.com/pebruwantoro/hackathon-efishery/internal/pkg"
	"github.com/pebruwantoro/hackathon-efishery/internal/pkg/helpers"
	"github.com/pebruwantoro/hackathon-efishery/internal/pkg/logger"
	"github.com/pebruwantoro/hackathon-efishery/internal/pkg/response"
	pkgValidator "github.com/pebruwantoro/hackathon-efishery/internal/pkg/validator"
)

func SetupMiddleware(server *echo.Echo, container *container.Container) {
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH, echo.OPTIONS},
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-App-Token, X-Client-Id"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	server.Use(middleware.Recover())
	server.Use(apmechov4.Middleware(apmechov4.WithTracer(container.Tracer)))
	server.Use(SetLoggerMiddleware(container))
	server.Use(LoggerMiddleware(container))

	server.HTTPErrorHandler = errorHandler
	server.Validator = &DataValidator{ValidatorData: pkgValidator.SetupValidator()}
}

func SetLoggerMiddleware(container *container.Container) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if SkipLoggerMiddleware(c.Path()) {
				return next(c)
			}

			cfg := container.Config
			ctxLogger := logger.Context{
				ServiceName:    cfg.App.Name,
				ServiceVersion: cfg.App.Version,
				ServicePort:    cfg.App.HttpPort,
				ReqMethod:      c.Request().Method,
				ReqURI:         c.Request().URL.String(),
			}

			var bodyByte []byte
			if c.Request().Body != nil {
				bodyByte, _ = io.ReadAll(c.Request().Body)
				ctxLogger.ReqBody = string(bodyByte)

				c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyByte))
			}

			request := c.Request()

			ctx := logger.InjectCtx(request.Context(), ctxLogger)
			c.SetRequest(request.WithContext(ctx))

			logger.Log.Info(ctx, "Request Header", c.Request().Header)

			if !logger.IsSkipLog(c.Request().Header.Get("Content-Type")) {
				logger.Log.Info(ctx, "Request Body", string(bodyByte))
			} else {
				logger.Log.Info(ctx, "Request Not Log Because Unsupported Content-Type")
			}

			return next(c)
		}
	}
}

func LoggerMiddleware(container *container.Container) echo.MiddlewareFunc {
	return middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		if SkipLoggerMiddleware(c.Path()) {
			return
		}

		// log request header, body & response
		ctx := c.Request().Context()

		if !logger.IsSkipLog(c.Response().Header().Get("Content-Type")) {
			logger.Log.Info(ctx, "Response Body", string(resBody))
		} else {
			logger.Log.Info(ctx, "Response Not Log Because Unsupported Content-Type")
		}
	})
}

type DataValidator struct {
	ValidatorData *validator.Validate
}

func (cv *DataValidator) Validate(i interface{}) error {
	return cv.ValidatorData.Struct(i)
}

func SkipLoggerMiddleware(path string) bool {
	switch path {
	case "/", "/metrics", "/favicon.ico":
		return true
	}

	return false
}

func AuthMiddleware(container *container.Container) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var (
				token    string
				jwtToken *jwt.Token
				err      error
			)

			cfg := container.Config
			ctx := c.Request().Context()
			reqPath := c.Request().URL.Path

			if SkipLoggerMiddleware(reqPath) {
				return next(c)
			}

			headerAuth := c.Request().Header.Get(pkg.HEADER_AUTHORIZATION)
			if headerAuth != "" {
				splitHeader := strings.Split(headerAuth, " ")
				if len(splitHeader) != 2 {
					logger.Log.Error(ctx, "invalid header authorization")
					return c.JSON(401, response.DefaultResponse{
						Success: false,
						Message: "Unauthorized",
					})
				}

				token = splitHeader[1]
			}
			claimToken := cfg.Token.Secret

			jwtToken, err = jwt.ParseWithClaims(token, &helpers.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
				}

				return []byte(claimToken), nil
			})
			if err != nil {
				logger.Log.Error(ctx, fmt.Sprintf("error parseWithClaim : %s", err.Error()))
				return c.JSON(401, response.DefaultResponse{
					Success: false,
					Message: "Unauthorized",
				})
			}

			if !jwtToken.Valid {
				logger.Log.Error(ctx, "invalid jwt token")
				return c.JSON(401, response.DefaultResponse{
					Success: false,
					Message: "Unauthorized",
				})
			}
			claims, ok := jwtToken.Claims.(*helpers.UserClaims)
			if !ok {
				return c.JSON(401, response.DefaultResponse{
					Success: false,
					Message: "Unauthorized",
				})
			}

			c.Request().Header.Add(pkg.HEADER_USER_EMAIL, claims.Email)
			c.Request().Header.Add(pkg.Header_USER_UUID, claims.UUID)

			return next(c)
		}
	}
}

func AuthAdminMiddleware(container *container.Container) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var (
				token    string
				jwtToken *jwt.Token
				err      error
			)

			cfg := container.Config
			ctx := c.Request().Context()
			reqPath := c.Request().URL.Path

			if SkipLoggerMiddleware(reqPath) {
				return next(c)
			}

			headerAuth := c.Request().Header.Get(pkg.HEADER_AUTHORIZATION)
			if headerAuth != "" {
				splitHeader := strings.Split(headerAuth, " ")
				if len(splitHeader) != 2 {
					logger.Log.Error(ctx, "invalid header authorization")
					return c.JSON(401, response.DefaultResponse{
						Success: false,
						Message: "Unauthorized",
					})
				}

				token = splitHeader[1]
			}
			claimToken := cfg.Token.Secret

			jwtToken, err = jwt.ParseWithClaims(token, &helpers.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
				}

				return []byte(claimToken), nil
			})
			if err != nil {
				logger.Log.Error(ctx, fmt.Sprintf("error parseWithClaim : %s", err.Error()))
				return c.JSON(401, response.DefaultResponse{
					Success: false,
					Message: "Unauthorized",
				})
			}

			if !jwtToken.Valid {
				logger.Log.Error(ctx, "invalid jwt token")
				return c.JSON(401, response.DefaultResponse{
					Success: false,
					Message: "Unauthorized",
				})
			}
			claims, ok := jwtToken.Claims.(*helpers.UserClaims)
			if !ok {
				return c.JSON(401, response.DefaultResponse{
					Success: false,
					Message: "Unauthorized",
				})
			}

			if claims.Role != "ADMIN" {
				return c.JSON(401, response.DefaultResponse{
					Success: false,
					Message: "Unauthorized",
				})
			}

			c.Request().Header.Add(pkg.HEADER_USER_EMAIL, claims.Email)
			c.Request().Header.Add(pkg.Header_USER_UUID, claims.UUID)

			return next(c)
		}
	}
}

func AuthManagerialMiddleware(container *container.Container) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var (
				token    string
				jwtToken *jwt.Token
				err      error
			)

			cfg := container.Config
			ctx := c.Request().Context()
			reqPath := c.Request().URL.Path

			if SkipLoggerMiddleware(reqPath) {
				return next(c)
			}

			headerAuth := c.Request().Header.Get(pkg.HEADER_AUTHORIZATION)
			if headerAuth != "" {
				splitHeader := strings.Split(headerAuth, " ")
				if len(splitHeader) != 2 {
					logger.Log.Error(ctx, "invalid header authorization")
					return c.JSON(401, response.DefaultResponse{
						Success: false,
						Message: "Unauthorized",
					})
				}

				token = splitHeader[1]
			}
			claimToken := cfg.Token.Secret

			jwtToken, err = jwt.ParseWithClaims(token, &helpers.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
				}

				return []byte(claimToken), nil
			})
			if err != nil {
				logger.Log.Error(ctx, fmt.Sprintf("error parseWithClaim : %s", err.Error()))
				return c.JSON(401, response.DefaultResponse{
					Success: false,
					Message: "Unauthorized",
				})
			}

			if !jwtToken.Valid {
				logger.Log.Error(ctx, "invalid jwt token")
				return c.JSON(401, response.DefaultResponse{
					Success: false,
					Message: "Unauthorized",
				})
			}
			claims, ok := jwtToken.Claims.(*helpers.UserClaims)
			if !ok {
				return c.JSON(401, response.DefaultResponse{
					Success: false,
					Message: "Unauthorized",
				})
			}

			if claims.Role == "STAFF" {
				return c.JSON(401, response.DefaultResponse{
					Success: false,
					Message: "Unauthorized",
				})
			}

			c.Request().Header.Add(pkg.HEADER_USER_EMAIL, claims.Email)
			c.Request().Header.Add(pkg.Header_USER_UUID, claims.UUID)

			return next(c)
		}
	}
}
