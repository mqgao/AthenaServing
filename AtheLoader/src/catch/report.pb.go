// Code generated by protoc-gen-go. DO NOT EDIT.
// source: report.proto

package catch

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type MetaData struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Format               string   `protobuf:"bytes,2,opt,name=format,proto3" json:"format,omitempty"`
	Encoding             string   `protobuf:"bytes,3,opt,name=encoding,proto3" json:"encoding,omitempty"`
	Data                 []byte   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetaData) Reset()         { *m = MetaData{} }
func (m *MetaData) String() string { return proto.CompactTextString(m) }
func (*MetaData) ProtoMessage()    {}
func (*MetaData) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eedb623aa6ca98c, []int{0}
}

func (m *MetaData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetaData.Unmarshal(m, b)
}
func (m *MetaData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetaData.Marshal(b, m, deterministic)
}
func (m *MetaData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetaData.Merge(m, src)
}
func (m *MetaData) XXX_Size() int {
	return xxx_messageInfo_MetaData.Size(m)
}
func (m *MetaData) XXX_DiscardUnknown() {
	xxx_messageInfo_MetaData.DiscardUnknown(m)
}

var xxx_messageInfo_MetaData proto.InternalMessageInfo

func (m *MetaData) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *MetaData) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

func (m *MetaData) GetEncoding() string {
	if m != nil {
		return m.Encoding
	}
	return ""
}

func (m *MetaData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type MetaDoubt struct {
	Score                int32             `protobuf:"varint,1,opt,name=score,proto3" json:"score,omitempty"`
	Sid                  string            `protobuf:"bytes,2,opt,name=sid,proto3" json:"sid,omitempty"`
	Params               map[string]string `protobuf:"bytes,3,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Doubt                map[string]string `protobuf:"bytes,4,rep,name=doubt,proto3" json:"doubt,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Datalist             []*MetaData       `protobuf:"bytes,5,rep,name=datalist,proto3" json:"datalist,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MetaDoubt) Reset()         { *m = MetaDoubt{} }
func (m *MetaDoubt) String() string { return proto.CompactTextString(m) }
func (*MetaDoubt) ProtoMessage()    {}
func (*MetaDoubt) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eedb623aa6ca98c, []int{1}
}

func (m *MetaDoubt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetaDoubt.Unmarshal(m, b)
}
func (m *MetaDoubt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetaDoubt.Marshal(b, m, deterministic)
}
func (m *MetaDoubt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetaDoubt.Merge(m, src)
}
func (m *MetaDoubt) XXX_Size() int {
	return xxx_messageInfo_MetaDoubt.Size(m)
}
func (m *MetaDoubt) XXX_DiscardUnknown() {
	xxx_messageInfo_MetaDoubt.DiscardUnknown(m)
}

var xxx_messageInfo_MetaDoubt proto.InternalMessageInfo

func (m *MetaDoubt) GetScore() int32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *MetaDoubt) GetSid() string {
	if m != nil {
		return m.Sid
	}
	return ""
}

func (m *MetaDoubt) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *MetaDoubt) GetDoubt() map[string]string {
	if m != nil {
		return m.Doubt
	}
	return nil
}

func (m *MetaDoubt) GetDatalist() []*MetaData {
	if m != nil {
		return m.Datalist
	}
	return nil
}

type MetaTrust struct {
	Sid                  string            `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"`
	Params               map[string]string `protobuf:"bytes,2,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MetaTrust) Reset()         { *m = MetaTrust{} }
func (m *MetaTrust) String() string { return proto.CompactTextString(m) }
func (*MetaTrust) ProtoMessage()    {}
func (*MetaTrust) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eedb623aa6ca98c, []int{2}
}

func (m *MetaTrust) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetaTrust.Unmarshal(m, b)
}
func (m *MetaTrust) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetaTrust.Marshal(b, m, deterministic)
}
func (m *MetaTrust) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetaTrust.Merge(m, src)
}
func (m *MetaTrust) XXX_Size() int {
	return xxx_messageInfo_MetaTrust.Size(m)
}
func (m *MetaTrust) XXX_DiscardUnknown() {
	xxx_messageInfo_MetaTrust.DiscardUnknown(m)
}

var xxx_messageInfo_MetaTrust proto.InternalMessageInfo

func (m *MetaTrust) GetSid() string {
	if m != nil {
		return m.Sid
	}
	return ""
}

func (m *MetaTrust) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

type CatchReport struct {
	Event                string       `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Service              string       `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	Address              string       `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Summary              string       `protobuf:"bytes,4,opt,name=summary,proto3" json:"summary,omitempty"`
	Pid                  string       `protobuf:"bytes,5,opt,name=pid,proto3" json:"pid,omitempty"`
	Stackc               []byte       `protobuf:"bytes,6,opt,name=stackc,proto3" json:"stackc,omitempty"`
	Stackgo              []byte       `protobuf:"bytes,7,opt,name=stackgo,proto3" json:"stackgo,omitempty"`
	Doubtlist            []*MetaDoubt `protobuf:"bytes,8,rep,name=doubtlist,proto3" json:"doubtlist,omitempty"`
	Trustlist            []*MetaTrust `protobuf:"bytes,9,rep,name=trustlist,proto3" json:"trustlist,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CatchReport) Reset()         { *m = CatchReport{} }
func (m *CatchReport) String() string { return proto.CompactTextString(m) }
func (*CatchReport) ProtoMessage()    {}
func (*CatchReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eedb623aa6ca98c, []int{3}
}

func (m *CatchReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CatchReport.Unmarshal(m, b)
}
func (m *CatchReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CatchReport.Marshal(b, m, deterministic)
}
func (m *CatchReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CatchReport.Merge(m, src)
}
func (m *CatchReport) XXX_Size() int {
	return xxx_messageInfo_CatchReport.Size(m)
}
func (m *CatchReport) XXX_DiscardUnknown() {
	xxx_messageInfo_CatchReport.DiscardUnknown(m)
}

var xxx_messageInfo_CatchReport proto.InternalMessageInfo

func (m *CatchReport) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *CatchReport) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *CatchReport) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *CatchReport) GetSummary() string {
	if m != nil {
		return m.Summary
	}
	return ""
}

func (m *CatchReport) GetPid() string {
	if m != nil {
		return m.Pid
	}
	return ""
}

func (m *CatchReport) GetStackc() []byte {
	if m != nil {
		return m.Stackc
	}
	return nil
}

func (m *CatchReport) GetStackgo() []byte {
	if m != nil {
		return m.Stackgo
	}
	return nil
}

func (m *CatchReport) GetDoubtlist() []*MetaDoubt {
	if m != nil {
		return m.Doubtlist
	}
	return nil
}

func (m *CatchReport) GetTrustlist() []*MetaTrust {
	if m != nil {
		return m.Trustlist
	}
	return nil
}

func init() {
	proto.RegisterType((*MetaData)(nil), "catch.MetaData")
	proto.RegisterType((*MetaDoubt)(nil), "catch.MetaDoubt")
	proto.RegisterMapType((map[string]string)(nil), "catch.MetaDoubt.DoubtEntry")
	proto.RegisterMapType((map[string]string)(nil), "catch.MetaDoubt.ParamsEntry")
	proto.RegisterType((*MetaTrust)(nil), "catch.MetaTrust")
	proto.RegisterMapType((map[string]string)(nil), "catch.MetaTrust.ParamsEntry")
	proto.RegisterType((*CatchReport)(nil), "catch.CatchReport")
}

func init() { proto.RegisterFile("report.proto", fileDescriptor_3eedb623aa6ca98c) }

var fileDescriptor_3eedb623aa6ca98c = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xcd, 0x8a, 0xdb, 0x30,
	0x10, 0xc6, 0x4e, 0x9c, 0x8d, 0x27, 0x0b, 0x5d, 0x44, 0x29, 0x22, 0xed, 0x21, 0xf8, 0x14, 0x28,
	0x18, 0xfa, 0x73, 0xd8, 0xf6, 0xda, 0xf6, 0x58, 0x28, 0xa2, 0x2f, 0xa0, 0xb5, 0xb4, 0x5b, 0xb3,
	0x6b, 0xcb, 0x48, 0xe3, 0x80, 0x5f, 0xa2, 0xb7, 0xbe, 0x4e, 0x9f, 0xad, 0xcc, 0x48, 0x4e, 0x96,
	0x66, 0x2f, 0xa5, 0x97, 0x30, 0x5f, 0xe6, 0xfb, 0xe6, 0xe7, 0xd3, 0x18, 0x2e, 0xbd, 0x1d, 0x9c,
	0xc7, 0x7a, 0xf0, 0x0e, 0x9d, 0x28, 0x1a, 0x8d, 0xcd, 0x8f, 0xea, 0x16, 0xd6, 0x5f, 0x2d, 0xea,
	0xcf, 0x1a, 0xb5, 0x10, 0xb0, 0xc4, 0x69, 0xb0, 0x32, 0xdb, 0x65, 0xfb, 0x52, 0x71, 0x2c, 0x5e,
	0xc0, 0xea, 0xd6, 0xf9, 0x4e, 0xa3, 0xcc, 0xf9, 0xdf, 0x84, 0xc4, 0x16, 0xd6, 0xb6, 0x6f, 0x9c,
	0x69, 0xfb, 0x3b, 0xb9, 0xe0, 0xcc, 0x11, 0x53, 0x1d, 0xa3, 0x51, 0xcb, 0xe5, 0x2e, 0xdb, 0x5f,
	0x2a, 0x8e, 0xab, 0xdf, 0x39, 0x94, 0xdc, 0xc8, 0x8d, 0x37, 0x28, 0x9e, 0x43, 0x11, 0x1a, 0xe7,
	0x63, 0xab, 0x42, 0x45, 0x20, 0xae, 0x60, 0x11, 0x5a, 0x93, 0x1a, 0x51, 0x28, 0xde, 0xc3, 0x6a,
	0xd0, 0x5e, 0x77, 0x41, 0x2e, 0x76, 0x8b, 0xfd, 0xe6, 0xed, 0xab, 0x9a, 0xa7, 0xae, 0x8f, 0x95,
	0xea, 0x6f, 0x9c, 0xfe, 0xd2, 0xa3, 0x9f, 0x54, 0xe2, 0x8a, 0x37, 0x50, 0x18, 0x4a, 0xca, 0x25,
	0x8b, 0x5e, 0x9e, 0x89, 0xf8, 0x37, 0x6a, 0x22, 0x53, 0xbc, 0x86, 0x35, 0x8d, 0xf9, 0xd0, 0x06,
	0x94, 0x05, 0xab, 0x9e, 0x3d, 0x56, 0x69, 0xd4, 0xea, 0x48, 0xd8, 0x7e, 0x80, 0xcd, 0xa3, 0xb6,
	0x34, 0xf6, 0xbd, 0x9d, 0x92, 0x6b, 0x14, 0xd2, 0x7a, 0x07, 0xfd, 0x30, 0xda, 0xb4, 0x4a, 0x04,
	0x1f, 0xf3, 0xeb, 0x6c, 0x7b, 0x0d, 0x70, 0x6a, 0xfe, 0x2f, 0xca, 0xea, 0x67, 0x16, 0x0d, 0xfc,
	0xee, 0xc7, 0x80, 0xb3, 0x55, 0xd9, 0x53, 0x56, 0xe5, 0x67, 0x56, 0xb1, 0xe6, 0x29, 0xab, 0xfe,
	0x63, 0x95, 0xea, 0x57, 0x0e, 0x9b, 0x4f, 0xd4, 0x42, 0xf1, 0x59, 0x11, 0xd3, 0x1e, 0x6c, 0x8f,
	0x49, 0x1d, 0x81, 0x90, 0x70, 0x11, 0xac, 0x3f, 0xb4, 0xcd, 0x5c, 0x61, 0x86, 0x94, 0xd1, 0xc6,
	0x78, 0x1b, 0x42, 0x3a, 0xa0, 0x19, 0xb2, 0x66, 0xec, 0x3a, 0xed, 0x27, 0x3e, 0x21, 0xd2, 0x44,
	0x48, 0xf3, 0x0d, 0xad, 0x91, 0x45, 0x9c, 0x6f, 0x68, 0x0d, 0xdd, 0x67, 0x40, 0xdd, 0xdc, 0x37,
	0x72, 0xc5, 0xd7, 0x96, 0x10, 0xd7, 0xa0, 0xe8, 0xce, 0xc9, 0x0b, 0x4e, 0xcc, 0x50, 0xd4, 0x50,
	0xf2, 0x9b, 0xf3, 0x5b, 0xaf, 0xd9, 0xab, 0xab, 0xbf, 0x2f, 0x44, 0x9d, 0x28, 0xc4, 0x47, 0xf2,
	0x8f, 0xf9, 0xe5, 0x19, 0x9f, 0xbd, 0x55, 0x27, 0xca, 0xcd, 0x8a, 0xbf, 0xaf, 0x77, 0x7f, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x8a, 0x3f, 0x82, 0x34, 0x6f, 0x03, 0x00, 0x00,
}
