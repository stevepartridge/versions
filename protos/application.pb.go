// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/application.proto

package versions

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Application struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	CreatedAt            string   `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Application) Reset()         { *m = Application{} }
func (m *Application) String() string { return proto.CompactTextString(m) }
func (*Application) ProtoMessage()    {}
func (*Application) Descriptor() ([]byte, []int) {
	return fileDescriptor_application_b24b4f22d5f4a4ca, []int{0}
}
func (m *Application) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Application.Unmarshal(m, b)
}
func (m *Application) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Application.Marshal(b, m, deterministic)
}
func (dst *Application) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Application.Merge(dst, src)
}
func (m *Application) XXX_Size() int {
	return xxx_messageInfo_Application.Size(m)
}
func (m *Application) XXX_DiscardUnknown() {
	xxx_messageInfo_Application.DiscardUnknown(m)
}

var xxx_messageInfo_Application proto.InternalMessageInfo

func (m *Application) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Application) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Application) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Application) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

type ApplicationRequest struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ApplicationId        string       `protobuf:"bytes,2,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	Application          *Application `protobuf:"bytes,3,opt,name=application,proto3" json:"application,omitempty"`
	Limit                int32        `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int32        `protobuf:"varint,5,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ApplicationRequest) Reset()         { *m = ApplicationRequest{} }
func (m *ApplicationRequest) String() string { return proto.CompactTextString(m) }
func (*ApplicationRequest) ProtoMessage()    {}
func (*ApplicationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_application_b24b4f22d5f4a4ca, []int{1}
}
func (m *ApplicationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplicationRequest.Unmarshal(m, b)
}
func (m *ApplicationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplicationRequest.Marshal(b, m, deterministic)
}
func (dst *ApplicationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplicationRequest.Merge(dst, src)
}
func (m *ApplicationRequest) XXX_Size() int {
	return xxx_messageInfo_ApplicationRequest.Size(m)
}
func (m *ApplicationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplicationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ApplicationRequest proto.InternalMessageInfo

func (m *ApplicationRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ApplicationRequest) GetApplicationId() string {
	if m != nil {
		return m.ApplicationId
	}
	return ""
}

func (m *ApplicationRequest) GetApplication() *Application {
	if m != nil {
		return m.Application
	}
	return nil
}

func (m *ApplicationRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ApplicationRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type ApplicationResponse struct {
	Application          *Application   `protobuf:"bytes,1,opt,name=application,proto3" json:"application,omitempty"`
	Applications         []*Application `protobuf:"bytes,2,rep,name=applications,proto3" json:"applications,omitempty"`
	Limit                int32          `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int32          `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ApplicationResponse) Reset()         { *m = ApplicationResponse{} }
func (m *ApplicationResponse) String() string { return proto.CompactTextString(m) }
func (*ApplicationResponse) ProtoMessage()    {}
func (*ApplicationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_application_b24b4f22d5f4a4ca, []int{2}
}
func (m *ApplicationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplicationResponse.Unmarshal(m, b)
}
func (m *ApplicationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplicationResponse.Marshal(b, m, deterministic)
}
func (dst *ApplicationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplicationResponse.Merge(dst, src)
}
func (m *ApplicationResponse) XXX_Size() int {
	return xxx_messageInfo_ApplicationResponse.Size(m)
}
func (m *ApplicationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplicationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ApplicationResponse proto.InternalMessageInfo

func (m *ApplicationResponse) GetApplication() *Application {
	if m != nil {
		return m.Application
	}
	return nil
}

func (m *ApplicationResponse) GetApplications() []*Application {
	if m != nil {
		return m.Applications
	}
	return nil
}

func (m *ApplicationResponse) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ApplicationResponse) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func init() {
	proto.RegisterType((*Application)(nil), "versions.Application")
	proto.RegisterType((*ApplicationRequest)(nil), "versions.ApplicationRequest")
	proto.RegisterType((*ApplicationResponse)(nil), "versions.ApplicationResponse")
}

func init() {
	proto.RegisterFile("protos/application.proto", fileDescriptor_application_b24b4f22d5f4a4ca)
}

var fileDescriptor_application_b24b4f22d5f4a4ca = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcd, 0x4a, 0x03, 0x31,
	0x14, 0x85, 0xc9, 0xfc, 0x14, 0x7b, 0x47, 0xbb, 0x88, 0x3f, 0x04, 0x51, 0x28, 0x03, 0x82, 0xab,
	0x19, 0xd0, 0x85, 0xb8, 0x9c, 0xa5, 0xdb, 0x79, 0x81, 0x12, 0x9b, 0xb4, 0x04, 0xa6, 0xb9, 0x71,
	0x72, 0xeb, 0x83, 0xb9, 0xf1, 0xf5, 0xc4, 0xcc, 0x94, 0xa6, 0xfe, 0x40, 0x77, 0xb9, 0xe7, 0x9c,
	0xcb, 0xf9, 0x2e, 0x01, 0xe1, 0x7a, 0x24, 0xf4, 0xb5, 0x74, 0xae, 0x33, 0x4b, 0x49, 0x06, 0x6d,
	0x15, 0x24, 0x7e, 0xf2, 0xae, 0x7b, 0x6f, 0xd0, 0xfa, 0xeb, 0x9b, 0x35, 0xe2, 0xba, 0xd3, 0xb5,
	0x74, 0xa6, 0x96, 0xd6, 0x22, 0x85, 0x98, 0x1f, 0x72, 0x25, 0x42, 0xd1, 0xec, 0x97, 0xf9, 0x0c,
	0x12, 0xa3, 0x04, 0x9b, 0xb3, 0xfb, 0x69, 0x9b, 0x18, 0xc5, 0x39, 0x64, 0x56, 0x6e, 0xb4, 0x48,
	0x82, 0x12, 0xde, 0xfc, 0x16, 0x60, 0xeb, 0x94, 0x24, 0xad, 0x16, 0x92, 0x44, 0x1a, 0x9c, 0xe9,
	0xa8, 0x34, 0xf4, 0x6d, 0x2f, 0x7b, 0xbd, 0xb3, 0xb3, 0xc1, 0x1e, 0x95, 0x86, 0xca, 0x0f, 0x06,
	0x3c, 0x6a, 0x6c, 0xf5, 0xdb, 0x56, 0x7b, 0xfa, 0x55, 0x7c, 0x07, 0xb3, 0xe8, 0xa8, 0x85, 0x51,
	0x23, 0xc2, 0x59, 0xa4, 0xbe, 0x28, 0xfe, 0x04, 0x45, 0x24, 0x04, 0x98, 0xe2, 0xe1, 0xb2, 0xda,
	0x1d, 0x5f, 0xc5, 0x4d, 0x71, 0x92, 0x5f, 0x40, 0xde, 0x99, 0x8d, 0x19, 0x00, 0xf3, 0x76, 0x18,
	0xf8, 0x15, 0x4c, 0x70, 0xb5, 0xf2, 0x9a, 0x44, 0x1e, 0xe4, 0x71, 0x2a, 0x3f, 0x19, 0x9c, 0x1f,
	0x40, 0x7b, 0x87, 0xd6, 0xeb, 0x9f, 0xf5, 0xec, 0xe8, 0xfa, 0x67, 0x38, 0x8d, 0x46, 0x2f, 0x92,
	0x79, 0xfa, 0xff, 0xe6, 0x41, 0x74, 0x4f, 0x9e, 0xfe, 0x4d, 0x9e, 0xc5, 0xe4, 0xaf, 0x93, 0xf0,
	0xcd, 0x8f, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xab, 0xfa, 0x67, 0x96, 0x2a, 0x02, 0x00, 0x00,
}
