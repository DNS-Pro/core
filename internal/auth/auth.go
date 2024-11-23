package auth

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type IAuther interface {
	Run(ctx context.Context) error
	GetType() AuthType
}

// ...
type AuthType int

const (
	AUTH_NONE AuthType = iota
	AUTH_HTTP
	AUTH_UNKNOWN
)

// ...
type IAuthenticator interface {
	Start(ctx context.Context) error
}
type authenticator struct {
	runEvery       time.Duration
	iAuthenticator IAuther
}

func (a *authenticator) getLogger() *logrus.Entry {
	return logrus.WithField("module", "authenticator").WithField("auther", a.iAuthenticator.GetType())
}

func (a *authenticator) Start(ctx context.Context) error {
	if a.iAuthenticator == nil {
		return nil
	}
	ticker := time.NewTicker(a.runEvery)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			a.getLogger().Info("Context canceled")
			return nil
		case <-ticker.C:
			a.getLogger().Debug("calling run")
			if err := a.iAuthenticator.Run(ctx); err != nil {
				a.getLogger().WithError(err).Error("error calling run")
			} else {
				a.getLogger().Debug("done calling run")
			}
		}
	}
}

// NewAuthenticator validates and creates new Authenticator.
//
// Using factory is the only way to create a authenticator, so validated configs are ensured.
func NewAuthenticator(interval time.Duration, auther IAuther) (IAuthenticator, error) {
	v := authenticator{
		runEvery:       interval,
		iAuthenticator: auther,
	}
	return &v, nil
}
