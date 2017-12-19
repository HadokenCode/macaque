package mongodb

import (
	"time"

	"github.com/wildnature/macaque/pkg/logger"
	pbStore "github.com/wildnature/macaque/pkg/pb/store"
	"gopkg.in/mgo.v2/bson"
)

const (
	schedulerColl string = "scheduler"
	statusField          = "status"
)

//Scheduler structure to be persisted
type Scheduler struct {
	ID          string            `bson:"_id,omitempty"`
	Name        string            `bson:"name,omitempty"`
	Description string            `bson:"description,omitempty"`
	Expression  string            `bson:"expression,omitempty"`
	Labels      []string          `bson:"labels,omitempty"`
	Type        string            `bson:"type,omitempty"`
	Status      string            `bson:"status,omitempty"`
	StartDate   *time.Time        `bson:"startDate,omitempty"`
	EndDate     *time.Time        `bson:"endDate,omitempty"`
	Properties  map[string]string `bson:"properties,omitempty"`
}

func (s *Scheduler) collection() string {
	return schedulerColl
}

//SaveScheduler method
func SaveScheduler(scheduler *Scheduler) error {
	mgoDial := buildDial()
	return insert(mgoDial, scheduler)
}

func toScheduler(res bson.M) (*Scheduler, error) {
	schd := &Scheduler{}
	bsonBytes, _ := bson.Marshal(res)
	err := bson.Unmarshal(bsonBytes, schd)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return schd, err
}

//GetScheduler method
func GetScheduler(schedulerID string) (*Scheduler, error) {
	mgoDial := buildDial()
	logger.Infof("\nSearching schedule for ID %s\n", schedulerID)
	res, err := findOne(mgoDial, schedulerColl, schedulerID)
	if err == nil {
		return toScheduler(res)
	}
	logger.Errorf("\nUnexpected error %s\n", err.Error())
	return nil, err
}

//UpdateScheduler function
func UpdateScheduler(schedulerID string, fields map[string]interface{}) error {
	mgoDial := buildDial()
	logger.Infof("\nPartial updating schedule for ID %s\n", schedulerID)
	return partialUpdate(mgoDial, schedulerColl, schedulerID, fields)
}

//UpdateSchedulerWithBytes function
func UpdateSchedulerWithBytes(schedulerID string, fields bson.Binary) error {
	mgoDial := buildDial()
	logger.Infof("\nPartial updating schedule for ID %s\n", schedulerID)
	return partialUpdateWithBytes(mgoDial, schedulerColl, schedulerID, fields)
}

//DeleteScheduler function, actually this only makes a change in scheduler status
func DeleteScheduler(schedulerID string) error {
	return UpdateScheduler(schedulerID, map[string]interface{}{statusField: pbStore.SchedulerStatus_name[int32(pbStore.SchedulerStatus_DELETED)]})
}
