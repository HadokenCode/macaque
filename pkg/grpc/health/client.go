package health

import (
	"errors"

	pbHealth "github.com/wildnature/macaque/pkg/pb/health"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

//Check methods to ensure health service
func Check(serverAddress string, service string) error {
	conn, _ := grpc.Dial(serverAddress, grpc.WithInsecure())
	defer conn.Close()
	client := pbHealth.NewHealthClient(conn)
	md := metadata.Pairs("sec-token", "A1B2")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	res, err := client.Check(ctx, &pbHealth.HealthCheckRequest{Service: service})
	if err != nil {
		return err
	}
	if res.GetStatus() != pbHealth.HealthCheckResponse_SERVING {
		return errors.New("Service is not up and running")
	}
	return nil
}
