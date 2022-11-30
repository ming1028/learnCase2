package main

import (
	"github.com/learnCase2/grpc/proto"
	"github.com/spf13/cast"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const PORT = 9002

func main() {
	conn, err := grpc.Dial(":"+cast.ToString(PORT), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.Dial err:%v", err)
	}
	defer conn.Close()

	client := proto.NewStreamServiceClient(conn)

	err = printLists(client, &proto.StreamReq{
		Pt: &proto.StreamPoint{
			Name:  "gRPC Stream Client: List",
			Value: 2022,
		},
	})
	if err != nil {
		log.Fatalf("printList err:%v", err)
	}

	err = printRecord(client, &proto.StreamReq{
		Pt: &proto.StreamPoint{
			Name:  "gRPC Stream Client: Record",
			Value: 2022,
		},
	})
	if err != nil {
		log.Fatalf("printRecord err:%v", err)
	}

	err = printRoute(client, &proto.StreamReq{
		Pt: &proto.StreamPoint{
			Name:  "gRPC Stream Client: Route",
			Value: 2022,
		},
	})
	if err != nil {
		log.Fatalf("printRoute err:%v", err)
	}

}

func printLists(
	client proto.StreamServiceClient,
	req *proto.StreamReq,
) error {
	return nil
}

func printRecord(
	client proto.StreamServiceClient,
	req *proto.StreamReq,
) error {
	return nil
}

func printRoute(
	client proto.StreamServiceClient,
	req *proto.StreamReq,
) error {
	return nil
}
