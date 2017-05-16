package convox_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestConvox(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Convox Suite")
}
