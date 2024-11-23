package app

import (
	"context"
	"fmt"

	"github.com/DNS-Pro/core/pkg/auth"
	"github.com/DNS-Pro/core/pkg/client"
	"github.com/sirupsen/logrus"
)

type app struct {
	config        *appConfig
	authenticator auth.IAuthenticator
	dnsCl         client.IClient
}

func (a *app) Run(ctx context.Context) {
	go func() {
		if err := a.authenticator.Start(ctx); err != nil {
			logrus.WithError(err).Error("error in authenticator process")
		}
	}()
	if err := a.dnsCl.Start(ctx); err != nil {
		logrus.WithError(err).Error("error in dns client process")
	}
}
func NewApp(cfg *appConfig) (*app, error) {
	v := app{
		config: cfg,
	}
	// ...
	authenticator, err := GetAuthenticator(&cfg.Authenticator, &cfg.Client)
	if err != nil {
		return nil, err
	}
	v.authenticator = authenticator
	// ...
	cl, err := GetClient(&cfg.DNS, &cfg.Client)
	if err != nil {
		return nil, err
	}
	v.dnsCl = cl
	// ...
	return &v, nil
}

func GetAuther(authCfg *AuthenticatorConfig) (auth.IAuther, error) {
	var auther auth.IAuther
	var err error
	switch authCfg.Type {
	case auth.AUTH_NONE:
		auther = nil
	case auth.AUTH_HTTP:
		auther, err = auth.NewHttpAuther(authCfg.Url)
	default:
		return nil, fmt.Errorf("unknown auther type %d", authCfg.Type)
	}
	return auther, err
}
func GetAuthenticator(authCfg *AuthenticatorConfig, clCfg *ClientConfig) (auth.IAuthenticator, error) {
	auther, err := GetAuther(authCfg)
	if err != nil {
		return nil, err
	}
	return auth.NewAuthenticator(clCfg.RunAuthEvery, auther)
}
func GetClient(dnsCfg *DnsConfig, clCfg *ClientConfig) (client.IClient, error) {
	cl, err := client.NewClient(
		client.DnsAddress{
			IP:   dnsCfg.IP,
			Port: dnsCfg.Port,
		},
		clCfg.BindAddress,
		clCfg.HttpListenPort,
		clCfg.SocksListenPort,
		clCfg.QueryStrategy,
		clCfg.LogLevel,
	)
	return cl, err
}
