package client_test

import (
	"context"
	"time"

	"github.com/DNS-Pro/core/internal/client"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client", func() {
	Describe("NewClient", Label("NewClient"), func() {
		tests := []struct {
			name                    string
			tType                   TestCaseType
			dnsAddr                 client.DnsAddress
			bindIP                  string
			httpPort, socksPort     uint32
			queryStrategy, logLevel string
			wantErr                 bool
		}{
			{
				name:  "Valid Parameters",
				tType: HAPPY_PATH,
				dnsAddr: client.DnsAddress{
					IP:   "8.8.8.8",
					Port: 53,
				},
				bindIP:        "127.0.0.1",
				httpPort:      8080,
				socksPort:     1080,
				queryStrategy: "UseIP",
				logLevel:      "info",
				wantErr:       false,
			},
			{
				name:  "Fail on invalid DNS IP",
				tType: FAILURE,
				dnsAddr: client.DnsAddress{
					IP:   "invalid-ip",
					Port: 53,
				},
				bindIP:        "127.0.0.1",
				httpPort:      8080,
				socksPort:     1080,
				queryStrategy: "UseIP",
				logLevel:      "info",
				wantErr:       true,
			},
			{
				name:  "Fail on invalid Bind IP",
				tType: FAILURE,
				dnsAddr: client.DnsAddress{
					IP:   "8.8.8.8",
					Port: 53,
				},
				bindIP:        "invalid-ip",
				httpPort:      8080,
				socksPort:     1080,
				queryStrategy: "UseIP",
				logLevel:      "info",
				wantErr:       true,
			},
			{
				name:  "Fail on invalid Query Strategy",
				tType: FAILURE,
				dnsAddr: client.DnsAddress{
					IP:   "8.8.8.8",
					Port: 53,
				},
				bindIP:        "127.0.0.1",
				httpPort:      8080,
				socksPort:     1080,
				queryStrategy: "InvalidStrategy",
				logLevel:      "info",
				wantErr:       true,
			},
			{
				name:  "Fail on invalid Log Level",
				tType: FAILURE,
				dnsAddr: client.DnsAddress{
					IP:   "8.8.8.8",
					Port: 53,
				},
				bindIP:        "127.0.0.1",
				httpPort:      8080,
				socksPort:     1080,
				queryStrategy: "UseIP",
				logLevel:      "invalid-loglevel",
				wantErr:       true,
			},
		}
		for _, tt := range tests {
			It(tt.name, func() {
				cl, err := client.NewClient(tt.dnsAddr, tt.bindIP, tt.httpPort, tt.socksPort, tt.queryStrategy, tt.logLevel)
				if tt.wantErr {
					Expect(err).NotTo(BeNil())
					Expect(cl).To(BeNil())
				} else {
					Expect(err).To(BeNil())
					Expect(cl).NotTo(BeNil())
				}
			})

		}
	})
	Describe("GenerateClient", Label("GenerateClient"), func() {
		tests := []struct {
			name                    string
			tType                   TestCaseType
			dnsAddr                 client.DnsAddress
			bindIP                  string
			httpPort, socksPort     uint32
			queryStrategy, logLevel string
		}{
			{
				name:  "Valid Parameters",
				tType: HAPPY_PATH,
				dnsAddr: client.DnsAddress{
					IP:   "8.8.8.8",
					Port: 53,
				},
				bindIP:        "127.0.0.1",
				httpPort:      8080,
				socksPort:     1080,
				queryStrategy: "UseIP",
				logLevel:      "info",
			},
		}
		for _, tt := range tests {
			It(tt.name, func() {
				client, err := client.NewClient(tt.dnsAddr, tt.bindIP, tt.httpPort, tt.socksPort, tt.queryStrategy, tt.logLevel)
				Expect(err).To(BeNil())
				cfg, err := client.GenerateConfig()
				Expect(err).To(BeNil())
				_, err = client.GenerateClient(cfg)
				Expect(err).To(BeNil())
			})
		}
	})
	Describe("Start", Label("StartClient"), func() {
		tests := []struct {
			name                    string
			tType                   TestCaseType
			dnsAddr                 client.DnsAddress
			bindIP                  string
			httpPort, socksPort     uint32
			queryStrategy, logLevel string
		}{
			{
				name:  "Valid Parameters",
				tType: HAPPY_PATH,
				dnsAddr: client.DnsAddress{
					IP:   "8.8.8.8",
					Port: 53,
				},
				bindIP:        "127.0.0.1",
				httpPort:      8080,
				socksPort:     1080,
				queryStrategy: "UseIP",
				logLevel:      "info",
			},
		}
		for _, tt := range tests {
			It(tt.name, func() {
				client, err := client.NewClient(tt.dnsAddr, tt.bindIP, tt.httpPort, tt.socksPort, tt.queryStrategy, tt.logLevel)
				Expect(err).To(BeNil())
				// ...
				ctx, cncl := context.WithTimeout(context.Background(), 100*time.Microsecond)
				err = client.Start(ctx)
				Expect(err).To(BeNil())
				Expect(ctx.Err()).ToNot(BeNil())
				cncl()
			})
		}
	})

})
