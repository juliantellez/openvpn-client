package prometheus

import (
	"context"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
)

const (
	handlerPatternMetrics = "/__/metrics"
)

// New instantiates a handler for prometheus
func New(ctx context.Context, address string) error {
	mux := http.NewServeMux()
	mux.Handle(handlerPatternMetrics, promhttp.Handler())

	server := http.Server{
		Handler: mux,
		Addr:    address,
	}

	go func() {
		<-ctx.Done()
		timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		if err := server.Shutdown(timeoutCtx); err != nil {
			log.Error(err)
		}
	}()

	return server.ListenAndServe()
}
