package mongodb

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/wildnature/macaque/pkg/grpc/health"
	"github.com/wildnature/macaque/pkg/logger"
	pb "github.com/wildnature/macaque/pkg/pb/store"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)

const (
	address       = "localhost:5051"
	healthAddress = "localhost:5052"
)

// TestMain before each test
func TestMain(m *testing.M) {
	tries := 3
	ready := false
	for !ready && tries > 0 {
		logger.Debug("Checking is service is up")
		err := health.Check(healthAddress, "mongodb")
		if err != nil {
			logger.Errorf("\nError checking service health %v \n", err)
			tries--
			time.Sleep(3 * time.Second)
		} else {
			ready = true
		}
	}
	if ready {
		logger.Debug("Running test")
		code := m.Run()
		os.Exit(code)
	} else {
		logger.Debug("The service is unavailable to be tested.")
		os.Exit(1)
	}

}

func TestCreateScheduler(t *testing.T) {
	logger.Debug("Running integration test..")
	conn, _ := grpc.Dial(address, grpc.WithInsecure())
	defer conn.Close()
	logger.Debug("Creating connection")
	client := pb.NewStoreServiceClient(conn)
	cases := []struct {
		description   string
		content       *pb.SchedulerEntity
		expectedError error
	}{
		{
			description:   "I - Happy path",
			content:       &pb.SchedulerEntity{},
			expectedError: nil,
		},
	}

	for _, c := range cases {
		md := metadata.Pairs("sec-token", "A1B2")
		ctx := metadata.NewOutgoingContext(context.Background(), md)
		result, err := client.SaveScheduler(ctx, c.content)
		logger.Debugf("\n\n%s: \n", c.description)
		if err != nil {
			fmt.Println(err.Error())
		}
		if c.expectedError == nil {
			assert.NotEmpty(t, result)
		} else {
			assert.EqualValues(t, c.expectedError, err)
		}
	}

}
