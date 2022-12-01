package main

import (
	"context"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/learnCase2/grpc/proto"
	"github.com/spf13/cast"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"net"
	"runtime/debug"
)

const PORT = 9002

func main() {
	opts := []grpc.ServerOption{
		grpc_middleware.WithStreamServerChain(
			LoggingServerInterceptor(),
		),
	}
	server := grpc.NewServer(opts...)
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
	for {
		recv, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(
				&proto.StreamResp{
					Pt: &proto.StreamPoint{
						Name:  "gRPC Stream Server:Record",
						Value: 1,
					},
				},
			)
		}
		if err != nil {
			return err
		}
		log.Printf("stream Recv pt.name: %s, pt.Value: %d\n", recv.GetPt().GetName(), recv.GetPt().GetValue())
	}
	return nil
}

func (s *StreamService) Route(
	stream proto.StreamService_RouteServer,
) error {
	n := 0
	for {
		err := stream.Send(&proto.StreamResp{
			Pt: &proto.StreamPoint{
				Name:  "gRPC stream Client: Route",
				Value: int32(n),
			},
		})
		if err != nil {
			return err
		}

		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}
		n++

		log.Printf("steam Recn pt.Name:%s, pt.Value:%d",
			req.GetPt().GetName(),
			req.GetPt().GetValue(),
		)
	}
	return nil
}

func LoggingInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (
	interface{},
	error,
) {
	log.Printf("gRPc method: %s, %v", info.FullMethod, req)
	resp, err := handler(ctx, req)
	log.Printf("gRPc method: %s, %v", info.FullMethod, req)
	return resp, err
}

func RecoveryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (
	resp interface{},
	err error,
) {
	defer func() {
		if e := recover(); e != nil {
			debug.PrintStack()
			err = status.Errorf(codes.Internal, "Panic err: %v")
		}
	}()
	return handler(ctx, req)
}

func LoggingServerInterceptor() grpc.StreamServerInterceptor {
	return func(
		srv interface{}, ss grpc.ServerStream,
		info *grpc.StreamServerInfo, handler grpc.StreamHandler,
	) error {
		log.Printf("gRPc method: %s, %v", info.FullMethod, ss)
		return handler(srv, ss)
	}
}
