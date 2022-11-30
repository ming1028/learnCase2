package main

import "github.com/learnCase2/grpc/proto"

const PORT = 9002

func main() {

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
