package grpcprotobuf

import (
	"log"
	"net"

	"github.com/dinel13/thesis-ac/protobuf-vs-json/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Start entrypoint
func Start() {
	lis, _ := net.Listen("tcp", ":60000")

	srv := grpc.NewServer()
	proto.RegisterAPIServer(srv, &Server{})
	log.Println(srv.Serve(lis))
}

// Server type
type Server struct{}

// CreateUser handler
func (s *Server) CreateKrs(ctx context.Context, in *proto.Krs) (*proto.Response, error) {
	return &proto.Response{
		Code:        200,
		Message:     "OK",
		MataKuliahs: in.MataKuliahs,
	}, nil
}
