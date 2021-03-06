// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/matcher/value.proto

package envoy_type_matcher

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// Specifies the way to match a ProtobufWkt::Value. Primitive values and ListValue are supported.
// StructValue is not supported and is always not matched.
// [#next-free-field: 7]
type ValueMatcher struct {
	// Specifies how to match a value.
	//
	// Types that are valid to be assigned to MatchPattern:
	//	*ValueMatcher_NullMatch_
	//	*ValueMatcher_DoubleMatch
	//	*ValueMatcher_StringMatch
	//	*ValueMatcher_BoolMatch
	//	*ValueMatcher_PresentMatch
	//	*ValueMatcher_ListMatch
	MatchPattern         isValueMatcher_MatchPattern `protobuf_oneof:"match_pattern"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ValueMatcher) Reset()         { *m = ValueMatcher{} }
func (m *ValueMatcher) String() string { return proto.CompactTextString(m) }
func (*ValueMatcher) ProtoMessage()    {}
func (*ValueMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_145b36501d266253, []int{0}
}

func (m *ValueMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValueMatcher.Unmarshal(m, b)
}
func (m *ValueMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValueMatcher.Marshal(b, m, deterministic)
}
func (m *ValueMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValueMatcher.Merge(m, src)
}
func (m *ValueMatcher) XXX_Size() int {
	return xxx_messageInfo_ValueMatcher.Size(m)
}
func (m *ValueMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_ValueMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_ValueMatcher proto.InternalMessageInfo

type isValueMatcher_MatchPattern interface {
	isValueMatcher_MatchPattern()
}

type ValueMatcher_NullMatch_ struct {
	NullMatch *ValueMatcher_NullMatch `protobuf:"bytes,1,opt,name=null_match,json=nullMatch,proto3,oneof"`
}

type ValueMatcher_DoubleMatch struct {
	DoubleMatch *DoubleMatcher `protobuf:"bytes,2,opt,name=double_match,json=doubleMatch,proto3,oneof"`
}

type ValueMatcher_StringMatch struct {
	StringMatch *StringMatcher `protobuf:"bytes,3,opt,name=string_match,json=stringMatch,proto3,oneof"`
}

type ValueMatcher_BoolMatch struct {
	BoolMatch bool `protobuf:"varint,4,opt,name=bool_match,json=boolMatch,proto3,oneof"`
}

type ValueMatcher_PresentMatch struct {
	PresentMatch bool `protobuf:"varint,5,opt,name=present_match,json=presentMatch,proto3,oneof"`
}

type ValueMatcher_ListMatch struct {
	ListMatch *ListMatcher `protobuf:"bytes,6,opt,name=list_match,json=listMatch,proto3,oneof"`
}

func (*ValueMatcher_NullMatch_) isValueMatcher_MatchPattern() {}

func (*ValueMatcher_DoubleMatch) isValueMatcher_MatchPattern() {}

func (*ValueMatcher_StringMatch) isValueMatcher_MatchPattern() {}

func (*ValueMatcher_BoolMatch) isValueMatcher_MatchPattern() {}

func (*ValueMatcher_PresentMatch) isValueMatcher_MatchPattern() {}

func (*ValueMatcher_ListMatch) isValueMatcher_MatchPattern() {}

func (m *ValueMatcher) GetMatchPattern() isValueMatcher_MatchPattern {
	if m != nil {
		return m.MatchPattern
	}
	return nil
}

func (m *ValueMatcher) GetNullMatch() *ValueMatcher_NullMatch {
	if x, ok := m.GetMatchPattern().(*ValueMatcher_NullMatch_); ok {
		return x.NullMatch
	}
	return nil
}

func (m *ValueMatcher) GetDoubleMatch() *DoubleMatcher {
	if x, ok := m.GetMatchPattern().(*ValueMatcher_DoubleMatch); ok {
		return x.DoubleMatch
	}
	return nil
}

func (m *ValueMatcher) GetStringMatch() *StringMatcher {
	if x, ok := m.GetMatchPattern().(*ValueMatcher_StringMatch); ok {
		return x.StringMatch
	}
	return nil
}

func (m *ValueMatcher) GetBoolMatch() bool {
	if x, ok := m.GetMatchPattern().(*ValueMatcher_BoolMatch); ok {
		return x.BoolMatch
	}
	return false
}

func (m *ValueMatcher) GetPresentMatch() bool {
	if x, ok := m.GetMatchPattern().(*ValueMatcher_PresentMatch); ok {
		return x.PresentMatch
	}
	return false
}

func (m *ValueMatcher) GetListMatch() *ListMatcher {
	if x, ok := m.GetMatchPattern().(*ValueMatcher_ListMatch); ok {
		return x.ListMatch
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ValueMatcher) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ValueMatcher_NullMatch_)(nil),
		(*ValueMatcher_DoubleMatch)(nil),
		(*ValueMatcher_StringMatch)(nil),
		(*ValueMatcher_BoolMatch)(nil),
		(*ValueMatcher_PresentMatch)(nil),
		(*ValueMatcher_ListMatch)(nil),
	}
}

// NullMatch is an empty message to specify a null value.
type ValueMatcher_NullMatch struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValueMatcher_NullMatch) Reset()         { *m = ValueMatcher_NullMatch{} }
func (m *ValueMatcher_NullMatch) String() string { return proto.CompactTextString(m) }
func (*ValueMatcher_NullMatch) ProtoMessage()    {}
func (*ValueMatcher_NullMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_145b36501d266253, []int{0, 0}
}

func (m *ValueMatcher_NullMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValueMatcher_NullMatch.Unmarshal(m, b)
}
func (m *ValueMatcher_NullMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValueMatcher_NullMatch.Marshal(b, m, deterministic)
}
func (m *ValueMatcher_NullMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValueMatcher_NullMatch.Merge(m, src)
}
func (m *ValueMatcher_NullMatch) XXX_Size() int {
	return xxx_messageInfo_ValueMatcher_NullMatch.Size(m)
}
func (m *ValueMatcher_NullMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_ValueMatcher_NullMatch.DiscardUnknown(m)
}

var xxx_messageInfo_ValueMatcher_NullMatch proto.InternalMessageInfo

// Specifies the way to match a list value.
type ListMatcher struct {
	// Types that are valid to be assigned to MatchPattern:
	//	*ListMatcher_OneOf
	MatchPattern         isListMatcher_MatchPattern `protobuf_oneof:"match_pattern"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ListMatcher) Reset()         { *m = ListMatcher{} }
func (m *ListMatcher) String() string { return proto.CompactTextString(m) }
func (*ListMatcher) ProtoMessage()    {}
func (*ListMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_145b36501d266253, []int{1}
}

func (m *ListMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMatcher.Unmarshal(m, b)
}
func (m *ListMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMatcher.Marshal(b, m, deterministic)
}
func (m *ListMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMatcher.Merge(m, src)
}
func (m *ListMatcher) XXX_Size() int {
	return xxx_messageInfo_ListMatcher.Size(m)
}
func (m *ListMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_ListMatcher proto.InternalMessageInfo

type isListMatcher_MatchPattern interface {
	isListMatcher_MatchPattern()
}

type ListMatcher_OneOf struct {
	OneOf *ValueMatcher `protobuf:"bytes,1,opt,name=one_of,json=oneOf,proto3,oneof"`
}

func (*ListMatcher_OneOf) isListMatcher_MatchPattern() {}

func (m *ListMatcher) GetMatchPattern() isListMatcher_MatchPattern {
	if m != nil {
		return m.MatchPattern
	}
	return nil
}

func (m *ListMatcher) GetOneOf() *ValueMatcher {
	if x, ok := m.GetMatchPattern().(*ListMatcher_OneOf); ok {
		return x.OneOf
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ListMatcher) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ListMatcher_OneOf)(nil),
	}
}

func init() {
	proto.RegisterType((*ValueMatcher)(nil), "envoy.type.matcher.ValueMatcher")
	proto.RegisterType((*ValueMatcher_NullMatch)(nil), "envoy.type.matcher.ValueMatcher.NullMatch")
	proto.RegisterType((*ListMatcher)(nil), "envoy.type.matcher.ListMatcher")
}

func init() { proto.RegisterFile("envoy/type/matcher/value.proto", fileDescriptor_145b36501d266253) }

var fileDescriptor_145b36501d266253 = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcd, 0x4a, 0xeb, 0x40,
	0x1c, 0xc5, 0x9b, 0xf6, 0x36, 0xdc, 0xfe, 0xd3, 0xc2, 0x65, 0xb8, 0x70, 0x2f, 0x05, 0x6d, 0x2c,
	0x08, 0xc5, 0x45, 0x02, 0xba, 0xd1, 0x9d, 0x04, 0x91, 0x82, 0x5f, 0xa5, 0x82, 0x4b, 0xcb, 0xc4,
	0x4c, 0x35, 0x30, 0x9d, 0x09, 0xf3, 0x51, 0xec, 0xce, 0xc7, 0xf0, 0x59, 0x7c, 0x02, 0xb7, 0xbe,
	0x8d, 0x4b, 0xc9, 0xcc, 0xc4, 0x16, 0x8c, 0xb8, 0x9b, 0xf9, 0x9f, 0x73, 0x7e, 0x39, 0x93, 0x19,
	0xd8, 0x26, 0x6c, 0xc9, 0x57, 0xb1, 0x5a, 0x15, 0x24, 0x5e, 0x60, 0x75, 0xf7, 0x40, 0x44, 0xbc,
	0xc4, 0x54, 0x93, 0xa8, 0x10, 0x5c, 0x71, 0x84, 0x8c, 0x1e, 0x95, 0x7a, 0xe4, 0xf4, 0xfe, 0xa0,
	0x26, 0xc3, 0xf4, 0x22, 0x25, 0xc2, 0x86, 0x6a, 0x0d, 0x52, 0x89, 0x9c, 0xdd, 0x3b, 0xc3, 0x96,
	0xce, 0x0a, 0x1c, 0x63, 0xc6, 0xb8, 0xc2, 0x2a, 0xe7, 0x4c, 0xc6, 0x52, 0x61, 0xa5, 0xa5, 0x93,
	0xff, 0x2d, 0x31, 0xcd, 0x33, 0xac, 0x48, 0x5c, 0x2d, 0xac, 0x30, 0x7c, 0x6e, 0x41, 0xf7, 0xa6,
	0x6c, 0x77, 0x61, 0xa9, 0xe8, 0x0c, 0x80, 0x69, 0x4a, 0x67, 0xe6, 0x2b, 0xff, 0xbd, 0xd0, 0x1b,
	0x05, 0xfb, 0x7b, 0xd1, 0xd7, 0xce, 0xd1, 0x66, 0x2a, 0xba, 0xd4, 0x94, 0x9a, 0xf5, 0xb8, 0x31,
	0xed, 0xb0, 0x6a, 0x83, 0x4e, 0xa1, 0x9b, 0x71, 0x9d, 0x52, 0xe2, 0x70, 0x4d, 0x83, 0xdb, 0xa9,
	0xc3, 0x9d, 0x18, 0x9f, 0xe3, 0x8d, 0x1b, 0xd3, 0x20, 0x5b, 0x0f, 0x4a, 0x8e, 0x3d, 0xad, 0xe3,
	0xb4, 0xbe, 0xe7, 0x5c, 0x1b, 0xdf, 0x06, 0x47, 0xae, 0x07, 0x68, 0x00, 0x90, 0x72, 0x5e, 0x1d,
	0xee, 0x57, 0xe8, 0x8d, 0x7e, 0x97, 0x85, 0xcb, 0x99, 0x35, 0xec, 0x42, 0xaf, 0x10, 0x44, 0x12,
	0xa6, 0x9c, 0xa7, 0xed, 0x3c, 0x5d, 0x37, 0xb6, 0xb6, 0x63, 0x00, 0x9a, 0xcb, 0xca, 0xe3, 0x9b,
	0x36, 0x83, 0xba, 0x36, 0xe7, 0xb9, 0x54, 0xeb, 0x2e, 0x1d, 0x5a, 0x6d, 0xfb, 0x01, 0x74, 0x3e,
	0xff, 0x59, 0xf2, 0x17, 0x7a, 0x26, 0x30, 0x2b, 0xb0, 0x52, 0x44, 0x30, 0xd4, 0x7a, 0x4f, 0xbc,
	0xe1, 0x2d, 0x04, 0x1b, 0x71, 0x74, 0x04, 0x3e, 0x67, 0x64, 0xc6, 0xe7, 0xee, 0x52, 0xc2, 0x9f,
	0x2e, 0x65, 0xdc, 0x98, 0xb6, 0x39, 0x23, 0x57, 0xf3, 0x7a, 0x7e, 0x72, 0xf8, 0xf2, 0xf4, 0xfa,
	0xe6, 0x37, 0xff, 0x78, 0x10, 0xe6, 0xdc, 0xc2, 0x0a, 0xc1, 0x1f, 0x57, 0x35, 0xdc, 0x04, 0x0c,
	0x78, 0x52, 0x3e, 0x99, 0x89, 0x97, 0xfa, 0xe6, 0xed, 0x1c, 0x7c, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x57, 0x34, 0x64, 0x98, 0xeb, 0x02, 0x00, 0x00,
}
