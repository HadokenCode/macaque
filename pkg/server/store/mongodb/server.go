package mongodb

import (
	"time"

	"github.com/golang/protobuf/ptypes"

	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/wildnature/macaque/pkg/logger"
	pbStore "github.com/wildnature/macaque/pkg/pb/store"
	"github.com/wildnature/macaque/pkg/store/mongodb"
	context "golang.org/x/net/context"
)

//Server structure
type Server struct {
}

func toTime(ts *timestamp.Timestamp) *time.Time {
	t, err := ptypes.Timestamp(ts)
	if err != nil {
		return nil
	}
	return &t
}

//Create function
func (s *Server) Create(ctx context.Context, in *pbStore.SchedulerEntity) (*google_protobuf.Empty, error) {
	logger.Infof("Retrieving request %v", in)
	err := mongodb.SaveScheduler(&mongodb.Scheduler{
		ID:          in.GetId().Id,
		Name:        in.GetName(),
		Description: in.GetDescription(),
		Expression:  in.GetExpression(),
		Labels:      in.GetLabels(),
		Properties:  in.GetProperties(),
		Status:      pbStore.SchedulerStatus_name[int32(in.GetStatus())],
		Type:        pbStore.SchedulerType_name[int32(in.GetType())],
		StartDate:   toTime(in.GetStartDate()),
		EndDate:     toTime(in.GetEndDate()),
	})
	if err != nil {
		return nil, err
	}
	return &google_protobuf.Empty{}, nil
}

//DeleteByID function
func (s *Server) DeleteByID(ctx context.Context, in *pbStore.EntityID) (*google_protobuf.Empty, error) {
	return nil, nil
}

//GetByID function
func (s *Server) GetByID(ctx context.Context, in *pbStore.EntityID) (*pbStore.SchedulerEntity, error) {
	schd, err := mongodb.GetScheduler(in.Id)
	if err != nil {
		return nil, err
	}
	return &pbStore.SchedulerEntity{Id: &pbStore.EntityID{Id: schd.ID}}, err
}
