package main

import (
	"bytes"
	context "context"
	"net"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logurs "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	lis := must(net.Listen("tcp", ":12345"))
	log := logrus.New()
	srv := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_logurs.StreamServerInterceptor(logrus.NewEntry(log)),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_logurs.UnaryServerInterceptor(logrus.NewEntry(log)),
		)),
	)
	svc := new(helloService)
	RegisterHelloServiceServer(srv, svc)
	println("serving at :12345")
	srv.Serve(lis)
}

type helloService struct{}

// SayHello implements HelloServiceServer
func (*helloService) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	var b bytes.Buffer
	b.WriteString(req.Greet)
	b.WriteString("\n\n")
	spew.Fdump(&b, os.Environ())
	md, _ := metadata.FromIncomingContext(ctx)
	spew.Fdump(&b, md)
	return &HelloResponse{
		GreetBack: b.String(),
	}, nil
}

// mustEmbedUnimplementedHelloServiceServer implements HelloServiceServer
func (*helloService) mustEmbedUnimplementedHelloServiceServer() {
	panic("unimplemented")
}

var _ HelloServiceServer = (*helloService)(nil)

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
