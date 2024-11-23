package app_test

import (
	"time"

	"github.com/DNS-Pro/core/internal/app"
	"github.com/DNS-Pro/core/internal/auth"
	"github.com/creasty/defaults"
	"github.com/go-playground/validator/v10"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {
	Describe("ClientConfig", Label("ClientConfig"), func() {
		type testCase struct {
			describe  string
			tType     TestCaseType
			cfg       app.ClientConfig
			expectErr error
		}
		tests := []testCase{
			{
				describe: "create a valid config",
				tType:    HAPPY_PATH,
				cfg: app.ClientConfig{
					BindAddress:     "1.2.3.4",
					HttpListenPort:  1000,
					SocksListenPort: 1001,
					QueryStrategy:   "UseIP",
					LogLevel:        "warning",
					RunAuthEvery:    10 * time.Second,
				},
				expectErr: nil,
			},
			{
				describe: "create a valid config (default bind ip, query strategy and log level)",
				tType:    HAPPY_PATH,
				cfg: app.ClientConfig{
					HttpListenPort:  1000,
					SocksListenPort: 1001,
					RunAuthEvery:    10 * time.Second,
				},
				expectErr: nil,
			},
			{
				describe: "error on empty HttpListenPort",
				tType:    FAILURE,
				cfg: app.ClientConfig{
					SocksListenPort: 1001,
					RunAuthEvery:    10 * time.Second,
				},
				expectErr: validator.ValidationErrors{},
			},
			{
				describe: "error on empty SocksListenPort",
				tType:    FAILURE,
				cfg: app.ClientConfig{
					HttpListenPort: 1000,
					RunAuthEvery:   10 * time.Second,
				},
				expectErr: validator.ValidationErrors{},
			},
			{
				describe: "error on empty RunAuthEvery",
				tType:    FAILURE,
				cfg: app.ClientConfig{
					HttpListenPort:  1000,
					SocksListenPort: 1001,
				},
				expectErr: validator.ValidationErrors{},
			},
			{
				describe: "error on invalid BindAddress",
				tType:    FAILURE,
				cfg: app.ClientConfig{
					BindAddress:     "invalid ip",
					HttpListenPort:  1000,
					SocksListenPort: 1001,
					RunAuthEvery:    10 * time.Second,
				},
				expectErr: validator.ValidationErrors{},
			},
			{
				describe: "error on invalid QueryStrategy",
				tType:    FAILURE,
				cfg: app.ClientConfig{
					QueryStrategy:   "invalid qs",
					HttpListenPort:  1000,
					SocksListenPort: 1001,
					RunAuthEvery:    10 * time.Second,
				},
				expectErr: validator.ValidationErrors{},
			},
			{
				describe: "error on invalid LogLevel",
				tType:    FAILURE,
				cfg: app.ClientConfig{
					LogLevel:        "invalid ll",
					HttpListenPort:  1000,
					SocksListenPort: 1001,
					RunAuthEvery:    10 * time.Second,
				},
				expectErr: validator.ValidationErrors{},
			},
		}
		// ...
		for _, tc := range tests {
			It(tc.describe, Label(string(tc.tType)), func() {
				// Arrange
				v := tc.cfg
				validate := validator.New(validator.WithRequiredStructEnabled())

				// Act
				if err := defaults.Set(&v); err != nil {
					AddReportEntry("error", err, ReportEntryVisibilityFailureOrVerbose)
					Expect(tc.expectErr).NotTo(BeNil())
					Expect(err).To(BeAssignableToTypeOf(tc.expectErr))
					return
				}
				if err := validate.Struct(&v); err != nil {
					AddReportEntry("error", err, ReportEntryVisibilityFailureOrVerbose)
					Expect(tc.expectErr).NotTo(BeNil())
					Expect(err).To(BeAssignableToTypeOf(tc.expectErr))
					return
				}
				// Assert
				Expect(tc.expectErr).To(BeNil())
			})
		}
	})
	Describe("DnsConfig", Label("DnsConfig"), func() {
		type testCase struct {
			describe  string
			tType     TestCaseType
			cfg       app.DnsConfig
			expectErr error
		}
		tests := []testCase{
			{
				describe: "create a valid config",
				tType:    HAPPY_PATH,
				cfg: app.DnsConfig{
					IP:   "1.2.3.4",
					Port: 1000,
				},
				expectErr: nil,
			},
			{
				describe: "error on empty ip",
				tType:    FAILURE,
				cfg: app.DnsConfig{
					Port: 1000,
				},
				expectErr: validator.ValidationErrors{},
			},
			{
				describe: "error on empty port",
				tType:    FAILURE,
				cfg: app.DnsConfig{
					IP: "1.2.3.4",
				},
				expectErr: validator.ValidationErrors{},
			},
			{
				describe: "error on invalid ip",
				tType:    FAILURE,
				cfg: app.DnsConfig{
					IP:   "invalid ip",
					Port: 1000,
				},
				expectErr: validator.ValidationErrors{},
			},
		}
		// ...
		for _, tc := range tests {
			It(tc.describe, Label(string(tc.tType)), func() {
				// Arrange
				v := tc.cfg
				validate := validator.New(validator.WithRequiredStructEnabled())

				// Act
				if err := defaults.Set(&v); err != nil {
					AddReportEntry("error", err, ReportEntryVisibilityFailureOrVerbose)
					Expect(tc.expectErr).NotTo(BeNil())
					Expect(err).To(BeAssignableToTypeOf(tc.expectErr))
					return
				}
				if err := validate.Struct(&v); err != nil {
					AddReportEntry("error", err, ReportEntryVisibilityFailureOrVerbose)
					Expect(tc.expectErr).NotTo(BeNil())
					Expect(err).To(BeAssignableToTypeOf(tc.expectErr))
					return
				}
				// Assert
				Expect(tc.expectErr).To(BeNil())
			})
		}
	})
	Describe("AuthenticatorConfig", Label("AuthenticatorConfig"), func() {
		type testCase struct {
			describe  string
			tType     TestCaseType
			cfg       app.AuthenticatorConfig
			expectErr error
		}
		tests := []testCase{
			{
				describe: "create a valid AUTH_NONE config",
				tType:    HAPPY_PATH,
				cfg: app.AuthenticatorConfig{
					Type: auth.AUTH_NONE,
				},
				expectErr: nil,
			},
			{
				describe: "create a valid AUTH_HTTP config",
				tType:    HAPPY_PATH,
				cfg: app.AuthenticatorConfig{
					Type: auth.AUTH_HTTP,
					Url:  "http://url.com",
				},
				expectErr: nil,
			},
			{
				describe: "create a valid AUTH_UNKNOWN config",
				tType:    HAPPY_PATH,
				cfg: app.AuthenticatorConfig{
					Type: auth.AUTH_UNKNOWN,
				},
				expectErr: nil,
			},
			{
				describe:  "create with default auth type",
				tType:     HAPPY_PATH,
				cfg:       app.AuthenticatorConfig{},
				expectErr: nil,
			},
		}
		// ...
		for _, tc := range tests {
			It(tc.describe, Label(string(tc.tType)), func() {
				// Arrange
				v := tc.cfg
				validate := validator.New(validator.WithRequiredStructEnabled())

				// Act
				if err := defaults.Set(&v); err != nil {
					AddReportEntry("error", err, ReportEntryVisibilityFailureOrVerbose)
					Expect(tc.expectErr).NotTo(BeNil())
					Expect(err).To(BeAssignableToTypeOf(tc.expectErr))
					return
				}
				if err := validate.Struct(&v); err != nil {
					AddReportEntry("error", err, ReportEntryVisibilityFailureOrVerbose)
					Expect(tc.expectErr).NotTo(BeNil())
					Expect(err).To(BeAssignableToTypeOf(tc.expectErr))
					return
				}
				// Assert
				Expect(tc.expectErr).To(BeNil())
			})
		}
	})
})
