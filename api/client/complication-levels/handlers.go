package complication_levels

import (
	"context"
	"github.com/jannyjacky1/barmen/api/client/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type ComplicationLevelsServer struct {
}

func (s *ComplicationLevelsServer) Get(ctx context.Context, request *pb.Request) (*pb.ComplicationLevelsList, error) {
	var response pb.ComplicationLevelsList
	return &response, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterComplicationLevelsServer(grpcServer, &ComplicationLevelsServer{})
	grpcServer.Serve(lis)
}
