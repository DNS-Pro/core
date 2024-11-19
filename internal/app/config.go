package app

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/DNS-Pro/core/internal/auth"
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
	Type auth.AuthType `validate:"required"`
	Url  string        `validate:"http_url,required_if=Type 0"`
}
type AppConfig struct {
	DNS           DnsConfig
	Client        ClientConfig `json:"-"`
	Authenticator AuthenticatorConfig
}

func (cfg *AppConfig) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(cfg)
}
func (cfg *AppConfig) EncodeString() (string, error) {
	v, err := json.Marshal(cfg)
	if err != nil {
		return "", fmt.Errorf("error marshaling config: %s", err)
	}
	return base64.StdEncoding.EncodeToString(v), nil
}
func (cfg *AppConfig) DecodeString(s string) error {
	decodedBytes, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return fmt.Errorf("error decoding config: %s", err)
	}
	if err := json.Unmarshal(decodedBytes, cfg); err != nil {
		return fmt.Errorf("error unmarshaling config: %s", err)
	}
	return nil
}
func NewAppConfig(clientConf *ClientConfig, configStr string) (*AppConfig, error) {
	v := &AppConfig{}
	if err := v.DecodeString(configStr); err != nil {
		return nil, err
	}
	v.Client = *clientConf
	if err := defaults.Set(v); err != nil {
		return nil, fmt.Errorf("error setting defaul values: %s", err)
	}
	if err := v.Validate(); err != nil {
		return nil, fmt.Errorf("error validating config: %s", err)
	}
	return v, nil
}
