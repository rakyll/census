// Package grpc contains OpenCensus gRPC integrations.
package grpc

import (
	ocstats "github.com/census-instrumentation/opencensus-go/stats"
	"google.golang.org/grpc"
	"google.golang.org/grpc/stats"
)

// NewClientStatsHandler returns a gRPC stats handlers that
// automatically registers the given views. If no views are
// specificed, it registers all of the default client views.
// Collected data will be sent to the view data channel, ch.
func NewClientStatsHandler(ch chan *ocstats.ViewData, v ...ocstats.View) stats.Handler {
	panic("not implemented")
}

// NewServerStatsHandler returns a gRPC stats handlers that
// automatically registers the given views. If no views are
// specificed, it registers all of the default server views.
// Collected data will be sent to the view data channel, ch.
func NewServerStatsHandler(ch chan *ocstats.ViewData, v ...ocstats.View) stats.Handler {
	panic("not implemented")
}

// NewTracingUnaryServerInterceptor returns a unary interceptor to be installed
// when initializing a gRPC server. The interceptor automatically
// extracts traces from the incoming requests and propagates them
// via the call's context.
func NewTracingUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	panic("not implemented")
}

// NewTracingStreamServerInterceptor returns a steam interceptor to be installed
// when initializing a gRPC server. It works similar to the
// TraceUnaryServerInterceptor, but with streaming incoming calls.
func NewTracingStreamServerInterceptor() grpc.StreamServerInterceptor {
	panic("not implemented")
}

// NewTracingUnaryClientInterceptor returns a unary interceptor to be installed
// when initializing a gRPC client. The interceptor automatically
// insert traces to the outgoing requests.
func NewTracingUnaryClientInterceptor() grpc.UnaryClientInterceptor {
	panic("not implemented")
}

// NewTracingStreamClientInterceptor is similar to NewTracingUnaryClientInterceptor,
// but works with streaming outgoing calls.
func NewTracingStreamClientInterceptor() grpc.StreamClientInterceptor {
	panic("not implemented")
}
