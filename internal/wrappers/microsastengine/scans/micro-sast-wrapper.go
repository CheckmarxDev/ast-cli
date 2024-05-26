package scans

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

type MicroSastWrapper struct {
	port int
}

var serviceConfig = `{
	"loadBalancingPolicy": "round_robin",
	"healthCheckConfig": {
		"serviceName": "MicroSastEngine"
	}
}`

func NewMicroSastWrapper(port int) *MicroSastWrapper {
	return &MicroSastWrapper{
		port: port,
	}
}

func (s *MicroSastWrapper) Scan(filePath string, dataBytes []byte) (*ScanResult, error) {
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithDefaultServiceConfig(serviceConfig),
	}
	localhostAddress := fmt.Sprintf("0.0.0.0:%d", s.port)
	conn, err := grpc.NewClient(localhostAddress, options...)

	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)

	if err != nil {
		log.Fatalf("grpc.NewClient(%q): %v", localhostAddress, err)
	}

	client := NewScanServiceClient(conn)

	request := &SingleScanRequest{
		ScanRequest: &ScanRequest{
			Id:         "idForTheScan",
			FileName:   filePath,
			SourceCode: string(dataBytes),
		},
	}

	return client.Scan(context.Background(), request)
}

func checkHealth(service string, conn *grpc.ClientConn) (*healthpb.HealthCheckResponse, error) {
	req := &healthpb.HealthCheckRequest{
		Service: service,
	}
	ctx := context.Background()
	rsp, err := healthpb.NewHealthClient(conn).Check(ctx, req)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func (s *MicroSastWrapper) CheckHealth() error {
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithDefaultServiceConfig(serviceConfig),
	}
	localhostAddress := fmt.Sprintf("0.0.0.0:%d", s.port)
	conn, err := grpc.Dial(localhostAddress, options...)
	if err != nil {
		log.Fatalf("grpc.Dial(%q): %v", localhostAddress, err)
	}

	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)

	healthRes, healthErr := checkHealth("MicroSastEngine", conn)
	if healthErr != nil {
		log.Printf("Health Check Failed: %v", healthErr)
		return healthErr
	}

	if healthRes.Status == healthpb.HealthCheckResponse_SERVING {
		log.Printf("End of Health Check! Status: %v", healthRes.Status)
		return nil
	}

	return fmt.Errorf("service not serving, status: %v", healthRes.Status)
}
