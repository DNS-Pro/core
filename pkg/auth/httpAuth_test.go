package auth_test

import (
	"context"
	"net/http"

	"github.com/DNS-Pro/core/pkg/auth"
	"github.com/DNS-Pro/core/pkg/errs"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("HttpAuth", func() {
	Describe("NewHttpAuther", Label("NewHttpAuther"), func() {
		type testCase struct {
			name      string
			tType     TestCaseType
			url       string
			expectErr error // expect validation error
		}
		// ...
		tests := []testCase{
			{
				name:  "Valid data",
				tType: HAPPY_PATH,
				url:   "http://example.com",
			},
			{
				name:  "Valid data (with path)",
				tType: HAPPY_PATH,
				url:   "http://example.com/example",
			},
			{
				name:  "Valid data (with port)",
				tType: HAPPY_PATH,
				url:   "http://example.com:1020",
			},
			{
				name:      "Invalid data (no url schema)",
				tType:     FAILURE,
				url:       "example.com",
				expectErr: errs.AppConfigValidationErr{},
			},
			{
				name:      "Invalid data (invalid url schema)",
				tType:     FAILURE,
				url:       "invalid://example.com",
				expectErr: errs.AppConfigValidationErr{},
			},
			{
				name:      "Invalid data (invalid port)",
				tType:     FAILURE,
				url:       "http://example.com:test",
				expectErr: errs.AppConfigValidationErr{},
			},
		}
		// ...
		for _, tt := range tests {
			It(tt.name, Label(string(tt.tType)), func() {
				// Arrange
				// Act
				auther, err := auth.NewHttpAuther(tt.url)
				// Assert
				if tt.expectErr != nil {
					Expect(err).ToNot(BeNil())
					Expect(err).To(BeAssignableToTypeOf(tt.expectErr))
					Expect(auther).To(BeNil())
				} else {
					Expect(err).To(BeNil())
					Expect(auther).NotTo(BeNil())
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
				httpAuther, err := auth.NewHttpAuther(server.URL())
				Expect(err).To(BeNil())
				// Act
				err = httpAuther.Run(context.TODO())
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
