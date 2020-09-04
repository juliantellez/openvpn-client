package main

import (
	"os"

	"github.com/juliantellez/openvpn-client/internal/client"
	"github.com/juliantellez/openvpn-client/shared/logger"
	"github.com/juliantellez/openvpn-client/shared/prometheus"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"golang.org/x/sync/errgroup"
)

const (
	appName             = "openvpnClient"
	appVersion          = "0.1.0"
	errorServiceFailure = "Failed to run service"
)

var (
	config struct {
		LogLevel          string
		LogOutput         string
		AddressPrometheus string
		AddressClient     string
		OpenVpnConfig     string
	}

	flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "log-level",
			Usage:       "Log Level",
			EnvVars:     []string{"LOG_LEVEL"},
			Value:       "trace",
			Destination: &config.LogLevel,
		},
		&cli.StringFlag{
			Name:        "log-output",
			Usage:       "Log Output `text or json`",
			EnvVars:     []string{"LOG_OUTPUT"},
			Value:       "json",
			Destination: &config.LogOutput,
		},
		&cli.StringFlag{
			Name:        "prometheus-address",
			Usage:       "Prometheus Address exposes '/__/metrics' ",
			EnvVars:     []string{"PROMETHEUS_ADDRESS"},
			Value:       ":8081",
			Destination: &config.AddressPrometheus,
		},
		&cli.StringFlag{
			Name:        "client-address",
			Usage:       "Client Address exposes the vpn wrapper",
			EnvVars:     []string{"CLIENT_ADDRESS"},
			Value:       ":7979",
			Destination: &config.AddressClient,
		},
		&cli.StringFlag{
			Name:        "openvpn-config",
			Usage:       "OVPN config file",
			EnvVars:     []string{"OPENVPN_CONFIG"},
			Destination: &config.OpenVpnConfig,
		},
	}
)

func appAction(cliCtx *cli.Context) error {
	if err := logger.New(config.LogLevel, config.LogOutput); err != nil {
		return err
	}

	ctx := cliCtx.Context

	errorGroup, ctx := errgroup.WithContext(ctx)
	metrics := prometheus.New(ctx, config.AddressPrometheus)
	client := client.New(ctx, config.AddressClient)

	errorGroup.Go(func() error {
		return metrics.Serve(ctx)
	})

	errorGroup.Go(func() error {
		return client.Serve()
	})

	return errorGroup.Wait()
}

func main() {
	app := cli.NewApp()
	app.Name = appName
	app.Flags = flags
	app.Version = appVersion

	app.Action = appAction

	if err := app.Run(os.Args); err != nil {
		logrus.WithError(err).Fatal(errorServiceFailure)
	}
}
