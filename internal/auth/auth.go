package auth

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type IAuthenticator interface {
	run(ctx context.Context) error
}
type AuthType int

const (
	AUTH_HTTP AuthType = iota
)

type Authenticator struct {
	runEvery       time.Duration
	aType          AuthType `validate:"required"`
	iAuthenticator IAuthenticator
}

func (a *Authenticator) getLogger() *logrus.Entry {
	return logrus.WithField("module", "auth").WithField("authenticator", a.aType)
}

func (a *Authenticator) Start(ctx context.Context) error {
	ticker := time.NewTicker(a.runEvery)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			a.getLogger().Info("Context canceled")
			return nil
		case <-ticker.C:
			a.getLogger().Debug("calling run")
			if err := a.iAuthenticator.run(ctx); err != nil {
				a.getLogger().WithError(err).Error("error calling run")
			} else {
				a.getLogger().Debug("done calling run")
			}
		}
	}
}
