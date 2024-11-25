package cmd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/DNS-Pro/core/pkg/app"
)

type serverCfg struct {
	DNS           app.DnsConfig
	Authenticator app.AuthenticatorConfig
}

type clientCfg app.ClientConfig

// ...
func (cfg *serverCfg) encodeString() (string, error) {
	v, err := json.Marshal(cfg)
	if err != nil {
		return "", fmt.Errorf("error marshaling config: %s", err)
	}
	return base64.StdEncoding.EncodeToString(v), nil
}

func decodeServerCfgString(s string) (*serverCfg, error) {
	var cfg serverCfg
	decodedBytes, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, fmt.Errorf("error decoding config: %s", err)
	}
	if err := json.Unmarshal(decodedBytes, &cfg); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %s", err)
	}
	return &cfg, nil
}
