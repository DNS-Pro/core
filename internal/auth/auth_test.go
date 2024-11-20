package auth_test

import (
	"time"

	"github.com/DNS-Pro/core/internal/auth"
	mockAuth "github.com/DNS-Pro/core/mocks/auth"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("Auth", func() {
	var iAuthMock *mockAuth.MockIAuthenticator

	resetMock := func() {
		iAuthMock = mockAuth.NewMockIAuthenticator(GinkgoT())
	}
	asserMockCall := func() {
		iAuthMock.AssertExpectations(GinkgoT())
	}

	Describe("AuthType", Label("AuthType"), func() {
		Describe("fromAuthenticator", func() {
			tests := []struct {
				name         string
				tType        TestCaseType
				input        auth.IAuther  // input to inference type from
				expectOutput auth.AuthType // expected inferenced type
			}{
				{
					name:         "Successfully infer AUTH_NONE",
					tType:        HAPPY_PATH,
					input:        nil,
					expectOutput: auth.AUTH_NONE,
				},
				{
					name:         "Successfully infer AUTH_HTTP",
					tType:        HAPPY_PATH,
					input:        &auth.HttpAuthenticator{},
					expectOutput: auth.AUTH_HTTP,
				},
			}
			for _, tt := range tests {
				It(tt.name, Label(string(tt.tType)), func() {
					v := new(auth.AuthType)
					v.FromAuthenticator(tt.input)
					Expect(*v).To(Equal(tt.expectOutput))
				})
			}
		})
	})
	Describe("Authenticator", Label("Authenticator"), func() {
		Describe("NewAuthenticator", Label("NewAuthenticator"), func() {
			type testCase struct {
				name              string
				tType             TestCaseType
				withAuther        bool // wether or not to create with auther
				expectSetDefault  bool // expect calling authenticator SetDefault method
				expectSetBaseAuth bool // expect calling authenticator SetBaseAuth method
				expectValidate    bool // expect calling authenticator Validate method
				expectErr         bool // expect error
			}
			// ...
			assertIAuth_Validate := func(tc testCase) {
				if tc.expectValidate {
					iAuthMock.EXPECT().Validate().Return(nil)
				}
			}
			assertIAuth_SetDefaults := func(tc testCase) {
				if tc.expectSetDefault {
					iAuthMock.EXPECT().SetDefaults().Return(nil)
				}
			}
			assertIAuth_SetBaseAuth := func(tc testCase) {
				if tc.expectSetBaseAuth {
					iAuthMock.EXPECT().SetBaseAuth(mock.Anything).Return()
				}
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
					name:              "Successfully Create authenticator",
					tType:             HAPPY_PATH,
					withAuther:        true,
					expectSetDefault:  true,
					expectSetBaseAuth: true,
					expectValidate:    true,
				},
			}
			const interval = 1 * time.Second
			// ...
			for _, tt := range tests {
				It(tt.name, Label(string(tt.tType)), func() {
					// Arrange
					assertIAuth_SetDefaults(tt)
					assertIAuth_SetBaseAuth(tt)
					assertIAuth_Validate(tt)
					// Act
					auther := new(auth.IAuther)
					if tt.withAuther {
						*auther = iAuthMock
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
	})
})
