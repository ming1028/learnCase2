package main

import (
	"context"
	"github.com/learnCase2/grpc/proto"
	"github.com/spf13/cast"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const PORT = 9001

func main() {
	conn, err := grpc.Dial(":"+cast.ToString(PORT), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.Dial err :%v", err)
	}
	defer conn.Close()

	client := proto.NewSearchServiceClient(conn)
	resp, err := client.Search(context.Background(), &proto.SearchReq{
		Request: "gRPC",
	})
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}
	log.Printf("resp:%s", resp.GetResponse())
}
