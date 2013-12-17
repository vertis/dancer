package dancer_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDancer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dancer Suite")
}
