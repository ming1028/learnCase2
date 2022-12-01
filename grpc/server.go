package main

import (
	"context"
	"github.com/learnCase2/grpc/proto"
	"github.com/spf13/cast"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

type SearchService struct {
	proto.UnimplementedSearchServiceServer
}

func (s *SearchService) Search(
	ctx context.Context,
	r *proto.SearchReq,
) (
	*proto.SearchResp,
	error,
) {
	return &proto.SearchResp{
		Response: r.GetRequest() + " Server",
	}, nil
}

const PORT = 9001

func main() {
	cred, err := credentials.NewServerTLSFromFile(
		"./grpc/conf/server.pem",
		"./grpc/conf/server.key",
	)
	if err != nil {
		log.Fatalf("credentials.NewServerTLSFromFile err:%v", err)
	}

	server := grpc.NewServer(grpc.Creds(cred))
	proto.RegisterSearchServiceServer(server, &SearchService{})

	lis, err := net.Listen("tcp", ":"+cast.ToString(PORT))
	if err != nil {
		log.Fatalf("net listen err: %v", err)
	}
	server.Serve(lis)
}
