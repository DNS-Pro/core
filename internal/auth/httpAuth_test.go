package auth_test

import (
	"context"
	"net/http"

	"github.com/DNS-Pro/core/internal/auth"
	"github.com/go-playground/validator/v10"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("HttpAuth", func() {
	Describe("Validate", Label("Validate"), func() {
		type testCase struct {
			name      string
			tType     TestCaseType
			input     *auth.HttpAuther
			expectErr error // expect validation error
		}
		// ...
		tests := []testCase{
			{
				name:  "Valid data",
				tType: HAPPY_PATH,
				input: &auth.HttpAuther{"http://example.com"},
			},
			{
				name:  "Valid data (with path)",
				tType: HAPPY_PATH,
				input: &auth.HttpAuther{"http://example.com/example"},
			},
			{
				name:  "Valid data (with port)",
				tType: HAPPY_PATH,
				input: &auth.HttpAuther{"http://example.com:1020"},
			},
			{
				name:      "Invalid data (no url schema)",
				tType:     FAILURE,
				input:     &auth.HttpAuther{"example.com"},
				expectErr: validator.ValidationErrors{},
			},
			{
				name:      "Invalid data (invalid url schema)",
				tType:     FAILURE,
				input:     &auth.HttpAuther{"invalid://example.com"},
				expectErr: validator.ValidationErrors{},
			},
			{
				name:      "Invalid data (invalid port)",
				tType:     FAILURE,
				input:     &auth.HttpAuther{"http://example.com:test"},
				expectErr: validator.ValidationErrors{},
			},
		}
		// ...
		for _, tt := range tests {
			It(tt.name, Label(string(tt.tType)), func() {
				// Arrange
				// Act
				err := tt.input.Validate()
				// Assert
				if tt.expectErr != nil {
					Expect(err).ToNot(BeNil())
					Expect(err).To(BeAssignableToTypeOf(tt.expectErr))
				} else {
					Expect(err).To(BeNil())
				}
			})
		}
	})
	Describe("Run", Label("Run"), func() {
		type testCase struct {
			name          string
			tType         TestCaseType
			expectHttpErr bool // expect error response from server
			expectErr     bool // expect error in call
		}
		var server *ghttp.Server
		// ...
		setupServer := func(tc testCase) {
			var stCode int
			if tc.expectHttpErr {
				stCode = http.StatusUnauthorized
			} else {
				stCode = http.StatusAccepted
			}
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/"),
					ghttp.RespondWith(stCode, nil),
				),
			)
		}
		// ...
		BeforeEach(func() {
			server = ghttp.NewServer()
		})

		AfterEach(func() {
			server.Close()
		})
		// ...
		tests := []testCase{
			{
				name:  "Successfully call remote server",
				tType: HAPPY_PATH,
			},
			{
				name:          "Error from remote server",
				tType:         FAILURE,
				expectHttpErr: true,
				expectErr:     true,
			},
		}
		// ...
		for _, tt := range tests {
			It(tt.name, Label(string(tt.tType)), func() {
				// Arrange
				setupServer(tt)
				httpAuther := auth.HttpAuther{Url: server.URL()}
				// Act
				err := httpAuther.Run(context.TODO())
				// Assert
				Expect(server.ReceivedRequests()).Should(HaveLen(1))
				if tt.expectErr {
					Expect(err).ToNot(BeNil())
				} else {
					Expect(err).To(BeNil())
				}
			})
		}

	})
})
