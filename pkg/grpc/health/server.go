package health

import (
	"net"

	context "golang.org/x/net/context"

	"github.com/wildnature/macaque/pkg/logger"
	pbHealth "github.com/wildnature/macaque/pkg/pb/health"
	"google.golang.org/grpc"
)

const (
	protocol string = "tcp"
	address  string = "0.0.0.0:5052"
)

func getGrpcServerConfiguration() []grpc.ServerOption {
	return []grpc.ServerOption{}
}

//Run method
func Run() {

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
	return &pbHealth.HealthCheckResponse{
		Status: pbHealth.HealthCheckResponse_SERVING,
	}, nil
}
