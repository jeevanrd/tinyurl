package shorturl_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestShorturl(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Shorturl Suite")
}
