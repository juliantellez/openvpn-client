package client

import (
	"context"
	"net/http"

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
	server := &http.Server{
		Addr: address,
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
