// Code generated by protoc-gen-grpc-gateway-cors
// source: authentication.proto
// DO NOT EDIT!

/*
Package v1beta1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package v1beta1

import "github.com/grpc-ecosystem/grpc-gateway/runtime"

// ExportAuthenticationCorsPatterns returns an array of grpc gatway mux patterns for
// Authentication service to enable CORS.
func ExportAuthenticationCorsPatterns() []runtime.Pattern {
	return []runtime.Pattern{
		pattern_Authentication_Login_0,
		pattern_Authentication_Logout_0,
		pattern_Authentication_Token_0,
	}
}
