package auth_test

import (
	"time"

	"github.com/DNS-Pro/core/internal/auth"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Auth", func() {
	Describe("AuthType", Label("AuthType"), func() {
	})
	Describe("Authenticator", Label("Authenticator"), func() {
		Describe("NewAuthenticator", Label("NewAuthenticator"), func() {
			type testCase struct {
				name             string
				tType            TestCaseType
				withAuther       bool // whether or not to create with auther
				expectSetDefault bool // expect calling authenticator SetDefault method
				expectValidate   bool // expect calling authenticator Validate method
				expectErr        bool // expect error
			}
			// ...
			BeforeEach(func() {
				resetMock()
			})

			AfterEach(func() {
				asserMockCall()
			})
			// ...
			tests := []testCase{
				{
					name:       "Successfully Create AUTH_NONE",
					tType:      HAPPY_PATH,
					withAuther: false,
				},
				{
					name:             "Successfully Create authenticator",
					tType:            HAPPY_PATH,
					withAuther:       true,
					expectSetDefault: true,
					expectValidate:   true,
				},
			}
			const interval = 1 * time.Second
			// ...
			for _, tt := range tests {
				It(tt.name, Label(string(tt.tType)), func() {
					// Arrange
					// Act
					auther := new(auth.IAuther)
					if tt.withAuther {
						*auther = mockIAuther
					}
					v, err := auth.NewAuthenticator(interval, *auther)
					// Assert
					if tt.expectErr {
						Expect(err).ToNot(BeNil())
						Expect(v).To(BeNil())
					} else {
						Expect(err).To(BeNil())
						Expect(v).ToNot(BeNil())
					}
				})
			}
		})
		// TODO: authenticator.Run tests
	})
})
