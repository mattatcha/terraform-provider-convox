package convox_test

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/mattatcha/terraform-provider-convox/convox"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Provider", func() {
	var provider terraform.ResourceProvider

	BeforeEach(func() {
		provider = convox.Provider()
	})

	It("should make one", func() {
		Expect(provider).ToNot(BeNil())
	})

	It("should be valid", func() {
		err := provider.(*schema.Provider).InternalValidate()
		Expect(err).To(BeNil())
	})
})
