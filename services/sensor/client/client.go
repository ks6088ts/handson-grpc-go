package client

import (
	"context"
	"io"

	"google.golang.org/grpc"

	pb "github.com/ks6088ts/handson-grpc-go/services/sensor/sensor"
	"google.golang.org/grpc/credentials/insecure"
)

// SensorClient represents a client for sensor service
type SensorClient struct {
	client pb.SensorClient
}

// SensorState denotes a state for sensor service
type SensorState struct {
	X float64
	Y float64
	Z float64
}

// NewSensorClient returns a newly created SensorClient
func NewSensorClient(addr string) (*SensorClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewSensorClient(conn)
	return &SensorClient{
		client: client,
	}, nil
}

// GetSensorState returns a sensor state
func (c *SensorClient) GetSensorState() (*SensorState, error) {
	state, err := c.client.GetSensorState(context.Background(), &pb.DummyRequest{})
	if err != nil {
		return nil, err
	}
	return &SensorState{
		X: state.X,
		Y: state.Y,
		Z: state.Z,
	}, nil
}

// GetSensorStates returns multiple sensor states
func (c *SensorClient) GetSensorStates() ([]SensorState, error) {
	stream, err := c.client.GetSensorStates(context.Background(), &pb.DummyRequest{})
	if err != nil {
		return nil, err
	}

	var states []SensorState
	for {
		state, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		states = append(states, SensorState{
			X: state.X,
			Y: state.Y,
			Z: state.Z,
		})
	}
	return states, nil
}
