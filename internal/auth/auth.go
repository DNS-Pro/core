package auth

import (
	"context"
	"time"

	"github.com/DNS-Pro/core/internal/errs"
	"github.com/sirupsen/logrus"
)

type IAuther interface {
	Run(ctx context.Context) error
	Validate() error
	SetDefaults() error
}
type AuthType int

const (
	AUTH_NONE AuthType = iota
	AUTH_HTTP
	AUTH_UNKNOWN
)

func (at *AuthType) FromAuthenticator(authenticator IAuther) {
	switch authenticator.(type) {
	case nil:
		*at = AUTH_NONE
	case *HttpAuther:
		*at = AUTH_HTTP
	default:
		*at = AUTH_UNKNOWN
	}
}

// ...
type Authenticator struct {
	runEvery       time.Duration
	aType          AuthType
	iAuthenticator IAuther
}

func (a *Authenticator) getLogger() *logrus.Entry {
	return logrus.WithField("module", "auth").WithField("authenticator", a.aType)
}

func (a *Authenticator) Start(ctx context.Context) error {
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

// validate auther and create new Authenticator.
func NewAuthenticator(interval time.Duration, auther IAuther) (*Authenticator, error) {
	v := Authenticator{
		runEvery:       interval,
		iAuthenticator: auther,
	}
	v.aType.FromAuthenticator(auther)
	if auther != nil {
		if err := auther.SetDefaults(); err != nil {
			return nil, errs.NewAppConfigDefaultValueErr(err)
		}
		if err := auther.Validate(); err != nil {
			return nil, errs.NewAppConfigValidationErr(err)
		}
	}
	return &v, nil
}
