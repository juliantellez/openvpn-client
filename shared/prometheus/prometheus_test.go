package prometheus_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/juliantellez/openvpn-client/shared/prometheus"
	"github.com/stretchr/testify/assert"
)

func TestPrometheus(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	metrics := prometheus.New(ctx, ":8081")

	go func() {
		err := metrics.Serve(ctx)
		assert.NoError(t, err)
	}()

	time.Sleep(time.Millisecond * 200)
	response, err := http.Get("http://localhost:8081/__/metrics")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	cancel()
}
