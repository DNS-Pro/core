package auth_test

import (
	"context"
	"net/http"

	"github.com/DNS-Pro/core/internal/auth"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("HttpAuth", func() {
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
				httpAuther := auth.NewHttpAuther(server.URL())
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
