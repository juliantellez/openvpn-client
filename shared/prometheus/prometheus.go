package prometheus

import (
	"context"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"github.com/sirupsen/logrus"
)

const (
	handlerPatternMetrics = "/__/metrics"
)

type (
	// Metrics struct for prometheus
	Metrics struct {
		server *http.Server
	}
)

// New configures a handler for prometheus
func New(ctx context.Context, address string) *Metrics {
	mux := http.NewServeMux()
	mux.Handle(handlerPatternMetrics, promhttp.Handler())

	metrics := &Metrics{}

	metrics.server = &http.Server{
		Handler: mux,
		Addr:    address,
	}

	return metrics
}

// Serve instantiates a server over http
func (metrics Metrics) Serve(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		if err := metrics.server.Shutdown(timeoutCtx); err != nil {
			log.Error(err)
		}
	}()

	logrus.WithFields(logrus.Fields{
		"address": metrics.server.Addr,
	}).Info("[ Metrics ] Listening")

	return metrics.server.ListenAndServe()
}
