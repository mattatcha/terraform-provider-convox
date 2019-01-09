package convox_test

import (
	"github.com/convox/rack/pkg/structs"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/mattatcha/terraform-provider-convox/convox"
	"github.com/stretchr/testify/mock"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ResourceConvoxResourceLink", func() {
	convoxClient := &structs.MockProvider{}
	var unpacker convox.ClientUnpacker = func(valueGetter convox.ValueGetter, meta interface{}) (structs.Provider, error) {
		return convoxClient, nil
	}

	var resourceData *schema.ResourceData

	BeforeEach(func() {
		convoxClient.On("ResourceList").Return(&structs.Resource{
			Status: "running",
		}, nil)

		resourceData = convox.ResourceConvoxResourceLink(unpacker).Data(&terraform.InstanceState{
			Attributes: map[string]string{
				"rack":          "test",
				"app_name":      "test_app",
				"resource_name": "test_resource",
			},
		})

		Expect(convoxClient).ToNot(BeNil())
	})

	Describe("CreateFunc", func() {
		var cut schema.CreateFunc

		BeforeEach(func() {
			cut = convox.ResourceConvoxResourceLinkCreateFactory(unpacker)
		})

		It("should create the func", func() {
			Expect(cut).ToNot(BeNil())
		})

		Describe("creating the resource", func() {
			var calledResourceName string
			var calledAppName string

			BeforeEach(func() {
				calledResourceName = ""
				calledAppName = ""

				convoxClient.On("ResourceLink", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
					calledResourceName = args[0].(string)
					calledAppName = args[1].(string)
				})

				Expect(cut(resourceData, "")).To(BeNil())
			})

			It("should call create with the specified resource name", func() {
				Expect(calledResourceName).To(Equal("test_resource"))
			})

			It("should call create with the specified app name", func() {
				Expect(calledAppName).To(Equal("test_app"))
			})
		})
	})

	Describe("DeleteFunc", func() {
		var cut schema.DeleteFunc

		BeforeEach(func() {
			cut = convox.ResourceConvoxResourceLinkDeleteFactory(unpacker)
		})

		It("should create the func", func() {
			Expect(cut).ToNot(BeNil())
		})

		Describe("deleting the resource", func() {
			var calledResourceName string
			var calledAppName string

			BeforeEach(func() {
				calledResourceName = ""
				calledAppName = ""

				convoxClient.On("ResourceUnlink", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
					calledResourceName = args[0].(string)
					calledAppName = args[1].(string)
				})

				Expect(cut(resourceData, "")).To(BeNil())
			})

			It("should call delete with the specified resource name", func() {
				Expect(calledResourceName).To(Equal("test_resource"))
			})

			It("should call delete with the specified app name", func() {
				Expect(calledAppName).To(Equal("test_app"))
			})
		})
	})
})
