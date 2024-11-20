package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/onsi/ginkgo/v2"
	"github.com/sirupsen/logrus"
)

type IAuther interface {
	Run(ctx context.Context) error
	SetBaseAuth(*Authenticator)
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
	case *HttpAuthenticator:
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

func NewAuthenticator(interval time.Duration, auther IAuther) (*Authenticator, error) {
	ginkgo.GinkgoWriter.Printf("auther: %v", auther)
	v := Authenticator{
		runEvery:       interval,
		iAuthenticator: auther,
	}
	v.aType.FromAuthenticator(auther)
	ginkgo.GinkgoWriter.Printf("auther2: %+v", v)
	if auther != nil {
		if err := auther.SetDefaults(); err != nil {
			return nil, fmt.Errorf("error setting defaul values: %s", err)
		}
		if err := auther.Validate(); err != nil {
			return nil, fmt.Errorf("error validating authenticator: %s", err)
		}
		auther.SetBaseAuth(&v)
	}
	return &v, nil
}
