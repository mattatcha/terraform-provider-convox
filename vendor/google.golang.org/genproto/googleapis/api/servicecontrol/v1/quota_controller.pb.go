// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/servicecontrol/v1/quota_controller.proto

package servicecontrol

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Supported quota modes.
type QuotaOperation_QuotaMode int32

const (
	// Guard against implicit default. Must not be used.
	QuotaOperation_UNSPECIFIED QuotaOperation_QuotaMode = 0
	// For AllocateQuota request, allocates quota for the amount specified in
	// the service configuration or specified using the quota metrics. If the
	// amount is higher than the available quota, allocation error will be
	// returned and no quota will be allocated.
	QuotaOperation_NORMAL QuotaOperation_QuotaMode = 1
	// The operation allocates quota for the amount specified in the service
	// configuration or specified using the quota metrics. If the amount is
	// higher than the available quota, request does not fail but all available
	// quota will be allocated.
	QuotaOperation_BEST_EFFORT QuotaOperation_QuotaMode = 2
	// For AllocateQuota request, only checks if there is enough quota
	// available and does not change the available quota. No lock is placed on
	// the available quota either.
	QuotaOperation_CHECK_ONLY QuotaOperation_QuotaMode = 3
)

var QuotaOperation_QuotaMode_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "NORMAL",
	2: "BEST_EFFORT",
	3: "CHECK_ONLY",
}

var QuotaOperation_QuotaMode_value = map[string]int32{
	"UNSPECIFIED": 0,
	"NORMAL":      1,
	"BEST_EFFORT": 2,
	"CHECK_ONLY":  3,
}

func (x QuotaOperation_QuotaMode) String() string {
	return proto.EnumName(QuotaOperation_QuotaMode_name, int32(x))
}

func (QuotaOperation_QuotaMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4401e348913df3b0, []int{1, 0}
}

// Error codes related to project config validations are deprecated since the
// quota controller methods do not perform these validations. Instead services
// have to call the Check method, without quota_properties field, to perform
// these validations before calling the quota controller methods. These
// methods check only for project deletion to be wipe out compliant.
type QuotaError_Code int32

const (
	// This is never used.
	QuotaError_UNSPECIFIED QuotaError_Code = 0
	// Quota allocation failed.
	// Same as [google.rpc.Code.RESOURCE_EXHAUSTED][].
	QuotaError_RESOURCE_EXHAUSTED QuotaError_Code = 8
	// Consumer cannot access the service because the service requires active
	// billing.
	QuotaError_BILLING_NOT_ACTIVE QuotaError_Code = 107
	// Consumer's project has been marked as deleted (soft deletion).
	QuotaError_PROJECT_DELETED QuotaError_Code = 108
	// Specified API key is invalid.
	QuotaError_API_KEY_INVALID QuotaError_Code = 105
	// Specified API Key has expired.
	QuotaError_API_KEY_EXPIRED QuotaError_Code = 112
)

var QuotaError_Code_name = map[int32]string{
	0:   "UNSPECIFIED",
	8:   "RESOURCE_EXHAUSTED",
	107: "BILLING_NOT_ACTIVE",
	108: "PROJECT_DELETED",
	105: "API_KEY_INVALID",
	112: "API_KEY_EXPIRED",
}

var QuotaError_Code_value = map[string]int32{
	"UNSPECIFIED":        0,
	"RESOURCE_EXHAUSTED": 8,
	"BILLING_NOT_ACTIVE": 107,
	"PROJECT_DELETED":    108,
	"API_KEY_INVALID":    105,
	"API_KEY_EXPIRED":    112,
}

func (x QuotaError_Code) String() string {
	return proto.EnumName(QuotaError_Code_name, int32(x))
}

func (QuotaError_Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4401e348913df3b0, []int{3, 0}
}

// Request message for the AllocateQuota method.
type AllocateQuotaRequest struct {
	// Name of the service as specified in the service configuration. For example,
	// `"pubsub.googleapis.com"`.
	//
	// See [google.api.Service][google.api.Service] for the definition of a
	// service name.
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// Operation that describes the quota allocation.
	AllocateOperation *QuotaOperation `protobuf:"bytes,2,opt,name=allocate_operation,json=allocateOperation,proto3" json:"allocate_operation,omitempty"`
	// Specifies which version of service configuration should be used to process
	// the request. If unspecified or no matching version can be found, the latest
	// one will be used.
	ServiceConfigId      string   `protobuf:"bytes,4,opt,name=service_config_id,json=serviceConfigId,proto3" json:"service_config_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AllocateQuotaRequest) Reset()         { *m = AllocateQuotaRequest{} }
func (m *AllocateQuotaRequest) String() string { return proto.CompactTextString(m) }
func (*AllocateQuotaRequest) ProtoMessage()    {}
func (*AllocateQuotaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4401e348913df3b0, []int{0}
}

func (m *AllocateQuotaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllocateQuotaRequest.Unmarshal(m, b)
}
func (m *AllocateQuotaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllocateQuotaRequest.Marshal(b, m, deterministic)
}
func (m *AllocateQuotaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllocateQuotaRequest.Merge(m, src)
}
func (m *AllocateQuotaRequest) XXX_Size() int {
	return xxx_messageInfo_AllocateQuotaRequest.Size(m)
}
func (m *AllocateQuotaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AllocateQuotaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AllocateQuotaRequest proto.InternalMessageInfo

func (m *AllocateQuotaRequest) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *AllocateQuotaRequest) GetAllocateOperation() *QuotaOperation {
	if m != nil {
		return m.AllocateOperation
	}
	return nil
}

func (m *AllocateQuotaRequest) GetServiceConfigId() string {
	if m != nil {
		return m.ServiceConfigId
	}
	return ""
}

// Represents information regarding a quota operation.
type QuotaOperation struct {
	// Identity of the operation. This is expected to be unique within the scope
	// of the service that generated the operation, and guarantees idempotency in
	// case of retries.
	//
	// UUID version 4 is recommended, though not required. In scenarios where an
	// operation is computed from existing information and an idempotent id is
	// desirable for deduplication purpose, UUID version 5 is recommended. See
	// RFC 4122 for details.
	OperationId string `protobuf:"bytes,1,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	// Fully qualified name of the API method for which this quota operation is
	// requested. This name is used for matching quota rules or metric rules and
	// billing status rules defined in service configuration. This field is not
	// required if the quota operation is performed on non-API resources.
	//
	// Example of an RPC method name:
	//     google.example.library.v1.LibraryService.CreateShelf
	MethodName string `protobuf:"bytes,2,opt,name=method_name,json=methodName,proto3" json:"method_name,omitempty"`
	// Identity of the consumer for whom this quota operation is being performed.
	//
	// This can be in one of the following formats:
	//   project:<project_id>,
	//   project_number:<project_number>,
	//   api_key:<api_key>.
	ConsumerId string `protobuf:"bytes,3,opt,name=consumer_id,json=consumerId,proto3" json:"consumer_id,omitempty"`
	// Labels describing the operation.
	Labels map[string]string `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Represents information about this operation. Each MetricValueSet
	// corresponds to a metric defined in the service configuration.
	// The data type used in the MetricValueSet must agree with
	// the data type specified in the metric definition.
	//
	// Within a single operation, it is not allowed to have more than one
	// MetricValue instances that have the same metric names and identical
	// label value combinations. If a request has such duplicated MetricValue
	// instances, the entire request is rejected with
	// an invalid argument error.
	QuotaMetrics []*MetricValueSet `protobuf:"bytes,5,rep,name=quota_metrics,json=quotaMetrics,proto3" json:"quota_metrics,omitempty"`
	// Quota mode for this operation.
	QuotaMode            QuotaOperation_QuotaMode `protobuf:"varint,6,opt,name=quota_mode,json=quotaMode,proto3,enum=google.api.servicecontrol.v1.QuotaOperation_QuotaMode" json:"quota_mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *QuotaOperation) Reset()         { *m = QuotaOperation{} }
func (m *QuotaOperation) String() string { return proto.CompactTextString(m) }
func (*QuotaOperation) ProtoMessage()    {}
func (*QuotaOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_4401e348913df3b0, []int{1}
}

func (m *QuotaOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuotaOperation.Unmarshal(m, b)
}
func (m *QuotaOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuotaOperation.Marshal(b, m, deterministic)
}
func (m *QuotaOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuotaOperation.Merge(m, src)
}
func (m *QuotaOperation) XXX_Size() int {
	return xxx_messageInfo_QuotaOperation.Size(m)
}
func (m *QuotaOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_QuotaOperation.DiscardUnknown(m)
}

var xxx_messageInfo_QuotaOperation proto.InternalMessageInfo

func (m *QuotaOperation) GetOperationId() string {
	if m != nil {
		return m.OperationId
	}
	return ""
}

func (m *QuotaOperation) GetMethodName() string {
	if m != nil {
		return m.MethodName
	}
	return ""
}

func (m *QuotaOperation) GetConsumerId() string {
	if m != nil {
		return m.ConsumerId
	}
	return ""
}

func (m *QuotaOperation) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *QuotaOperation) GetQuotaMetrics() []*MetricValueSet {
	if m != nil {
		return m.QuotaMetrics
	}
	return nil
}

func (m *QuotaOperation) GetQuotaMode() QuotaOperation_QuotaMode {
	if m != nil {
		return m.QuotaMode
	}
	return QuotaOperation_UNSPECIFIED
}

// Response message for the AllocateQuota method.
type AllocateQuotaResponse struct {
	// The same operation_id value used in the AllocateQuotaRequest. Used for
	// logging and diagnostics purposes.
	OperationId string `protobuf:"bytes,1,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	// Indicates the decision of the allocate.
	AllocateErrors []*QuotaError `protobuf:"bytes,2,rep,name=allocate_errors,json=allocateErrors,proto3" json:"allocate_errors,omitempty"`
	// Quota metrics to indicate the result of allocation. Depending on the
	// request, one or more of the following metrics will be included:
	//
	// 1. Per quota group or per quota metric incremental usage will be specified
	// using the following delta metric :
	//   "serviceruntime.googleapis.com/api/consumer/quota_used_count"
	//
	// 2. The quota limit reached condition will be specified using the following
	// boolean metric :
	//   "serviceruntime.googleapis.com/quota/exceeded"
	QuotaMetrics []*MetricValueSet `protobuf:"bytes,3,rep,name=quota_metrics,json=quotaMetrics,proto3" json:"quota_metrics,omitempty"`
	// ID of the actual config used to process the request.
	ServiceConfigId      string   `protobuf:"bytes,4,opt,name=service_config_id,json=serviceConfigId,proto3" json:"service_config_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AllocateQuotaResponse) Reset()         { *m = AllocateQuotaResponse{} }
func (m *AllocateQuotaResponse) String() string { return proto.CompactTextString(m) }
func (*AllocateQuotaResponse) ProtoMessage()    {}
func (*AllocateQuotaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4401e348913df3b0, []int{2}
}

func (m *AllocateQuotaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllocateQuotaResponse.Unmarshal(m, b)
}
func (m *AllocateQuotaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllocateQuotaResponse.Marshal(b, m, deterministic)
}
func (m *AllocateQuotaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllocateQuotaResponse.Merge(m, src)
}
func (m *AllocateQuotaResponse) XXX_Size() int {
	return xxx_messageInfo_AllocateQuotaResponse.Size(m)
}
func (m *AllocateQuotaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AllocateQuotaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AllocateQuotaResponse proto.InternalMessageInfo

func (m *AllocateQuotaResponse) GetOperationId() string {
	if m != nil {
		return m.OperationId
	}
	return ""
}

func (m *AllocateQuotaResponse) GetAllocateErrors() []*QuotaError {
	if m != nil {
		return m.AllocateErrors
	}
	return nil
}

func (m *AllocateQuotaResponse) GetQuotaMetrics() []*MetricValueSet {
	if m != nil {
		return m.QuotaMetrics
	}
	return nil
}

func (m *AllocateQuotaResponse) GetServiceConfigId() string {
	if m != nil {
		return m.ServiceConfigId
	}
	return ""
}

// Represents error information for
// [QuotaOperation][google.api.servicecontrol.v1.QuotaOperation].
type QuotaError struct {
	// Error code.
	Code QuotaError_Code `protobuf:"varint,1,opt,name=code,proto3,enum=google.api.servicecontrol.v1.QuotaError_Code" json:"code,omitempty"`
	// Subject to whom this error applies. See the specific enum for more details
	// on this field. For example, "clientip:<ip address of client>" or
	// "project:<Google developer project id>".
	Subject string `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	// Free-form text that provides details on the cause of the error.
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuotaError) Reset()         { *m = QuotaError{} }
func (m *QuotaError) String() string { return proto.CompactTextString(m) }
func (*QuotaError) ProtoMessage()    {}
func (*QuotaError) Descriptor() ([]byte, []int) {
	return fileDescriptor_4401e348913df3b0, []int{3}
}

func (m *QuotaError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuotaError.Unmarshal(m, b)
}
func (m *QuotaError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuotaError.Marshal(b, m, deterministic)
}
func (m *QuotaError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuotaError.Merge(m, src)
}
func (m *QuotaError) XXX_Size() int {
	return xxx_messageInfo_QuotaError.Size(m)
}
func (m *QuotaError) XXX_DiscardUnknown() {
	xxx_messageInfo_QuotaError.DiscardUnknown(m)
}

var xxx_messageInfo_QuotaError proto.InternalMessageInfo

func (m *QuotaError) GetCode() QuotaError_Code {
	if m != nil {
		return m.Code
	}
	return QuotaError_UNSPECIFIED
}

func (m *QuotaError) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *QuotaError) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterEnum("google.api.servicecontrol.v1.QuotaOperation_QuotaMode", QuotaOperation_QuotaMode_name, QuotaOperation_QuotaMode_value)
	proto.RegisterEnum("google.api.servicecontrol.v1.QuotaError_Code", QuotaError_Code_name, QuotaError_Code_value)
	proto.RegisterType((*AllocateQuotaRequest)(nil), "google.api.servicecontrol.v1.AllocateQuotaRequest")
	proto.RegisterType((*QuotaOperation)(nil), "google.api.servicecontrol.v1.QuotaOperation")
	proto.RegisterMapType((map[string]string)(nil), "google.api.servicecontrol.v1.QuotaOperation.LabelsEntry")
	proto.RegisterType((*AllocateQuotaResponse)(nil), "google.api.servicecontrol.v1.AllocateQuotaResponse")
	proto.RegisterType((*QuotaError)(nil), "google.api.servicecontrol.v1.QuotaError")
}

func init() {
	proto.RegisterFile("google/api/servicecontrol/v1/quota_controller.proto", fileDescriptor_4401e348913df3b0)
}

var fileDescriptor_4401e348913df3b0 = []byte{
	// 775 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xc1, 0x6e, 0xea, 0x46,
	0x14, 0xed, 0x18, 0x42, 0x9b, 0xeb, 0x04, 0x9c, 0x69, 0x5a, 0x59, 0x28, 0x52, 0x28, 0x2b, 0x1a,
	0xb5, 0x46, 0x21, 0x55, 0x95, 0xa6, 0x2b, 0x30, 0x93, 0xc6, 0x09, 0x01, 0x62, 0x20, 0x4a, 0xda,
	0x85, 0xe5, 0xd8, 0x53, 0xea, 0xc6, 0x78, 0x1c, 0xdb, 0x20, 0x45, 0x55, 0x37, 0x5d, 0x54, 0xaa,
	0xd4, 0x5d, 0xfb, 0x1d, 0xfd, 0x88, 0xfc, 0x42, 0x7f, 0xe1, 0xfd, 0xc3, 0x7b, 0xcb, 0x27, 0x8f,
	0x0d, 0x0f, 0x22, 0xc4, 0x0b, 0x7a, 0x3b, 0xcf, 0xf1, 0x9c, 0x33, 0xf7, 0xde, 0x73, 0xe7, 0x0e,
	0x1c, 0x0d, 0x19, 0x1b, 0xba, 0xb4, 0x6a, 0xfa, 0x4e, 0x35, 0xa4, 0xc1, 0xc4, 0xb1, 0xa8, 0xc5,
	0xbc, 0x28, 0x60, 0x6e, 0x75, 0x72, 0x58, 0x7d, 0x18, 0xb3, 0xc8, 0x34, 0x52, 0xc0, 0xa5, 0x81,
	0xe2, 0x07, 0x2c, 0x62, 0x78, 0x2f, 0x21, 0x29, 0xa6, 0xef, 0x28, 0x8b, 0x24, 0x65, 0x72, 0x58,
	0xdc, 0x9b, 0x93, 0x34, 0x3d, 0x8f, 0x45, 0x66, 0xe4, 0x30, 0x2f, 0x4c, 0xb8, 0xc5, 0xea, 0xca,
	0x03, 0x47, 0x34, 0x0a, 0x1c, 0xcb, 0x98, 0x98, 0xee, 0x98, 0x26, 0x84, 0xf2, 0x13, 0x82, 0xdd,
	0xba, 0xeb, 0x32, 0xcb, 0x8c, 0xe8, 0x55, 0x1c, 0x8f, 0x4e, 0x1f, 0xc6, 0x34, 0x8c, 0xf0, 0x17,
	0xb0, 0x95, 0x0a, 0x18, 0x9e, 0x39, 0xa2, 0x32, 0x2a, 0xa1, 0xca, 0xa6, 0x2e, 0xa6, 0x58, 0xdb,
	0x1c, 0x51, 0xfc, 0x13, 0x60, 0x33, 0xa5, 0x1a, 0xcc, 0xa7, 0x01, 0x8f, 0x44, 0x16, 0x4a, 0xa8,
	0x22, 0xd6, 0xbe, 0x52, 0x56, 0x65, 0xa1, 0xf0, 0xa3, 0x3a, 0x53, 0x8e, 0xbe, 0x33, 0xd5, 0x99,
	0x41, 0xf8, 0x00, 0x76, 0xa6, 0xe7, 0x5b, 0xcc, 0xfb, 0xd9, 0x19, 0x1a, 0x8e, 0x2d, 0x67, 0x79,
	0x10, 0x85, 0xf4, 0x87, 0xca, 0x71, 0xcd, 0x2e, 0xbf, 0xce, 0x40, 0x7e, 0x51, 0x31, 0x0e, 0x7f,
	0x16, 0x52, 0xcc, 0x4c, 0xc3, 0x9f, 0x61, 0x9a, 0x8d, 0xf7, 0x41, 0x1c, 0xd1, 0xe8, 0x17, 0x66,
	0x27, 0x09, 0x0a, 0x7c, 0x07, 0x24, 0x10, 0xcf, 0x6f, 0x1f, 0x44, 0x8b, 0x79, 0xe1, 0x78, 0x44,
	0x83, 0x58, 0x22, 0x93, 0x6c, 0x98, 0x42, 0x9a, 0x8d, 0xbb, 0x90, 0x73, 0xcd, 0x3b, 0xea, 0x86,
	0x72, 0xb6, 0x94, 0xa9, 0x88, 0xb5, 0xe3, 0x75, 0x92, 0x56, 0x5a, 0x9c, 0x4a, 0xbc, 0x28, 0x78,
	0xd4, 0x53, 0x1d, 0x7c, 0x05, 0xdb, 0x49, 0x57, 0x24, 0x56, 0x85, 0xf2, 0x06, 0x17, 0x7e, 0x4f,
	0x35, 0x2f, 0xf9, 0xe6, 0xeb, 0xd8, 0xd6, 0x1e, 0x8d, 0xf4, 0x2d, 0x2e, 0x91, 0x80, 0x21, 0x1e,
	0x00, 0xa4, 0x92, 0xcc, 0xa6, 0x72, 0xae, 0x84, 0x2a, 0xf9, 0xda, 0xb7, 0x6b, 0x05, 0xca, 0x97,
	0x97, 0xcc, 0xa6, 0xfa, 0xe6, 0xc3, 0xf4, 0xb3, 0xf8, 0x1d, 0x88, 0x73, 0x09, 0x60, 0x09, 0x32,
	0xf7, 0xf4, 0x31, 0x2d, 0x73, 0xfc, 0x89, 0x77, 0x61, 0x83, 0x37, 0x5a, 0x5a, 0xd8, 0x64, 0x71,
	0x22, 0x1c, 0xa3, 0xb2, 0x06, 0x9b, 0x33, 0x49, 0x5c, 0x00, 0x71, 0xd0, 0xee, 0x75, 0x89, 0xaa,
	0x9d, 0x6a, 0xa4, 0x29, 0x7d, 0x84, 0x01, 0x72, 0xed, 0x8e, 0x7e, 0x59, 0x6f, 0x49, 0x28, 0xfe,
	0xd9, 0x20, 0xbd, 0xbe, 0x41, 0x4e, 0x4f, 0x3b, 0x7a, 0x5f, 0x12, 0x70, 0x1e, 0x40, 0x3d, 0x23,
	0xea, 0x85, 0xd1, 0x69, 0xb7, 0x6e, 0xa5, 0x4c, 0xf9, 0x6f, 0x01, 0x3e, 0x7b, 0xd6, 0xbe, 0xa1,
	0xcf, 0xbc, 0x90, 0xbe, 0xa4, 0x01, 0xae, 0xa0, 0x30, 0xeb, 0x5f, 0x1a, 0x04, 0x2c, 0x08, 0x65,
	0x81, 0x97, 0xbb, 0xf2, 0x82, 0xf2, 0x90, 0x98, 0xa0, 0xe7, 0xa7, 0x02, 0x7c, 0xb9, 0xc4, 0xbf,
	0xcc, 0x07, 0xfb, 0xb7, 0xce, 0x45, 0xf8, 0x57, 0x00, 0x78, 0x17, 0x1d, 0xae, 0x43, 0xd6, 0x8a,
	0x4d, 0x47, 0xdc, 0xf4, 0xaf, 0x5f, 0x9a, 0x95, 0xa2, 0xc6, 0x5e, 0x73, 0x2a, 0x96, 0xe1, 0xe3,
	0x70, 0x7c, 0xf7, 0x2b, 0xb5, 0xa2, 0xd4, 0xc7, 0xe9, 0x12, 0x97, 0x40, 0xb4, 0x69, 0x68, 0x05,
	0x8e, 0xcf, 0xaf, 0x7d, 0x72, 0x3b, 0xe6, 0xa1, 0xf2, 0x9f, 0x08, 0xb2, 0xea, 0x52, 0x8f, 0x3f,
	0x07, 0xac, 0x93, 0x5e, 0x67, 0xa0, 0xab, 0xc4, 0x20, 0x37, 0x67, 0xf5, 0x41, 0xaf, 0x4f, 0x9a,
	0xd2, 0x27, 0x31, 0xde, 0xd0, 0x5a, 0x2d, 0xad, 0xfd, 0x83, 0xd1, 0xee, 0xf4, 0x8d, 0xba, 0xda,
	0xd7, 0xae, 0x89, 0x74, 0x8f, 0x3f, 0x85, 0x42, 0x57, 0xef, 0x9c, 0x13, 0xb5, 0x6f, 0x34, 0x49,
	0x8b, 0xc4, 0x9b, 0xdd, 0x18, 0xac, 0x77, 0x35, 0xe3, 0x82, 0xdc, 0x1a, 0x5a, 0xfb, 0xba, 0xde,
	0xd2, 0x9a, 0x92, 0x33, 0x0f, 0x92, 0x9b, 0xae, 0xa6, 0x93, 0xa6, 0xe4, 0xd7, 0x9e, 0x10, 0x14,
	0x78, 0x7a, 0xea, 0x6c, 0xd6, 0xe2, 0xff, 0x10, 0x6c, 0x2f, 0x74, 0x0e, 0xae, 0xad, 0xae, 0xcf,
	0xb2, 0x29, 0x59, 0x3c, 0x5a, 0x8b, 0x93, 0xb4, 0x66, 0xf9, 0x9b, 0x3f, 0xfe, 0x7f, 0xf5, 0x8f,
	0xa0, 0x94, 0xbf, 0x8c, 0x67, 0x72, 0x4a, 0x0a, 0xab, 0xbf, 0xcd, 0x8f, 0xdb, 0xdf, 0x4f, 0xcc,
	0x79, 0xea, 0x09, 0x3a, 0x68, 0xfc, 0x85, 0xa0, 0x64, 0xb1, 0xd1, 0xca, 0x03, 0x1b, 0xbb, 0xcf,
	0xd2, 0xec, 0xc6, 0x43, 0xbe, 0x8b, 0x7e, 0x3c, 0x4f, 0x59, 0x43, 0xe6, 0x9a, 0xde, 0x50, 0x61,
	0xc1, 0xb0, 0x3a, 0xa4, 0x1e, 0x7f, 0x02, 0xd2, 0x27, 0xc3, 0xf4, 0x9d, 0x70, 0xf9, 0xb3, 0xf1,
	0xfd, 0x22, 0xf2, 0x06, 0xa1, 0xbb, 0x1c, 0x67, 0x1e, 0xbd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xbb,
	0x98, 0x03, 0x4f, 0xe0, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QuotaControllerClient is the client API for QuotaController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QuotaControllerClient interface {
	// Attempts to allocate quota for the specified consumer. It should be called
	// before the operation is executed.
	//
	// This method requires the `servicemanagement.services.quota`
	// permission on the specified service. For more information, see
	// [Cloud IAM](https://cloud.google.com/iam).
	//
	// **NOTE:** The client **must** fail-open on server errors `INTERNAL`,
	// `UNKNOWN`, `DEADLINE_EXCEEDED`, and `UNAVAILABLE`. To ensure system
	// reliability, the server may inject these errors to prohibit any hard
	// dependency on the quota functionality.
	AllocateQuota(ctx context.Context, in *AllocateQuotaRequest, opts ...grpc.CallOption) (*AllocateQuotaResponse, error)
}

type quotaControllerClient struct {
	cc *grpc.ClientConn
}

func NewQuotaControllerClient(cc *grpc.ClientConn) QuotaControllerClient {
	return &quotaControllerClient{cc}
}

func (c *quotaControllerClient) AllocateQuota(ctx context.Context, in *AllocateQuotaRequest, opts ...grpc.CallOption) (*AllocateQuotaResponse, error) {
	out := new(AllocateQuotaResponse)
	err := c.cc.Invoke(ctx, "/google.api.servicecontrol.v1.QuotaController/AllocateQuota", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuotaControllerServer is the server API for QuotaController service.
type QuotaControllerServer interface {
	// Attempts to allocate quota for the specified consumer. It should be called
	// before the operation is executed.
	//
	// This method requires the `servicemanagement.services.quota`
	// permission on the specified service. For more information, see
	// [Cloud IAM](https://cloud.google.com/iam).
	//
	// **NOTE:** The client **must** fail-open on server errors `INTERNAL`,
	// `UNKNOWN`, `DEADLINE_EXCEEDED`, and `UNAVAILABLE`. To ensure system
	// reliability, the server may inject these errors to prohibit any hard
	// dependency on the quota functionality.
	AllocateQuota(context.Context, *AllocateQuotaRequest) (*AllocateQuotaResponse, error)
}

func RegisterQuotaControllerServer(s *grpc.Server, srv QuotaControllerServer) {
	s.RegisterService(&_QuotaController_serviceDesc, srv)
}

func _QuotaController_AllocateQuota_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllocateQuotaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuotaControllerServer).AllocateQuota(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.api.servicecontrol.v1.QuotaController/AllocateQuota",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuotaControllerServer).AllocateQuota(ctx, req.(*AllocateQuotaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QuotaController_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.api.servicecontrol.v1.QuotaController",
	HandlerType: (*QuotaControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AllocateQuota",
			Handler:    _QuotaController_AllocateQuota_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/api/servicecontrol/v1/quota_controller.proto",
}
