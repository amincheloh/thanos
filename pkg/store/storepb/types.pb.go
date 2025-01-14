// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: store/storepb/types.proto

package storepb

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/thanos-io/thanos/pkg/store/labelpb"
	github_com_thanos_io_thanos_pkg_store_labelpb "github.com/thanos-io/thanos/pkg/store/labelpb"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// / PartialResponseStrategy controls partial response handling.
type PartialResponseStrategy int32

const (
	/// WARN strategy tells server to treat any error that will related to single StoreAPI (e.g missing chunk series because of underlying
	/// storeAPI is temporarily not available) as warning which will not fail the whole query (still OK response).
	/// Server should produce those as a warnings field in response.
	PartialResponseStrategy_WARN PartialResponseStrategy = 0
	/// ABORT strategy tells server to treat any error that will related to single StoreAPI (e.g missing chunk series because of underlying
	/// storeAPI is temporarily not available) as the gRPC error that aborts the query.
	///
	/// This is especially useful for any rule/alert evaluations on top of StoreAPI which usually does not tolerate partial
	/// errors.
	PartialResponseStrategy_ABORT PartialResponseStrategy = 1
)

var PartialResponseStrategy_name = map[int32]string{
	0: "WARN",
	1: "ABORT",
}

var PartialResponseStrategy_value = map[string]int32{
	"WARN":  0,
	"ABORT": 1,
}

func (x PartialResponseStrategy) String() string {
	return proto.EnumName(PartialResponseStrategy_name, int32(x))
}

func (PartialResponseStrategy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_121fba57de02d8e0, []int{0}
}

type Chunk_Encoding int32

const (
	Chunk_XOR Chunk_Encoding = 0
)

var Chunk_Encoding_name = map[int32]string{
	0: "XOR",
}

var Chunk_Encoding_value = map[string]int32{
	"XOR": 0,
}

func (x Chunk_Encoding) String() string {
	return proto.EnumName(Chunk_Encoding_name, int32(x))
}

func (Chunk_Encoding) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_121fba57de02d8e0, []int{0, 0}
}

type LabelMatcher_Type int32

const (
	LabelMatcher_EQ  LabelMatcher_Type = 0
	LabelMatcher_NEQ LabelMatcher_Type = 1
	LabelMatcher_RE  LabelMatcher_Type = 2
	LabelMatcher_NRE LabelMatcher_Type = 3
)

var LabelMatcher_Type_name = map[int32]string{
	0: "EQ",
	1: "NEQ",
	2: "RE",
	3: "NRE",
}

var LabelMatcher_Type_value = map[string]int32{
	"EQ":  0,
	"NEQ": 1,
	"RE":  2,
	"NRE": 3,
}

func (x LabelMatcher_Type) String() string {
	return proto.EnumName(LabelMatcher_Type_name, int32(x))
}

func (LabelMatcher_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_121fba57de02d8e0, []int{3, 0}
}

type Chunk struct {
	Type Chunk_Encoding `protobuf:"varint,1,opt,name=type,proto3,enum=thanos.Chunk_Encoding" json:"type,omitempty"`
	Data []byte         `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Chunk) Reset()         { *m = Chunk{} }
func (m *Chunk) String() string { return proto.CompactTextString(m) }
func (*Chunk) ProtoMessage()    {}
func (*Chunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_121fba57de02d8e0, []int{0}
}
func (m *Chunk) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Chunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Chunk.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Chunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chunk.Merge(m, src)
}
func (m *Chunk) XXX_Size() int {
	return m.Size()
}
func (m *Chunk) XXX_DiscardUnknown() {
	xxx_messageInfo_Chunk.DiscardUnknown(m)
}

var xxx_messageInfo_Chunk proto.InternalMessageInfo

type Series struct {
	Labels []github_com_thanos_io_thanos_pkg_store_labelpb.ZLabel `protobuf:"bytes,1,rep,name=labels,proto3,customtype=github.com/thanos-io/thanos/pkg/store/labelpb.ZLabel" json:"labels"`
	Chunks []AggrChunk                                            `protobuf:"bytes,2,rep,name=chunks,proto3" json:"chunks"`
}

func (m *Series) Reset()         { *m = Series{} }
func (m *Series) String() string { return proto.CompactTextString(m) }
func (*Series) ProtoMessage()    {}
func (*Series) Descriptor() ([]byte, []int) {
	return fileDescriptor_121fba57de02d8e0, []int{1}
}
func (m *Series) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Series) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Series.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Series) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Series.Merge(m, src)
}
func (m *Series) XXX_Size() int {
	return m.Size()
}
func (m *Series) XXX_DiscardUnknown() {
	xxx_messageInfo_Series.DiscardUnknown(m)
}

var xxx_messageInfo_Series proto.InternalMessageInfo

type AggrChunk struct {
	MinTime int64  `protobuf:"varint,1,opt,name=min_time,json=minTime,proto3" json:"min_time,omitempty"`
	MaxTime int64  `protobuf:"varint,2,opt,name=max_time,json=maxTime,proto3" json:"max_time,omitempty"`
	Raw     *Chunk `protobuf:"bytes,3,opt,name=raw,proto3" json:"raw,omitempty"`
	Count   *Chunk `protobuf:"bytes,4,opt,name=count,proto3" json:"count,omitempty"`
	Sum     *Chunk `protobuf:"bytes,5,opt,name=sum,proto3" json:"sum,omitempty"`
	Min     *Chunk `protobuf:"bytes,6,opt,name=min,proto3" json:"min,omitempty"`
	Max     *Chunk `protobuf:"bytes,7,opt,name=max,proto3" json:"max,omitempty"`
	Counter *Chunk `protobuf:"bytes,8,opt,name=counter,proto3" json:"counter,omitempty"`
}

func (m *AggrChunk) Reset()         { *m = AggrChunk{} }
func (m *AggrChunk) String() string { return proto.CompactTextString(m) }
func (*AggrChunk) ProtoMessage()    {}
func (*AggrChunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_121fba57de02d8e0, []int{2}
}
func (m *AggrChunk) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AggrChunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AggrChunk.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AggrChunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AggrChunk.Merge(m, src)
}
func (m *AggrChunk) XXX_Size() int {
	return m.Size()
}
func (m *AggrChunk) XXX_DiscardUnknown() {
	xxx_messageInfo_AggrChunk.DiscardUnknown(m)
}

var xxx_messageInfo_AggrChunk proto.InternalMessageInfo

// Matcher specifies a rule, which can match or set of labels or not.
type LabelMatcher struct {
	Type  LabelMatcher_Type `protobuf:"varint,1,opt,name=type,proto3,enum=thanos.LabelMatcher_Type" json:"type,omitempty"`
	Name  string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Value string            `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *LabelMatcher) Reset()         { *m = LabelMatcher{} }
func (m *LabelMatcher) String() string { return proto.CompactTextString(m) }
func (*LabelMatcher) ProtoMessage()    {}
func (*LabelMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_121fba57de02d8e0, []int{3}
}
func (m *LabelMatcher) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LabelMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LabelMatcher.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LabelMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabelMatcher.Merge(m, src)
}
func (m *LabelMatcher) XXX_Size() int {
	return m.Size()
}
func (m *LabelMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_LabelMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_LabelMatcher proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("thanos.PartialResponseStrategy", PartialResponseStrategy_name, PartialResponseStrategy_value)
	proto.RegisterEnum("thanos.Chunk_Encoding", Chunk_Encoding_name, Chunk_Encoding_value)
	proto.RegisterEnum("thanos.LabelMatcher_Type", LabelMatcher_Type_name, LabelMatcher_Type_value)
	proto.RegisterType((*Chunk)(nil), "thanos.Chunk")
	proto.RegisterType((*Series)(nil), "thanos.Series")
	proto.RegisterType((*AggrChunk)(nil), "thanos.AggrChunk")
	proto.RegisterType((*LabelMatcher)(nil), "thanos.LabelMatcher")
}

func init() { proto.RegisterFile("store/storepb/types.proto", fileDescriptor_121fba57de02d8e0) }

var fileDescriptor_121fba57de02d8e0 = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0xc1, 0x6f, 0xd3, 0x3e,
	0x14, 0xc7, 0xe3, 0x24, 0x4d, 0x5b, 0xff, 0xf6, 0x43, 0xc1, 0x4c, 0x90, 0xee, 0x90, 0x56, 0x41,
	0x88, 0x6a, 0xd2, 0x12, 0x69, 0x70, 0xe4, 0xd2, 0xa2, 0xde, 0x60, 0x63, 0x5e, 0x25, 0xd0, 0x84,
	0x84, 0xdc, 0xcc, 0x4a, 0xad, 0x35, 0x76, 0x94, 0x38, 0xd0, 0xfe, 0x17, 0x20, 0xee, 0xfc, 0x3d,
	0x3d, 0xee, 0x88, 0x38, 0x4c, 0xd0, 0xfe, 0x23, 0xc8, 0x4e, 0x0a, 0x54, 0xca, 0x25, 0x7a, 0x79,
	0xdf, 0xcf, 0x7b, 0xdf, 0xf8, 0xe5, 0x19, 0xf6, 0x0a, 0x29, 0x72, 0x1a, 0xe9, 0x67, 0x36, 0x8b,
	0xe4, 0x2a, 0xa3, 0x45, 0x98, 0xe5, 0x42, 0x0a, 0xe4, 0xc8, 0x39, 0xe1, 0xa2, 0x38, 0x3a, 0x4c,
	0x44, 0x22, 0x74, 0x2a, 0x52, 0x51, 0xa5, 0x1e, 0xd5, 0x85, 0x0b, 0x32, 0xa3, 0x8b, 0xfd, 0xc2,
	0xe0, 0x3d, 0x6c, 0xbd, 0x9c, 0x97, 0xfc, 0x06, 0x1d, 0x43, 0x5b, 0xe5, 0x3d, 0x30, 0x00, 0xc3,
	0x7b, 0xa7, 0x0f, 0xc3, 0xaa, 0x61, 0xa8, 0xc5, 0x70, 0xc2, 0x63, 0x71, 0xcd, 0x78, 0x82, 0x35,
	0x83, 0x10, 0xb4, 0xaf, 0x89, 0x24, 0x9e, 0x39, 0x00, 0xc3, 0x03, 0xac, 0xe3, 0xe0, 0x01, 0xec,
	0xec, 0x28, 0xd4, 0x86, 0xd6, 0xbb, 0x73, 0xec, 0x1a, 0xc1, 0x37, 0x00, 0x9d, 0x4b, 0x9a, 0x33,
	0x5a, 0xa0, 0x18, 0x3a, 0xda, 0xbf, 0xf0, 0xc0, 0xc0, 0x1a, 0xfe, 0x77, 0xfa, 0xff, 0xce, 0xe1,
	0x95, 0xca, 0x8e, 0x5f, 0xac, 0xef, 0xfa, 0xc6, 0x8f, 0xbb, 0xfe, 0xf3, 0x84, 0xc9, 0x79, 0x39,
	0x0b, 0x63, 0x91, 0x46, 0x15, 0x70, 0xc2, 0x44, 0x1d, 0x45, 0xd9, 0x4d, 0x12, 0xed, 0x1d, 0x25,
	0xbc, 0xd2, 0xd5, 0xb8, 0x6e, 0x8d, 0x22, 0xe8, 0xc4, 0xea, 0x83, 0x0b, 0xcf, 0xd4, 0x26, 0xf7,
	0x77, 0x26, 0xa3, 0x24, 0xc9, 0xf5, 0x51, 0xc6, 0xb6, 0x32, 0xc2, 0x35, 0x16, 0x7c, 0x35, 0x61,
	0xf7, 0x8f, 0x86, 0x7a, 0xb0, 0x93, 0x32, 0xfe, 0x41, 0xb2, 0xb4, 0x9a, 0x83, 0x85, 0xdb, 0x29,
	0xe3, 0x53, 0x96, 0x52, 0x2d, 0x91, 0x65, 0x25, 0x99, 0xb5, 0x44, 0x96, 0x5a, 0xea, 0x43, 0x2b,
	0x27, 0x9f, 0x3c, 0x6b, 0x00, 0xfe, 0x3d, 0x96, 0xee, 0x88, 0x95, 0x82, 0x1e, 0xc3, 0x56, 0x2c,
	0x4a, 0x2e, 0x3d, 0xbb, 0x09, 0xa9, 0x34, 0xd5, 0xa5, 0x28, 0x53, 0xaf, 0xd5, 0xd8, 0xa5, 0x28,
	0x53, 0x05, 0xa4, 0x8c, 0x7b, 0x4e, 0x23, 0x90, 0x32, 0xae, 0x01, 0xb2, 0xf4, 0xda, 0xcd, 0x00,
	0x59, 0xa2, 0xa7, 0xb0, 0xad, 0xbd, 0x68, 0xee, 0x75, 0x9a, 0xa0, 0x9d, 0x1a, 0x7c, 0x01, 0xf0,
	0x40, 0x0f, 0xf6, 0x35, 0x91, 0xf1, 0x9c, 0xe6, 0xe8, 0x64, 0x6f, 0x39, 0x7a, 0x7b, 0xbf, 0xae,
	0x66, 0xc2, 0xe9, 0x2a, 0xa3, 0x7f, 0xf7, 0x83, 0x93, 0x7a, 0x50, 0x5d, 0xac, 0x63, 0x74, 0x08,
	0x5b, 0x1f, 0xc9, 0xa2, 0xa4, 0x7a, 0x4e, 0x5d, 0x5c, 0xbd, 0x04, 0x43, 0x68, 0xab, 0x3a, 0xe4,
	0x40, 0x73, 0x72, 0xe1, 0x1a, 0x6a, 0x73, 0xce, 0x26, 0x17, 0x2e, 0x50, 0x09, 0x3c, 0x71, 0x4d,
	0x9d, 0xc0, 0x13, 0xd7, 0x3a, 0x0e, 0xe1, 0xa3, 0x37, 0x24, 0x97, 0x8c, 0x2c, 0x30, 0x2d, 0x32,
	0xc1, 0x0b, 0x7a, 0x29, 0x73, 0x22, 0x69, 0xb2, 0x42, 0x1d, 0x68, 0xbf, 0x1d, 0xe1, 0x33, 0xd7,
	0x40, 0x5d, 0xd8, 0x1a, 0x8d, 0xcf, 0xf1, 0xd4, 0x05, 0xe3, 0x27, 0xeb, 0x5f, 0xbe, 0xb1, 0xde,
	0xf8, 0xe0, 0x76, 0xe3, 0x83, 0x9f, 0x1b, 0x1f, 0x7c, 0xde, 0xfa, 0xc6, 0xed, 0xd6, 0x37, 0xbe,
	0x6f, 0x7d, 0xe3, 0xaa, 0x5d, 0x5f, 0xa2, 0x99, 0xa3, 0xaf, 0xc1, 0xb3, 0xdf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xd5, 0x63, 0x3a, 0x23, 0x5c, 0x03, 0x00, 0x00,
}

func (m *Chunk) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Chunk) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Chunk) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if m.Type != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Series) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Series) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Series) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Chunks) > 0 {
		for iNdEx := len(m.Chunks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Chunks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Labels) > 0 {
		for iNdEx := len(m.Labels) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.Labels[iNdEx].Size()
				i -= size
				if _, err := m.Labels[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *AggrChunk) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AggrChunk) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AggrChunk) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Counter != nil {
		{
			size, err := m.Counter.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.Max != nil {
		{
			size, err := m.Max.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.Min != nil {
		{
			size, err := m.Min.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.Sum != nil {
		{
			size, err := m.Sum.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.Count != nil {
		{
			size, err := m.Count.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Raw != nil {
		{
			size, err := m.Raw.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.MaxTime != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.MaxTime))
		i--
		dAtA[i] = 0x10
	}
	if m.MinTime != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.MinTime))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *LabelMatcher) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LabelMatcher) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LabelMatcher) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Type != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Chunk) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovTypes(uint64(m.Type))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *Series) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Labels) > 0 {
		for _, e := range m.Labels {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if len(m.Chunks) > 0 {
		for _, e := range m.Chunks {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *AggrChunk) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MinTime != 0 {
		n += 1 + sovTypes(uint64(m.MinTime))
	}
	if m.MaxTime != 0 {
		n += 1 + sovTypes(uint64(m.MaxTime))
	}
	if m.Raw != nil {
		l = m.Raw.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Count != nil {
		l = m.Count.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Sum != nil {
		l = m.Sum.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Min != nil {
		l = m.Min.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Max != nil {
		l = m.Max.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Counter != nil {
		l = m.Counter.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *LabelMatcher) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovTypes(uint64(m.Type))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Chunk) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Chunk: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Chunk: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= Chunk_Encoding(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Series) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Series: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Series: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Labels = append(m.Labels, github_com_thanos_io_thanos_pkg_store_labelpb.ZLabel{})
			if err := m.Labels[len(m.Labels)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chunks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chunks = append(m.Chunks, AggrChunk{})
			if err := m.Chunks[len(m.Chunks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AggrChunk) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AggrChunk: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AggrChunk: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinTime", wireType)
			}
			m.MinTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxTime", wireType)
			}
			m.MaxTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Raw", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Raw == nil {
				m.Raw = &Chunk{}
			}
			if err := m.Raw.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Count == nil {
				m.Count = &Chunk{}
			}
			if err := m.Count.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sum", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Sum == nil {
				m.Sum = &Chunk{}
			}
			if err := m.Sum.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Min", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Min == nil {
				m.Min = &Chunk{}
			}
			if err := m.Min.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Max", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Max == nil {
				m.Max = &Chunk{}
			}
			if err := m.Max.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Counter", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Counter == nil {
				m.Counter = &Chunk{}
			}
			if err := m.Counter.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LabelMatcher) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LabelMatcher: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LabelMatcher: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= LabelMatcher_Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
