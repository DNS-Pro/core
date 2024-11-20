package auth_test

import (
	"testing"

	mockAuth "github.com/DNS-Pro/core/mocks/auth"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAuth(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Auth Suite")
}

type TestCaseType string

const (
	HAPPY_PATH TestCaseType = "Happy"
	FAILURE    TestCaseType = "Failuer"
)

// ...
var iAuthMock *mockAuth.MockIAuthenticator

func resetMock() {
	iAuthMock = mockAuth.NewMockIAuthenticator(GinkgoT())
}
func asserMockCall() {
	iAuthMock.AssertExpectations(GinkgoT())
}
