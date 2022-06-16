package server

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"

	"google.golang.org/grpc"

	pb "github.com/ks6088ts/handson-grpc-go/services/sensor/sensor"
)

type sensorServer struct {
	pb.UnimplementedSensorServer
}

// GetSensorState returns a state.
func (s *sensorServer) GetSensorState(ctx context.Context, request *pb.DummyRequest) (*pb.SensorState, error) {
	return &pb.SensorState{
		X: rand.Float64(),
		Y: rand.Float64(),
		Z: rand.Float64(),
	}, nil
}

// GetSensorStates returns multiple states.
func (s *sensorServer) GetSensorStates(request *pb.DummyRequest, stream pb.Sensor_GetSensorStatesServer) error {
	for i := 0; i < 10; i++ {
		if err := stream.Send(&pb.SensorState{
			X: rand.Float64(),
			Y: rand.Float64(),
			Z: rand.Float64(),
		}); err != nil {
			return err
		}
	}
	return nil
}

// Serve accepts incoming connections
func Serve(port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSensorServer(grpcServer, &sensorServer{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
