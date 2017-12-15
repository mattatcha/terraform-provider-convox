package convox_test

import (
	"fmt"

	"github.com/convox/rack/client"
	"github.com/mattaitchison/terraform-provider-convox/convox"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RateLimitRespectingClient", func() {
	var cut convox.Client
	convoxClient := &MockClient{}
	var err error

	BeforeEach(func() {
		cut = convox.NewRateLimitRespectingClient(convoxClient)
	})

	JustBeforeEach(func() {
		_, err = cut.GetResource("foo")
	})

	Describe("when the call is throttled once", func() {
		var calls int

		BeforeEach(func() {
			convoxClient.ResetNoop()
			calls = 0
			convoxClient.GetResourceFunc = func(name string) (*client.Resource, error) {
				calls++
				if calls == 1 {
					return nil, fmt.Errorf("Throttling: Rate exceeded")
				} else {
					return nil, fmt.Errorf("some other error")
				}
			}
		})

		It("should eventually return the other error", func() {
			Expect(err.Error()).To(Equal("some other error"))
		})

		It("should have called twice", func() {
			Expect(calls).To(Equal(2))
		})
	})
})
