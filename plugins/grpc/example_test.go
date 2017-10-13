package grpc_test

import (
	"log"

	ocstats "github.com/census-instrumentation/opencensus-go/stats"
	ocgrpc "github.com/rakyll/census/plugins/grpc"
	"google.golang.org/grpc"
)

func ExampleClient() {
	ch := make(chan *ocstats.ViewData, 256)
	server, err := grpc.Dial("server:7656",
		grpc.WithStatsHandler(ocgrpc.NewClientStatsHandler(ch)),
		grpc.WithUnaryInterceptor(ocgrpc.NewTracingUnaryClientInterceptor()),
		grpc.WithStreamInterceptor(ocgrpc.NewTracingStreamClientInterceptor()),
	)
	if err != nil {
		log.Fatal(err)
	}
	_ = server
}
