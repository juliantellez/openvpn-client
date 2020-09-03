package client_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/juliantellez/openvpn-client/internal/client"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	client := client.New(ctx, ":7979")

	go func() {
		err := client.Serve()
		assert.NoError(t, err)
	}()

	time.Sleep(time.Millisecond * 200)
	response, err := http.Get("http://localhost:7979")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, response.StatusCode)
	cancel()
}
