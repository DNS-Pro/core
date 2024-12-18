package app_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestApp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "App Suite")
}

// ...
type TestCaseType string

const (
	HAPPY_PATH TestCaseType = "Happy"
	FAILURE    TestCaseType = "Failuer"
)
