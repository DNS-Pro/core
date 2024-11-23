package app

import (
	"context"
	"fmt"

	"github.com/DNS-Pro/core/internal/auth"
	"github.com/DNS-Pro/core/internal/client"
	"github.com/sirupsen/logrus"
)

type app struct {
	config        *appConfig
	authenticator *auth.Authenticator
	dnsCl         *client.Client
}

func (a *app) Run(ctx context.Context) {
	go func() {
		if err := a.authenticator.Start(ctx); err != nil {
			logrus.WithError(err).Error("error in authenticator process")
		}
	}()
	if err := a.dnsCl.AutoStart(ctx); err != nil {
		logrus.WithError(err).Error("error in dns client process")
	}
}
func NewApp(cfg *appConfig) (*app, error) {
	v := app{
		config: cfg,
	}
	// ...
	authenticator, err := getAuthenticator(&cfg.Authenticator, &cfg.Client)
	if err != nil {
		return nil, err
	}
	v.authenticator = authenticator
	// ...
	cl, err := client.NewClient(
		client.DnsAddress{
			IP:   cfg.DNS.IP,
			Port: cfg.DNS.Port,
		},
		cfg.Client.BindAddress,
		cfg.Client.HttpListenPort,
		cfg.Client.SocksListenPort,
		cfg.Client.QueryStrategy,
		cfg.Client.LogLevel)
	if err != nil {
		return nil, err
	}
	v.dnsCl = cl
	// ...
	return &v, nil
}

func getAuthenticator(authCfg *AuthenticatorConfig, clCfg *ClientConfig) (*auth.Authenticator, error) {
	var auther auth.IAuther
	switch authCfg.Type {
	case auth.AUTH_NONE:
		auther = nil
	case auth.AUTH_HTTP:
		auther = &auth.HttpAuther{Url: authCfg.Url}
	default:
		return nil, fmt.Errorf("unknown auther type %d", authCfg.Type)
	}
	return auth.NewAuthenticator(clCfg.RunAuthEvery, auther)
}
