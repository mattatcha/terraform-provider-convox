package convox_test

import (
	"github.com/convox/rack/client"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/mattaitchison/terraform-provider-convox/convox"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ResourceConvoxResourceLink", func() {
	convoxClient := &MockClient{}
	var unpacker convox.ClientUnpacker = func(valueGetter convox.ValueGetter, meta interface{}) (convox.Client, error) {
		return convoxClient, nil
	}

	var resourceData *schema.ResourceData

	BeforeEach(func() {
		convoxClient.ResetNoop()
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
				convoxClient.CreateLinkFunc = func(resource string, app string) (*client.Resource, error) {
					calledResourceName = resource
					calledAppName = app

					return nil, nil
				}

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
				convoxClient.DeleteLinkFunc = func(resource string, app string) (*client.Resource, error) {
					calledResourceName = resource
					calledAppName = app

					return nil, nil
				}

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
