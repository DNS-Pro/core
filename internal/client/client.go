package client

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime"

	"github.com/DNS-Pro/core/internal/errs"
	"github.com/go-playground/validator/v10"
	core "github.com/v2fly/v2ray-core/v5"
	v2net "github.com/v2fly/v2ray-core/v5/common/net"
	"github.com/v2fly/v2ray-core/v5/infra/conf/cfgcommon"
	"github.com/v2fly/v2ray-core/v5/infra/conf/cfgcommon/sniffer"
	"github.com/v2fly/v2ray-core/v5/infra/conf/synthetic/dns"
	logCfg "github.com/v2fly/v2ray-core/v5/infra/conf/synthetic/log"
	"github.com/v2fly/v2ray-core/v5/infra/conf/synthetic/router"
	conf "github.com/v2fly/v2ray-core/v5/infra/conf/v4"
)

// DnsAddress represents the DNS server's IP and port.
type DnsAddress struct {
	IP   string `validate:"required,ip"`
	Port uint16 `validate:"required"`
}

type IClient interface {
	Start(context.Context) error
	GenerateConfig() (*core.Config, error)
	GenerateClient(cfg *core.Config) (*core.Instance, error)
}

// Client encapsulates configuration details for the V2Ray Client.
type Client struct {
	DnsAddress      DnsAddress `validate:"required"`
	BindAddress     string     `validate:"required,ip"`
	HttpListenPort  uint32     `validate:"required"`
	SocksListenPort uint32     `validate:"required"`
	QueryStrategy   string     `validate:"required,oneof=UseIP UseIPv4 UseIPv6"`
	LogLevel        string     `validate:"required,oneof=debug info warning error none"`
}

// GenerateV4Config creates a V4 configuration for the V2Ray client.
func (cl *Client) GenerateV4Config() *conf.Config {
	listenAddr := cfgcommon.Address{Address: v2net.ParseAddress(cl.BindAddress)}
	dnsAddr := cfgcommon.Address{Address: v2net.ParseAddress(cl.DnsAddress.IP)}

	socksInboundSettings := json.RawMessage([]byte(`{"udp":true}`))
	directOutboundSettings := json.RawMessage([]byte(`{"domainStrategy":"UseIP"}`))
	dnsOutboundSettings := json.RawMessage([]byte(fmt.Sprintf(`{"address":"%s","network":"udp","port":%d,"userLevel":1}`, cl.DnsAddress.IP, cl.DnsAddress.Port)))

	routeDomainStrategy := "AsIs"

	return &conf.Config{
		LogConfig: &logCfg.LogConfig{
			LogLevel: cl.LogLevel,
		},
		InboundConfigs: []conf.InboundDetourConfig{
			createInboundConfig("socks", cl.SocksListenPort, listenAddr, &socksInboundSettings, "socks-in"),
			createInboundConfig("http", cl.HttpListenPort, listenAddr, nil, "http-in"),
		},
		RouterConfig: &router.RouterConfig{
			RuleList: []json.RawMessage{
				json.RawMessage([]byte(`{
					"inboundTag":["socks-in","http-in"],
					"outboundTag":"dns-out",
					"port":"53",
					"type":"field"
				}`)),
				json.RawMessage([]byte(`{
					"outboundTag":"direct",
					"port":"0-65535",
					"type":"field"
				}`)),
			},
			DomainStrategy: &routeDomainStrategy,
		},
		DNSConfig: &dns.DNSConfig{
			Servers: []*dns.NameServerConfig{
				{
					Address: &dnsAddr,
					Port:    cl.DnsAddress.Port,
				},
			},
			QueryStrategy: cl.QueryStrategy,
			Tag:           "dns",
		},
		OutboundConfigs: []conf.OutboundDetourConfig{
			{
				Protocol: "freedom",
				Tag:      "direct",
				Settings: &directOutboundSettings,
			},
			{
				Protocol: "dns",
				Tag:      "dns-out",
				Settings: &dnsOutboundSettings,
			},
		},
	}
}

// GenerateConfig builds the core configuration for the V2Ray client.
func (cl *Client) GenerateConfig() (*core.Config, error) {
	return cl.GenerateV4Config().Build()
}

// GenerateClient initializes a new V2Ray client instance.
func (cl *Client) GenerateClient(cfg *core.Config) (*core.Instance, error) {
	return core.New(cfg)
}

// Start launches the client and waits for termination signals.
func (cl *Client) Start_(ctx context.Context, clientInstance *core.Instance) error {
	if err := clientInstance.Start(); err != nil {
		return fmt.Errorf("failed to start client: %w", err)
	}
	defer clientInstance.Close()

	runtime.GC()

	<-ctx.Done()

	return nil
}

// AutoStart, generates config and client and launches the client
func (cl *Client) Start(ctx context.Context) error {
	cfg, err := cl.GenerateConfig()
	if err != nil {
		return fmt.Errorf("error generating client config: %s", err)
	}
	client, err := cl.GenerateClient(cfg)
	if err != nil {
		return fmt.Errorf("error generating client: %s", err)
	}
	return cl.Start_(ctx, client)
}

// NewClient validates config and initializes a new client. default values are injected if none are provided.
//
// Using factory is the only way to create a client, so validated configs are ensured.
func NewClient(dnsAddr DnsAddress, bindAddr string, httpPort, socksPort uint32, queryStrategy, logLevel string) (IClient, error) {
	cl := &Client{
		DnsAddress:      dnsAddr,
		BindAddress:     bindAddr,
		HttpListenPort:  httpPort,
		SocksListenPort: socksPort,
		QueryStrategy:   queryStrategy,
		LogLevel:        logLevel,
	}
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(cl); err != nil {
		return nil, errs.NewConfigValidationErr(err)
	}
	return cl, nil
}

// createInboundConfig is a helper to simplify inbound configuration creation.
func createInboundConfig(protocol string, port uint32, listenAddr cfgcommon.Address, settings *json.RawMessage, tag string) conf.InboundDetourConfig {
	return conf.InboundDetourConfig{
		Protocol:  protocol,
		PortRange: &cfgcommon.PortRange{From: port, To: port},
		ListenOn:  &listenAddr,
		Settings:  settings,
		Tag:       tag,
		SniffingConfig: &sniffer.SniffingConfig{
			Enabled:      true,
			DestOverride: cfgcommon.NewStringList([]string{"http", "tls", "quic"}),
			MetadataOnly: false,
		},
	}
}
