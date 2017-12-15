package mongodb

import (
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	"github.com/wildnature/macaque/pkg/logger"
	pbStore "github.com/wildnature/macaque/pkg/pb/store"
	"github.com/wildnature/macaque/pkg/store/mongodb"
	context "golang.org/x/net/context"
)

//Server structure
type Server struct {
}

//SaveScheduler function
func (s *Server) SaveScheduler(ctx context.Context, in *pbStore.SchedulerEntity) (*pbStore.EntityID, error) {
	logger.Infof("Retrieving request %v", in)
	generatedID := "generated-id"
	return &pbStore.EntityID{Id: generatedID}, mongodb.SaveScheduler(&mongodb.Scheduler{
		ID:     generatedID,
		Name:   "bla-bla-bla",
		Labels: in.GetLabels(),
	})
}

//DeleteSchedulerByID function
func (s *Server) DeleteSchedulerByID(ctx context.Context, in *pbStore.EntityID) (*google_protobuf.Empty, error) {
	return nil, nil
}

//GetSchedulerByID function
func (s *Server) GetSchedulerByID(ctx context.Context, in *pbStore.EntityID) (*pbStore.SchedulerEntity, error) {
	schd,err:=mongodb.GetScheduler(in.Id)
	return pbStore.SchedulerEntity{
		Description: schd.Description,
	},err
}
