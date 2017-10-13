// Package grpc contains OpenCensus gRPC integrations.
package grpc

import (
	ocstats "github.com/census-instrumentation/opencensus-go/stats"
	"google.golang.org/grpc"
	"google.golang.org/grpc/stats"
)

// ClientStatsHandler returns a gRPC stats handlers that
// automatically registers the given views. If no views are
// specificed, it registers all of the default client views.
// Collected data will be sent to the view data channel, ch.
func ClientStatsHandler(ch chan *ocstats.ViewData, v ...ocstats.View) stats.Handler {
	panic("not implemented")
}

// ServerStatsHandler returns a gRPC stats handlers that
// automatically registers the given views. If no views are
// specificed, it registers all of the default server views.
// Collected data will be sent to the view data channel, ch.
func ServerStatsHandler(ch chan *ocstats.ViewData, v ...ocstats.View) stats.Handler {
	panic("not implemented")
}

// TracingUnaryServerInterceptor returns a unary interceptor to be installed
// when initializing a gRPC server. The interceptor automatically
// extracts traces from the incoming requests and propagates them
// via the call's context.
func TracingUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	panic("not implemented")
}

// TracingStreamServerInterceptor returns a steam interceptor to be installed
// when initializing a gRPC server. It works similar to the
// TraceUnaryServerInterceptor, but with streaming incoming calls.
func TracingStreamServerInterceptor() grpc.StreamServerInterceptor {
	panic("not implemented")
}

// TracingUnaryClientInterceptor returns a unary interceptor to be installed
// when initializing a gRPC client. The interceptor automatically
// insert traces to the outgoing requests.
func TracingUnaryClientInterceptor() grpc.UnaryClientInterceptor {
	panic("not implemented")
}

// TracingStreamClientInterceptor is similar to TracingUnaryClientInterceptor,
// but works with streaming outgoing calls.
func TracingStreamClientInterceptor() grpc.StreamClientInterceptor {
	panic("not implemented")
}
