package main

import (
	"github.com/learnCase2/grpc/proto"
	"github.com/spf13/cast"
	"google.golang.org/grpc"
	"log"
	"net"
)

const PORT = 9002

func main() {
	server := grpc.NewServer()
	proto.RegisterStreamServiceServer(server, &StreamService{})

	listen, err := net.Listen("tcp", ":"+cast.ToString(PORT))
	if err != nil {
		log.Fatalf("net.Listen err:%v", err)
	}
	server.Serve(listen)
}

type StreamService struct {
	proto.UnimplementedStreamServiceServer
}

func (s *StreamService) List(
	req *proto.StreamReq,
	stream proto.StreamService_ListServer,
) error {
	for n := 0; n <= 6; n++ {
		err := stream.Send(&proto.StreamResp{
			Pt: &proto.StreamPoint{
				Name:  req.GetPt().GetName(),
				Value: req.GetPt().GetValue() + int32(n),
			},
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *StreamService) Record(
	stream proto.StreamService_RecordServer,
) error {
	return nil
}

func (s *StreamService) Route(
	stream proto.StreamService_RouteServer,
) error {
	return nil
}
