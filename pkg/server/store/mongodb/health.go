package mongodb

import (
	"net"

	"golang.org/x/net/context"

	"github.com/wildnature/macaque/pkg/logger"
	pbHealth "github.com/wildnature/macaque/pkg/pb/health"
	"github.com/wildnature/macaque/pkg/store/mongodb"
	"google.golang.org/grpc"
)

const (
	protocol   string = "tcp"
	address    string = "0.0.0.0:5052"
	srvMongoDB        = "mongodb"
)

func getGrpcServerConfiguration() []grpc.ServerOption {
	return []grpc.ServerOption{}
}

//RunHealthServer method
func RunHealthServer() {
	conn, err := net.Listen(protocol, address)
	if err != nil {
		logger.Errorf("Error: '%s' \n", err.Error())
		return
	}
	logger.Infof("Running health endpoint on: '%s' \n", address)
	grpcServer := grpc.NewServer(getGrpcServerConfiguration()...)
	pbHealth.RegisterHealthServer(grpcServer, &healthServer{})
	grpcServer.Serve(conn)
}

type healthServer struct {
}

func (s *healthServer) Check(ctx context.Context, in *pbHealth.HealthCheckRequest) (*pbHealth.HealthCheckResponse, error) {
	logger.Infof("\n - Checking status for service %s\n", in.Service)
	switch in.Service {
	case srvMongoDB:
		err := mongodb.CheckStatus()
		if err != nil {
			return &pbHealth.HealthCheckResponse{
				Status: pbHealth.HealthCheckResponse_NOT_SERVING,
			}, err
		}
		return &pbHealth.HealthCheckResponse{
			Status: pbHealth.HealthCheckResponse_SERVING,
		}, nil

	}
	return &pbHealth.HealthCheckResponse{
		Status: pbHealth.HealthCheckResponse_UNKNOWN,
	}, nil
}
