package grpc_test

import (
	"log"

	ocstats "github.com/census-instrumentation/opencensus-go/stats"
	ocgrpc "github.com/rakyll/census/plugins/grpc"
	"google.golang.org/grpc"
)

func ExampleClient() {
	ch := make(chan *ocstats.ViewData, 256)
	conn, err := grpc.Dial("server:7656",
		grpc.WithStatsHandler(ocgrpc.ClientStatsHandler(ch)),
		grpc.WithUnaryInterceptor(ocgrpc.TracingUnaryClientInterceptor()),
		grpc.WithStreamInterceptor(ocgrpc.TracingStreamClientInterceptor()),
	)
	if err != nil {
		log.Fatal(err)
	}
	_ = conn
}
