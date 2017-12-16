package mongodb

import (
	"errors"
	"time"

	"github.com/wildnature/macaque/pkg/logger"
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
	return "scheduler"
}

//SaveScheduler method
func SaveScheduler(scheduler *Scheduler) error {
	mgoDial := buildDial()
	return insert(mgoDial, scheduler)
}

//GetScheduler method
func GetScheduler(schedulerID string) (*Scheduler, error) {
	mgoDial := buildDial()
	logger.Infof("\nSearching schedule for ID %s\n", schedulerID)
	res, err := findOne(mgoDial, "scheduler", schedulerID)
	if err == nil {
		s, ok := res.(*Scheduler)
		if !ok {
			return nil, errors.New("Invalid response from mongodb")
		}
		logger.Infof("\nResponse from mongodb is %+v\n", s)
		return s, nil
	}
	logger.Errorf("\nUnexpected error %s\n", err.Error())
	return nil, err
}
