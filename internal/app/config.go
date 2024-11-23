package app

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/DNS-Pro/core/internal/auth"
	"github.com/DNS-Pro/core/internal/errs"
	"github.com/creasty/defaults"
	"github.com/go-playground/validator/v10"
)

type ClientConfig struct {
	BindAddress     string        `validate:"required,ip" default:"127.0.0.1"`
	HttpListenPort  uint32        `validate:"required"`
	SocksListenPort uint32        `validate:"required"`
	QueryStrategy   string        `validate:"required,oneof=UseIP UseIPv4 UseIPv6" default:"UseIP"`
	LogLevel        string        `validate:"required,oneof=debug info warning error none" default:"warning"`
	RunAuthEvery    time.Duration `validate:"required"`
}
type DnsConfig struct {
	IP   string `validate:"required,ip"`
	Port uint16 `validate:"required"`
}
type AuthenticatorConfig struct {
	Type auth.AuthType `default:"0"`
	Url  string
}
type appConfig struct {
	DNS           DnsConfig
	Client        ClientConfig `json:"-"`
	Authenticator AuthenticatorConfig
}

func (cfg *appConfig) EncodeString() (string, error) {
	v, err := json.Marshal(cfg)
	if err != nil {
		return "", fmt.Errorf("error marshaling config: %s", err)
	}
	return base64.StdEncoding.EncodeToString(v), nil
}
func (cfg *appConfig) DecodeString(s string) error {
	decodedBytes, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return fmt.Errorf("error decoding config: %s", err)
	}
	if err := json.Unmarshal(decodedBytes, cfg); err != nil {
		return fmt.Errorf("error unmarshaling config: %s", err)
	}
	return nil
}

// NewAppConfig validates and creates appConfig
//
// Using factory is the only way to create an appConfig, so validated configs are ensured.
func NewAppConfig(clientConf *ClientConfig, dnsConfig *DnsConfig, authenticatorConfig *AuthenticatorConfig) (*appConfig, error) {
	v := &appConfig{
		Client:        *clientConf,
		DNS:           *dnsConfig,
		Authenticator: *authenticatorConfig,
	}
	if err := defaults.Set(v); err != nil {
		return nil, errs.NewConfigDefaultValueErr(err)
	}
	if err := getAppConfigValidator().Struct(&v); err != nil {
		return nil, errs.NewConfigValidationErr(err)
	}
	return v, nil
}
func NewAppConfigFromString(clientConf *ClientConfig, configStr string) (*appConfig, error) {
	v := &appConfig{}
	if err := v.DecodeString(configStr); err != nil {
		return nil, err
	}
	return NewAppConfig(clientConf, &v.DNS, &v.Authenticator)
}

// ...
func getAppConfigValidator() *validator.Validate {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate
}
