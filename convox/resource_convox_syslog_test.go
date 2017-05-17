package convox_test

import (
	"github.com/convox/rack/client"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/mattaitchison/terraform-provider-convox/convox"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ResourceConvoxSyslog", func() {
	convoxClient := &MockClient{}
	var unpacker convox.ClientUnpacker = func(valueGetter convox.ValueGetter, meta interface{}) (convox.Client, error) {
		return convoxClient, nil
	}

	var resourceData *schema.ResourceData

	BeforeEach(func() {
		convoxClient.ResetNoop()
		resourceData = convox.ResourceConvoxSyslog(unpacker).Data(&terraform.InstanceState{
			Attributes: map[string]string{
				"name":     "test",
				"hostname": "logs.foo.com",
				"port":     "12345",
				"scheme":   "tcp+tls",
				"private":  "true",
			},
		})
	})

	Describe("CreateFunc", func() {
		var cut schema.CreateFunc
		BeforeEach(func() {
			cut = convox.ResourceConvoxSyslogCreateFactory(unpacker)
		})

		Describe("#ResourceConvoxSyslogCreateFactory", func() {
			It("should create the CreateFunc", func() {
				Expect(cut).ToNot(BeNil())
			})
		})

		Describe("creating the resource", func() {
			var createdKind string
			var createdOptions map[string]string

			BeforeEach(func() {
				createdKind = ""
				createdOptions = make(map[string]string)
				convoxClient.CreateResourceFunc = func(kind string, options map[string]string) (*client.Resource, error) {
					createdKind = kind
					createdOptions = options
					return &client.Resource{
						Name:   "test",
						Status: "running",
						Exports: map[string]string{
							"URL": "tcp://1.1.1.1:1111",
						},
					}, nil
				}

				Expect(cut(resourceData, resourceData)).To(BeNil())
			})

			It("should set the URL", func() {
				Expect(resourceData.Get("url")).To(Equal("tcp+tls://logs.foo.com:12345"))
			})

			It("should call the convox API with the right kind", func() {
				Expect(createdKind).To(Equal("syslog"))
			})

			It("should call the convox API with the right URL", func() {
				Expect(createdOptions["url"]).To(Equal("tcp+tls://logs.foo.com:12345"))
			})

			It("should call the convox API with the right Private value", func() {
				Expect(createdOptions["private"]).To(Equal("true"))
			})
		})
	})

	Describe("ReadFunc", func() {
		var cut schema.ReadFunc
		BeforeEach(func() {
			cut = convox.ResourceConvoxSyslogReadFactory(unpacker)
		})

		Describe("#ResourceConvoxSyslogReadFactory", func() {
			It("should create the ReadFunc", func() {
				Expect(cut).ToNot(BeNil())
			})
		})

		Describe("Reading state", func() {
			var requestedName string

			BeforeEach(func() {
				convoxClient.GetResourceFunc = func(name string) (*client.Resource, error) {
					requestedName = name
					return &client.Resource{
						Name:   "test",
						Status: "running",
						Exports: map[string]string{
							"URL": "tcp://192.168.1.23:4567",
						},
					}, nil
				}

				Expect(cut(resourceData, resourceData)).To(BeNil())
			})

			It("should ask for the resource by name", func() {
				Expect(requestedName).To(Equal("test"))
			})

			It("should read the URL", func() {
				Expect(resourceData.Get("url")).To(Equal("tcp://192.168.1.23:4567"))
			})
		})
	})
})
