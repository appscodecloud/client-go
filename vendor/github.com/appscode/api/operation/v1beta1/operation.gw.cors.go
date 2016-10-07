// Code generated by protoc-gen-grpc-gateway-cors
// source: operation.proto
// DO NOT EDIT!

/*
Package v1beta1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package v1beta1

import "github.com/grpc-ecosystem/grpc-gateway/runtime"

// ExportOperationsCorsPatterns returns an array of grpc gatway mux patterns for
// Operations service to enable CORS.
func ExportOperationsCorsPatterns() []runtime.Pattern {
	return []runtime.Pattern{
		pattern_Operations_Describe_0,
	}
}
