package main

import (
	"context"
	"github.com/learnCase2/grpc/proto"
	"github.com/spf13/cast"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
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
	stream, err := client.List(context.Background(), req)
	if err != nil {
		return err
	}

	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}
		if err == io.EOF {
			break
		}
		log.Printf("resp: pj.name:%s, pt.Value: %d\n", resp.GetPt().GetName(), resp.GetPt().GetValue())
	}
	return nil
}

func printRecord(
	client proto.StreamServiceClient,
	req *proto.StreamReq,
) error {
	stream, err := client.Record(context.Background())
	if err != nil {
		return err
	}
	for n := 0; n < 6; n++ {
		err := stream.Send(req)
		if err != nil {
			return err
		}
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}
	log.Printf("stream Recv pt.name: %s, pt.Value: %d\n", resp.GetPt().GetName(), resp.GetPt().GetValue())
	return nil
}

func printRoute(
	client proto.StreamServiceClient,
	req *proto.StreamReq,
) error {
	stream, err := client.Route(context.Background())
	if err != nil {
		return err
	}
	for n := 0; n <= 6; n++ {
		err := stream.Send(req)
		if err != nil {
			return err
		}
		resp, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("resp Recv pt.Name:%s, pt.Value:%d",
			resp.GetPt().GetName(),
			resp.GetPt().GetValue(),
		)
	}
	stream.CloseSend()
	return nil
}
