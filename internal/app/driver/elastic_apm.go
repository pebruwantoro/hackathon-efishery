package driver

import (
	"net/url"

	"github.com/pebruwantoro/hackathon-efishery/config"
	"go.elastic.co/apm/v2"
	"go.elastic.co/apm/v2/transport"
)

func NewElasticAPMTracer(cfg config.ElasticAPMConfig) (tracer *apm.Tracer) {
	u, err := url.Parse(cfg.ServerUrl)
	if err != nil {
		panic(err)
	}

	transport, err := transport.NewHTTPTransport(transport.HTTPTransportOptions{
		ServerURLs: []*url.URL{u},
	})
	if err != nil {
		panic(err)
	}

	opt := apm.TracerOptions{
		ServiceName:        cfg.ServiceName,
		ServiceEnvironment: cfg.ServiceEnv,
		Transport:          transport,
	}

	tracer, err = apm.NewTracerOptions(opt)
	if err != nil {
		panic(err)
	}

	return
}
