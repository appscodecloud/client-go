// Code generated by protoc-gen-go.
// source: appscode/api/annotations.proto
// DO NOT EDIT!

/*
Package appscode_api is a generated protocol buffer package.

It is generated from these files:
	appscode/api/annotations.proto
	appscode/api/cors.proto

It has these top-level messages:
	CorsRule
*/
package appscode_api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

var E_Cors = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MethodOptions)(nil),
	ExtensionType: (*CorsRule)(nil),
	Field:         50000,
	Name:          "appscode.api.cors",
	Tag:           "bytes,50000,opt,name=cors",
}

func init() {
	proto.RegisterExtension(E_Cors)
}

func init() { proto.RegisterFile("appscode/api/annotations.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0x2c, 0x28, 0x28,
	0x4e, 0xce, 0x4f, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c,
	0xc9, 0xcc, 0xcf, 0x2b, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x81, 0xc9, 0xeb, 0x25,
	0x16, 0x64, 0x4a, 0x89, 0xa3, 0xa8, 0x4e, 0xce, 0x2f, 0x82, 0x2a, 0x93, 0x52, 0x48, 0xcf, 0xcf,
	0x4f, 0xcf, 0x49, 0xd5, 0x07, 0xf3, 0x92, 0x4a, 0xd3, 0xf4, 0x53, 0x52, 0x8b, 0x93, 0x8b, 0x32,
	0x0b, 0x4a, 0xf2, 0x8b, 0x20, 0x2a, 0xac, 0x7c, 0xb8, 0x58, 0x40, 0xea, 0x85, 0xe4, 0xf4, 0x20,
	0x4a, 0xf5, 0x60, 0x4a, 0xf5, 0x7c, 0x53, 0x4b, 0x32, 0xf2, 0x53, 0xfc, 0x0b, 0xc0, 0xd6, 0x4a,
	0x5c, 0x68, 0x63, 0x56, 0x60, 0xd4, 0xe0, 0x36, 0x12, 0xd3, 0x43, 0xb6, 0x59, 0xcf, 0x39, 0xbf,
	0xa8, 0x38, 0xa8, 0x34, 0x27, 0x35, 0x08, 0x6c, 0x8a, 0x93, 0x1a, 0x97, 0x40, 0x72, 0x7e, 0x2e,
	0x8a, 0x12, 0x27, 0x01, 0x47, 0x84, 0xeb, 0x03, 0x40, 0x86, 0x07, 0x30, 0x26, 0xb1, 0x81, 0x6d,
	0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xef, 0x78, 0xac, 0xdf, 0xe7, 0x00, 0x00, 0x00,
}
