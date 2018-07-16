// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/service.proto

package versions

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ServiceInfoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceInfoRequest) Reset()         { *m = ServiceInfoRequest{} }
func (m *ServiceInfoRequest) String() string { return proto.CompactTextString(m) }
func (*ServiceInfoRequest) ProtoMessage()    {}
func (*ServiceInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_19d8b9402dc3447b, []int{0}
}
func (m *ServiceInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceInfoRequest.Unmarshal(m, b)
}
func (m *ServiceInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceInfoRequest.Marshal(b, m, deterministic)
}
func (dst *ServiceInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceInfoRequest.Merge(dst, src)
}
func (m *ServiceInfoRequest) XXX_Size() int {
	return xxx_messageInfo_ServiceInfoRequest.Size(m)
}
func (m *ServiceInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceInfoRequest proto.InternalMessageInfo

type ServiceInfoResponse struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Build                string   `protobuf:"bytes,2,opt,name=build,proto3" json:"build,omitempty"`
	BuiltAt              string   `protobuf:"bytes,3,opt,name=built_at,json=builtAt,proto3" json:"built_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceInfoResponse) Reset()         { *m = ServiceInfoResponse{} }
func (m *ServiceInfoResponse) String() string { return proto.CompactTextString(m) }
func (*ServiceInfoResponse) ProtoMessage()    {}
func (*ServiceInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_19d8b9402dc3447b, []int{1}
}
func (m *ServiceInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceInfoResponse.Unmarshal(m, b)
}
func (m *ServiceInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceInfoResponse.Marshal(b, m, deterministic)
}
func (dst *ServiceInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceInfoResponse.Merge(dst, src)
}
func (m *ServiceInfoResponse) XXX_Size() int {
	return xxx_messageInfo_ServiceInfoResponse.Size(m)
}
func (m *ServiceInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceInfoResponse proto.InternalMessageInfo

func (m *ServiceInfoResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ServiceInfoResponse) GetBuild() string {
	if m != nil {
		return m.Build
	}
	return ""
}

func (m *ServiceInfoResponse) GetBuiltAt() string {
	if m != nil {
		return m.BuiltAt
	}
	return ""
}

func init() {
	proto.RegisterType((*ServiceInfoRequest)(nil), "versions.ServiceInfoRequest")
	proto.RegisterType((*ServiceInfoResponse)(nil), "versions.ServiceInfoResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VersionsClient is the client API for Versions service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VersionsClient interface {
	// Get service version and build info
	Info(ctx context.Context, in *ServiceInfoRequest, opts ...grpc.CallOption) (*ServiceInfoResponse, error)
	// Get version by ID
	GetVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
	// Get applications versions
	GetVersions(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
	// Get applications latest version
	GetLatestVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
	// Create new version
	CreateVersion(ctx context.Context, in *Version, opts ...grpc.CallOption) (*VersionResponse, error)
	// Update existing version
	UpdateVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
	// Delete version
	DeleteVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Get applications list
	GetApplications(ctx context.Context, in *ApplicationRequest, opts ...grpc.CallOption) (*ApplicationResponse, error)
	// Get application details
	GetApplication(ctx context.Context, in *ApplicationRequest, opts ...grpc.CallOption) (*ApplicationResponse, error)
	// Create new application
	CreateApplication(ctx context.Context, in *Application, opts ...grpc.CallOption) (*ApplicationResponse, error)
	// Update existing application
	UpdateApplication(ctx context.Context, in *ApplicationRequest, opts ...grpc.CallOption) (*ApplicationResponse, error)
	// Delete application by ID
	DeleteApplication(ctx context.Context, in *ApplicationRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Get downloads list
	GetDownloads(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (*DownloadResponse, error)
	// Get download details
	GetDownload(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (*DownloadResponse, error)
	// Create new download
	// HTTP Fallback is handled in the service file download.go
	//    handleFallbackDownloads
	//
	// option (google.api.http) = {
	//   post: "/v1/versions/{version_id}/downloads"
	//   body: "*"
	// };
	CreateDownload(ctx context.Context, in *Download, opts ...grpc.CallOption) (*DownloadResponse, error)
	// Get download file
	// HTTP Fallback is handled in the service file download.go
	//    handleFallbackDownloads
	//
	// option (google.api.http) = {
	//   get: "/v1/downloads/{download_id}/file"
	// };
	GetDownloadFile(ctx context.Context, in *Download, opts ...grpc.CallOption) (*Download, error)
	// Update existing download
	UpdateDownload(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (*DownloadResponse, error)
	// Delete download by ID
	DeleteDownload(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type versionsClient struct {
	cc *grpc.ClientConn
}

func NewVersionsClient(cc *grpc.ClientConn) VersionsClient {
	return &versionsClient{cc}
}

func (c *versionsClient) Info(ctx context.Context, in *ServiceInfoRequest, opts ...grpc.CallOption) (*ServiceInfoResponse, error) {
	out := new(ServiceInfoResponse)
	err := c.cc.Invoke(ctx, "/versions.Versions/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionsClient) GetVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/versions.Versions/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionsClient) GetVersions(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/versions.Versions/GetVersions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionsClient) GetLatestVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/versions.Versions/GetLatestVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionsClient) CreateVersion(ctx context.Context, in *Version, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/versions.Versions/CreateVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionsClient) UpdateVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/versions.Versions/UpdateVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionsClient) DeleteVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/versions.Versions/DeleteVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionsClient) GetApplications(ctx context.Context, in *ApplicationRequest, opts ...grpc.CallOption) (*ApplicationResponse, error) {
	out := new(ApplicationResponse)
	err := c.cc.Invoke(ctx, "/versions.Versions/GetApplications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionsClient) GetApplication(ctx context.Context, in *ApplicationRequest, opts ...grpc.CallOption) (*ApplicationResponse, error) {
	out := new(ApplicationResponse)
	err := c.cc.Invoke(ctx, "/versions.Versions/GetApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionsClient) CreateApplication(ctx context.Context, in *Application, opts ...grpc.CallOption) (*ApplicationResponse, error) {
	out := new(ApplicationResponse)
	err := c.cc.Invoke(ctx, "/versions.Versions/CreateApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionsClient) UpdateApplication(ctx context.Context, in *ApplicationRequest, opts ...grpc.CallOption) (*ApplicationResponse, error) {
	out := new(ApplicationResponse)
	err := c.cc.Invoke(ctx, "/versions.Versions/UpdateApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionsClient) DeleteApplication(ctx context.Context, in *ApplicationRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/versions.Versions/DeleteApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionsClient) GetDownloads(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (*DownloadResponse, error) {
	out := new(DownloadResponse)
	err := c.cc.Invoke(ctx, "/versions.Versions/GetDownloads", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionsClient) GetDownload(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (*DownloadResponse, error) {
	out := new(DownloadResponse)
	err := c.cc.Invoke(ctx, "/versions.Versions/GetDownload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionsClient) CreateDownload(ctx context.Context, in *Download, opts ...grpc.CallOption) (*DownloadResponse, error) {
	out := new(DownloadResponse)
	err := c.cc.Invoke(ctx, "/versions.Versions/CreateDownload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionsClient) GetDownloadFile(ctx context.Context, in *Download, opts ...grpc.CallOption) (*Download, error) {
	out := new(Download)
	err := c.cc.Invoke(ctx, "/versions.Versions/GetDownloadFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionsClient) UpdateDownload(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (*DownloadResponse, error) {
	out := new(DownloadResponse)
	err := c.cc.Invoke(ctx, "/versions.Versions/UpdateDownload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionsClient) DeleteDownload(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/versions.Versions/DeleteDownload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VersionsServer is the server API for Versions service.
type VersionsServer interface {
	// Get service version and build info
	Info(context.Context, *ServiceInfoRequest) (*ServiceInfoResponse, error)
	// Get version by ID
	GetVersion(context.Context, *VersionRequest) (*VersionResponse, error)
	// Get applications versions
	GetVersions(context.Context, *VersionRequest) (*VersionResponse, error)
	// Get applications latest version
	GetLatestVersion(context.Context, *VersionRequest) (*VersionResponse, error)
	// Create new version
	CreateVersion(context.Context, *Version) (*VersionResponse, error)
	// Update existing version
	UpdateVersion(context.Context, *VersionRequest) (*VersionResponse, error)
	// Delete version
	DeleteVersion(context.Context, *VersionRequest) (*empty.Empty, error)
	// Get applications list
	GetApplications(context.Context, *ApplicationRequest) (*ApplicationResponse, error)
	// Get application details
	GetApplication(context.Context, *ApplicationRequest) (*ApplicationResponse, error)
	// Create new application
	CreateApplication(context.Context, *Application) (*ApplicationResponse, error)
	// Update existing application
	UpdateApplication(context.Context, *ApplicationRequest) (*ApplicationResponse, error)
	// Delete application by ID
	DeleteApplication(context.Context, *ApplicationRequest) (*empty.Empty, error)
	// Get downloads list
	GetDownloads(context.Context, *DownloadRequest) (*DownloadResponse, error)
	// Get download details
	GetDownload(context.Context, *DownloadRequest) (*DownloadResponse, error)
	// Create new download
	// HTTP Fallback is handled in the service file download.go
	//    handleFallbackDownloads
	//
	// option (google.api.http) = {
	//   post: "/v1/versions/{version_id}/downloads"
	//   body: "*"
	// };
	CreateDownload(context.Context, *Download) (*DownloadResponse, error)
	// Get download file
	// HTTP Fallback is handled in the service file download.go
	//    handleFallbackDownloads
	//
	// option (google.api.http) = {
	//   get: "/v1/downloads/{download_id}/file"
	// };
	GetDownloadFile(context.Context, *Download) (*Download, error)
	// Update existing download
	UpdateDownload(context.Context, *DownloadRequest) (*DownloadResponse, error)
	// Delete download by ID
	DeleteDownload(context.Context, *DownloadRequest) (*empty.Empty, error)
}

func RegisterVersionsServer(s *grpc.Server, srv VersionsServer) {
	s.RegisterService(&_Versions_serviceDesc, srv)
}

func _Versions_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionsServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/versions.Versions/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionsServer).Info(ctx, req.(*ServiceInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Versions_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionsServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/versions.Versions/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionsServer).GetVersion(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Versions_GetVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionsServer).GetVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/versions.Versions/GetVersions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionsServer).GetVersions(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Versions_GetLatestVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionsServer).GetLatestVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/versions.Versions/GetLatestVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionsServer).GetLatestVersion(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Versions_CreateVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Version)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionsServer).CreateVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/versions.Versions/CreateVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionsServer).CreateVersion(ctx, req.(*Version))
	}
	return interceptor(ctx, in, info, handler)
}

func _Versions_UpdateVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionsServer).UpdateVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/versions.Versions/UpdateVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionsServer).UpdateVersion(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Versions_DeleteVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionsServer).DeleteVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/versions.Versions/DeleteVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionsServer).DeleteVersion(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Versions_GetApplications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionsServer).GetApplications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/versions.Versions/GetApplications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionsServer).GetApplications(ctx, req.(*ApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Versions_GetApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionsServer).GetApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/versions.Versions/GetApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionsServer).GetApplication(ctx, req.(*ApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Versions_CreateApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Application)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionsServer).CreateApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/versions.Versions/CreateApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionsServer).CreateApplication(ctx, req.(*Application))
	}
	return interceptor(ctx, in, info, handler)
}

func _Versions_UpdateApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionsServer).UpdateApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/versions.Versions/UpdateApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionsServer).UpdateApplication(ctx, req.(*ApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Versions_DeleteApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionsServer).DeleteApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/versions.Versions/DeleteApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionsServer).DeleteApplication(ctx, req.(*ApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Versions_GetDownloads_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionsServer).GetDownloads(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/versions.Versions/GetDownloads",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionsServer).GetDownloads(ctx, req.(*DownloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Versions_GetDownload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionsServer).GetDownload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/versions.Versions/GetDownload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionsServer).GetDownload(ctx, req.(*DownloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Versions_CreateDownload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Download)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionsServer).CreateDownload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/versions.Versions/CreateDownload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionsServer).CreateDownload(ctx, req.(*Download))
	}
	return interceptor(ctx, in, info, handler)
}

func _Versions_GetDownloadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Download)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionsServer).GetDownloadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/versions.Versions/GetDownloadFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionsServer).GetDownloadFile(ctx, req.(*Download))
	}
	return interceptor(ctx, in, info, handler)
}

func _Versions_UpdateDownload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionsServer).UpdateDownload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/versions.Versions/UpdateDownload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionsServer).UpdateDownload(ctx, req.(*DownloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Versions_DeleteDownload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionsServer).DeleteDownload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/versions.Versions/DeleteDownload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionsServer).DeleteDownload(ctx, req.(*DownloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Versions_serviceDesc = grpc.ServiceDesc{
	ServiceName: "versions.Versions",
	HandlerType: (*VersionsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Info",
			Handler:    _Versions_Info_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _Versions_GetVersion_Handler,
		},
		{
			MethodName: "GetVersions",
			Handler:    _Versions_GetVersions_Handler,
		},
		{
			MethodName: "GetLatestVersion",
			Handler:    _Versions_GetLatestVersion_Handler,
		},
		{
			MethodName: "CreateVersion",
			Handler:    _Versions_CreateVersion_Handler,
		},
		{
			MethodName: "UpdateVersion",
			Handler:    _Versions_UpdateVersion_Handler,
		},
		{
			MethodName: "DeleteVersion",
			Handler:    _Versions_DeleteVersion_Handler,
		},
		{
			MethodName: "GetApplications",
			Handler:    _Versions_GetApplications_Handler,
		},
		{
			MethodName: "GetApplication",
			Handler:    _Versions_GetApplication_Handler,
		},
		{
			MethodName: "CreateApplication",
			Handler:    _Versions_CreateApplication_Handler,
		},
		{
			MethodName: "UpdateApplication",
			Handler:    _Versions_UpdateApplication_Handler,
		},
		{
			MethodName: "DeleteApplication",
			Handler:    _Versions_DeleteApplication_Handler,
		},
		{
			MethodName: "GetDownloads",
			Handler:    _Versions_GetDownloads_Handler,
		},
		{
			MethodName: "GetDownload",
			Handler:    _Versions_GetDownload_Handler,
		},
		{
			MethodName: "CreateDownload",
			Handler:    _Versions_CreateDownload_Handler,
		},
		{
			MethodName: "GetDownloadFile",
			Handler:    _Versions_GetDownloadFile_Handler,
		},
		{
			MethodName: "UpdateDownload",
			Handler:    _Versions_UpdateDownload_Handler,
		},
		{
			MethodName: "DeleteDownload",
			Handler:    _Versions_DeleteDownload_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/service.proto",
}

func init() { proto.RegisterFile("protos/service.proto", fileDescriptor_service_19d8b9402dc3447b) }

var fileDescriptor_service_19d8b9402dc3447b = []byte{
	// 621 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0x9b, 0x42, 0x69, 0x19, 0x9a, 0xd0, 0x0c, 0x09, 0x38, 0x4e, 0x8b, 0x5a, 0x07, 0x2a,
	0x08, 0x95, 0xad, 0xb6, 0x27, 0xe0, 0x54, 0x28, 0x44, 0x48, 0x9c, 0x82, 0xe8, 0xb5, 0x38, 0xf5,
	0xa6, 0x35, 0x32, 0x5e, 0x93, 0xdd, 0x04, 0x50, 0x04, 0x87, 0xbe, 0x02, 0x8f, 0xc6, 0x2b, 0x70,
	0xe0, 0x31, 0xd0, 0x7e, 0xa5, 0x0e, 0x71, 0x3e, 0x94, 0x70, 0xdb, 0x99, 0xc9, 0xfe, 0x7f, 0x3b,
	0xb3, 0xff, 0x75, 0xa0, 0x94, 0x74, 0x28, 0xa7, 0xcc, 0x63, 0xa4, 0xd3, 0x0b, 0xcf, 0x88, 0x2b,
	0x43, 0x5c, 0xeb, 0x91, 0x0e, 0x0b, 0x69, 0xcc, 0xec, 0xcd, 0x73, 0x4a, 0xcf, 0x23, 0xe2, 0xf9,
	0x49, 0xe8, 0xf9, 0x71, 0x4c, 0xb9, 0xcf, 0x45, 0x5e, 0xfd, 0xce, 0xae, 0xea, 0xaa, 0x8c, 0x5a,
	0xdd, 0xb6, 0x47, 0x3e, 0x25, 0xfc, 0x9b, 0x2e, 0x96, 0xb5, 0x74, 0x40, 0xbf, 0xc4, 0x11, 0xf5,
	0x03, 0x9d, 0x36, 0x44, 0x8d, 0xd0, 0x59, 0x4b, 0x67, 0xfd, 0x24, 0x89, 0xc2, 0x33, 0x09, 0x51,
	0x15, 0xa7, 0x04, 0xf8, 0x4e, 0x1d, 0xee, 0x4d, 0xdc, 0xa6, 0x4d, 0xf2, 0xb9, 0x4b, 0x18, 0x77,
	0x3e, 0xc0, 0x9d, 0xa1, 0x2c, 0x4b, 0x68, 0xcc, 0x08, 0x5a, 0xb0, 0xaa, 0x75, 0xad, 0xdc, 0x76,
	0xee, 0xd1, 0xcd, 0xa6, 0x09, 0xb1, 0x04, 0x2b, 0xad, 0x6e, 0x18, 0x05, 0xd6, 0xb2, 0xcc, 0xab,
	0x00, 0x2b, 0xb0, 0x26, 0x16, 0xfc, 0xd4, 0xe7, 0xd6, 0x35, 0xb5, 0x41, 0xc6, 0x47, 0xfc, 0xe0,
	0x4f, 0x01, 0xd6, 0x4e, 0xf4, 0x18, 0xb0, 0x09, 0xd7, 0x05, 0x07, 0x37, 0x5d, 0x33, 0x19, 0x77,
	0xf4, 0x50, 0xf6, 0xd6, 0x98, 0xaa, 0x3a, 0x9c, 0x93, 0xbf, 0xfc, 0xf5, 0xfb, 0xe7, 0xf2, 0x2a,
	0xae, 0x78, 0xa1, 0xd0, 0x0a, 0x00, 0x1a, 0x84, 0x6b, 0x04, 0x5a, 0x57, 0x7b, 0x75, 0xca, 0xa8,
	0x56, 0x32, 0x2a, 0x5a, 0x71, 0x47, 0x2a, 0x56, 0xb1, 0xe2, 0xf5, 0xf6, 0xcd, 0x40, 0x99, 0xd7,
	0xd7, 0xab, 0xd3, 0x30, 0xf8, 0x8e, 0x3d, 0xb8, 0x75, 0x45, 0x61, 0xf3, 0x61, 0x0e, 0x24, 0x66,
	0x0f, 0xeb, 0x02, 0x93, 0xba, 0x21, 0xe6, 0xf5, 0x53, 0x91, 0xc0, 0x0d, 0x0e, 0x81, 0x97, 0x39,
	0xd8, 0x68, 0x10, 0xfe, 0xd6, 0xe7, 0x84, 0x2d, 0xd6, 0xe4, 0x53, 0x49, 0x3f, 0xc4, 0xfd, 0xd9,
	0xe9, 0x5e, 0x24, 0xb1, 0x78, 0x02, 0xf9, 0x97, 0x1d, 0xe2, 0x73, 0x62, 0x0e, 0x50, 0x1c, 0xc1,
	0x4c, 0x22, 0xdf, 0x93, 0xe4, 0xa2, 0xb3, 0x9e, 0x1e, 0xef, 0xb3, 0x5c, 0x1d, 0x3f, 0x42, 0xfe,
	0x7d, 0x12, 0xa4, 0x74, 0xe7, 0x6a, 0xec, 0x81, 0x94, 0xbf, 0x6f, 0x8f, 0xbf, 0x3d, 0xc1, 0x0a,
	0x20, 0x7f, 0x4c, 0x22, 0x32, 0x0b, 0xeb, 0xae, 0xab, 0xde, 0xa3, 0x6b, 0xde, 0xa3, 0xfb, 0x4a,
	0xbc, 0x47, 0x63, 0x93, 0xfa, 0x04, 0x9b, 0x5c, 0xc0, 0xed, 0x06, 0xe1, 0x47, 0xa9, 0xe9, 0xa6,
	0xbd, 0x9e, 0xca, 0x67, 0x78, 0x7d, 0xa8, 0xaa, 0x7b, 0xb3, 0x24, 0x12, 0x71, 0xe3, 0xdf, 0x4b,
	0xc3, 0xaf, 0x50, 0x18, 0x26, 0x2d, 0x06, 0x7a, 0x2c, 0x41, 0x35, 0xdc, 0x99, 0xea, 0x0e, 0x6c,
	0x43, 0x51, 0xb9, 0x21, 0x0d, 0x2f, 0x67, 0xca, 0x4f, 0xa3, 0x56, 0x25, 0xb5, 0xec, 0x8c, 0xb4,
	0x27, 0x6e, 0xec, 0x07, 0x14, 0x95, 0x3b, 0xfe, 0x5b, 0x93, 0x7b, 0x12, 0xb7, 0x6b, 0x4f, 0x6f,
	0x52, 0xf0, 0x39, 0x14, 0x95, 0x63, 0x66, 0xe7, 0x8f, 0x73, 0x8e, 0x9e, 0x6e, 0x7d, 0x86, 0xe9,
	0x76, 0x60, 0xbd, 0x41, 0xf8, 0xb1, 0xfe, 0xd8, 0x33, 0x4c, 0x19, 0xdf, 0x24, 0x0d, 0xcd, 0xce,
	0x2a, 0xe9, 0x56, 0x9f, 0x48, 0xe2, 0x43, 0xac, 0x8d, 0xf5, 0xea, 0xe0, 0x0f, 0x85, 0x61, 0x28,
	0x3f, 0x6e, 0x46, 0x63, 0x5e, 0x64, 0x4d, 0x22, 0xb7, 0xb0, 0x2a, 0x90, 0x03, 0x71, 0xaf, 0x6f,
	0x96, 0xb2, 0xbd, 0x17, 0x50, 0x50, 0xe6, 0x19, 0xd0, 0x70, 0x54, 0x72, 0x22, 0x66, 0x09, 0x9f,
	0xcb, 0x47, 0x66, 0x0a, 0xaf, 0xc3, 0x88, 0x64, 0x8a, 0x64, 0xe4, 0x9c, 0x25, 0xa4, 0x50, 0x50,
	0xae, 0x5a, 0xb4, 0xdd, 0x5d, 0xd9, 0xee, 0xb6, 0x3d, 0xa9, 0x5d, 0x61, 0xa3, 0x0b, 0x28, 0x28,
	0x1b, 0xcd, 0x02, 0x1c, 0x67, 0x20, 0x3d, 0xdb, 0xfa, 0x24, 0x58, 0xeb, 0x86, 0xdc, 0x74, 0xf8,
	0x37, 0x00, 0x00, 0xff, 0xff, 0x70, 0x5e, 0x88, 0xa5, 0x8d, 0x08, 0x00, 0x00,
}
