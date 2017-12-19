package mongodb

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"time"

	"github.com/golang/protobuf/ptypes"
	google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
	"github.com/wildnature/macaque/pkg/grpc/health"
	"github.com/wildnature/macaque/pkg/logger"
	pb "github.com/wildnature/macaque/pkg/pb/store"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	address       = "localhost:5051"
	healthAddress = "localhost:5052"
)

// TestMain before each test
func TestMain(m *testing.M) {
	rand.Seed(time.Now().UnixNano())
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

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func now(seconds int64) *google_protobuf.Timestamp {
	ts := ptypes.TimestampNow()
	ts.Seconds += seconds
	return ts
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
			description: "I - Happy path",
			content: &pb.SchedulerEntity{
				Id:          &pb.EntityID{Id: randomStr(32)},
				Name:        "Scheduler Test I",
				Description: "This is a testing purpose scheduler",
				Labels:      []string{"scheduler", "demo", "testing"},
				Status:      pb.SchedulerStatus_CREATED,
				Type:        pb.SchedulerType_PRIVATE,
				Expression:  "0/15 * * * * *",
				Properties: map[string]string{
					"a": "b",
					"c": "d",
				},
				StartDate: now(0),
				EndDate:   now(10000),
			},
			expectedError: nil,
		},
	}

	for _, c := range cases {
		md := metadata.Pairs("sec-token", "A1B2")
		ctx := metadata.NewOutgoingContext(context.Background(), md)
		result, err := client.Create(ctx, c.content)
		logger.Debugf("\n\n%s: \n", c.description)
		if err != nil {
			t.Error(err.Error())
		}
		if c.expectedError == nil {
			assert.NotEmpty(t, result)
		} else {
			assert.EqualValues(t, c.expectedError, err)
		}
		scheduler, err := client.GetByID(ctx, &pb.EntityID{Id: c.content.GetId().Id})
		if err != nil {
			t.Error(err.Error())
		}
		fmt.Println(scheduler)

		fmt.Println("\n\n=============\n\n")
		schdUpdate := &pb.SchedulerEntity{
			Id:   c.content.GetId(),
			Name: "New name",
		}
		schdUpdate.Name = " New name"
		_, err = client.Update(ctx, schdUpdate)
		if err != nil {
			t.Error(err.Error())
		}
		fmt.Println("\n\n=============\n\n")
		_, err = client.DeleteByID(ctx, &pb.EntityID{Id: c.content.GetId().Id})
		if err != nil {
			t.Error(err.Error())
		}
		scheduler, err = client.GetByID(ctx, &pb.EntityID{Id: c.content.GetId().Id})
		if err != nil {
			t.Error(err.Error())
		}
		fmt.Println(scheduler)
	}

}
