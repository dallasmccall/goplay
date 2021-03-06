// Code generated by protoc-gen-go.
// source: dataset/dataset.proto
// DO NOT EDIT!

/*
Package dataset is a generated protocol buffer package.

It is generated from these files:
	dataset/dataset.proto

It has these top-level messages:
	DataSet
*/
package dataset

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DataSet struct {
	SchemaId  uint32         `protobuf:"varint,1,opt,name=schema_id,json=schemaId" json:"schema_id,omitempty"`
	Timestamp int64          `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Row       []*DataSet_Row `protobuf:"bytes,3,rep,name=row" json:"row,omitempty"`
}

func (m *DataSet) Reset()                    { *m = DataSet{} }
func (m *DataSet) String() string            { return proto.CompactTextString(m) }
func (*DataSet) ProtoMessage()               {}
func (*DataSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DataSet) GetRow() []*DataSet_Row {
	if m != nil {
		return m.Row
	}
	return nil
}

type DataSet_Column struct {
	BoolValue   bool    `protobuf:"varint,1,opt,name=bool_value,json=boolValue" json:"bool_value,omitempty"`
	IntValue    int32   `protobuf:"varint,2,opt,name=int_value,json=intValue" json:"int_value,omitempty"`
	FloatValue  float32 `protobuf:"fixed32,3,opt,name=float_value,json=floatValue" json:"float_value,omitempty"`
	LongValue   int64   `protobuf:"varint,4,opt,name=long_value,json=longValue" json:"long_value,omitempty"`
	DoubleValue float64 `protobuf:"fixed64,5,opt,name=double_value,json=doubleValue" json:"double_value,omitempty"`
	StringValue string  `protobuf:"bytes,6,opt,name=string_value,json=stringValue" json:"string_value,omitempty"`
}

func (m *DataSet_Column) Reset()                    { *m = DataSet_Column{} }
func (m *DataSet_Column) String() string            { return proto.CompactTextString(m) }
func (*DataSet_Column) ProtoMessage()               {}
func (*DataSet_Column) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type DataSet_Row struct {
	Column []*DataSet_Column `protobuf:"bytes,1,rep,name=column" json:"column,omitempty"`
}

func (m *DataSet_Row) Reset()                    { *m = DataSet_Row{} }
func (m *DataSet_Row) String() string            { return proto.CompactTextString(m) }
func (*DataSet_Row) ProtoMessage()               {}
func (*DataSet_Row) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

func (m *DataSet_Row) GetColumn() []*DataSet_Column {
	if m != nil {
		return m.Column
	}
	return nil
}

func init() {
	proto.RegisterType((*DataSet)(nil), "DataSet")
	proto.RegisterType((*DataSet_Column)(nil), "DataSet.Column")
	proto.RegisterType((*DataSet_Row)(nil), "DataSet.Row")
}

func init() { proto.RegisterFile("dataset/dataset.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x44, 0x90, 0xcd, 0x4a, 0xf4, 0x30,
	0x14, 0x86, 0xc9, 0xe4, 0x9b, 0x7e, 0xed, 0xe9, 0x88, 0x50, 0x10, 0xca, 0xe8, 0x68, 0x74, 0x63,
	0x56, 0x15, 0xf4, 0x12, 0x74, 0xe3, 0xf6, 0x08, 0x6e, 0x87, 0x74, 0x1a, 0xb5, 0x90, 0xf6, 0x0c,
	0x6d, 0x6a, 0xef, 0xd1, 0xbb, 0xf0, 0x4e, 0x24, 0x3f, 0xd6, 0x55, 0xe0, 0x79, 0x5e, 0xce, 0x39,
	0x6f, 0xe0, 0xac, 0x51, 0x56, 0x8d, 0xda, 0xde, 0xc5, 0xb7, 0x3a, 0x0e, 0x64, 0xe9, 0xe6, 0x7b,
	0x05, 0xff, 0x9f, 0x94, 0x55, 0x2f, 0xda, 0x16, 0xe7, 0x90, 0x8d, 0x87, 0x0f, 0xdd, 0xa9, 0x7d,
	0xdb, 0x94, 0x4c, 0x30, 0x79, 0x82, 0x69, 0x00, 0xcf, 0x4d, 0x71, 0x01, 0x99, 0x6d, 0x3b, 0x3d,
	0x5a, 0xd5, 0x1d, 0xcb, 0x95, 0x60, 0x92, 0xe3, 0x1f, 0x28, 0x2e, 0x81, 0x0f, 0x34, 0x97, 0x5c,
	0x70, 0x99, 0xdf, 0x6f, 0xaa, 0x38, 0xb1, 0x42, 0x9a, 0xd1, 0x89, 0xed, 0x17, 0x83, 0xe4, 0x91,
	0xcc, 0xd4, 0xf5, 0xc5, 0x0e, 0xa0, 0x26, 0x32, 0xfb, 0x4f, 0x65, 0x26, 0xed, 0xd7, 0xa4, 0x98,
	0x39, 0xf2, 0xea, 0x80, 0x3b, 0xa2, 0xed, 0x6d, 0xb4, 0x6e, 0xcf, 0x1a, 0xd3, 0xb6, 0xb7, 0x41,
	0x5e, 0x41, 0xfe, 0x66, 0x48, 0xfd, 0x6a, 0x2e, 0x98, 0x5c, 0x21, 0x78, 0x14, 0x02, 0x3b, 0x00,
	0x43, 0xfd, 0x7b, 0xf4, 0xff, 0xc2, 0x99, 0x8e, 0x04, 0x7d, 0x0d, 0x9b, 0x86, 0xa6, 0xda, 0xe8,
	0x18, 0x58, 0x0b, 0x26, 0x19, 0xe6, 0x81, 0x2d, 0x91, 0xd1, 0x0e, 0xed, 0x32, 0x23, 0x11, 0x4c,
	0x66, 0x98, 0x07, 0xe6, 0x23, 0xdb, 0x0a, 0x38, 0xd2, 0x5c, 0xdc, 0x42, 0x72, 0xf0, 0x95, 0x4a,
	0xe6, 0x6b, 0x9f, 0x2e, 0xb5, 0x43, 0x53, 0x8c, 0xba, 0x4e, 0xfc, 0x57, 0x3f, 0xfc, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xed, 0xa0, 0xa4, 0x24, 0x83, 0x01, 0x00, 0x00,
}
