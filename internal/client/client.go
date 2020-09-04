package client

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juliantellez/openvpn-client/internal/client/handlers"
	"github.com/juliantellez/openvpn-client/internal/client/middlewares"
	"github.com/sirupsen/logrus"
)

type (
	// Client vpn
	Client struct {
		server *http.Server
	}
)

// New configures the vpn client
func New(ctx context.Context, address string) *Client {
	gin.SetMode(gin.ReleaseMode)
	engineHandler := gin.New()
	middlewares.Bind(engineHandler)
	handlers.Bind(engineHandler)

	server := &http.Server{
		Addr:    address,
		Handler: engineHandler,
	}

	client := &Client{}
	client.server = server
	return client
}

// Serve instantiates a server over http
func (client Client) Serve() error {
	logrus.WithFields(logrus.Fields{
		"address": client.server.Addr,
	}).Info("[ Client ] Listening")

	return client.server.ListenAndServe()
}
