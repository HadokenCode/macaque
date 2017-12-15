package mongodb

import (
	constant "github.com/wildnature/macaque/pkg/constant/store/mongodb"
	"github.com/wildnature/macaque/pkg/logger"
	mgo "gopkg.in/mgo.v2"
)

//Scheduler structure
type Scheduler struct {
	Name        string
	ID          string
	Description string
	Expression  string
	Labels      []string
}

func (s *Scheduler) collection() string {
	return "scheduler"
}

//SaveScheduler method
func SaveScheduler(scheduler *Scheduler) error {
	schedulerCollection := getConfigValueOrPanic(constant.SchedulerCollection)
	logger.Infof("Getting collection name %s", schedulerCollection)
	mgoDial := buildDial()
	logger.Infof("Connection was already set up %v", mgoDial)
	_, err := run(mgoDial, scheduler, insertSchedulerEntity)
	return err
}

//SaveScheduler method
func GetScheduler(schedulerID string) (*Scheduler,error) {
	schedulerCollection := getConfigValueOrPanic(constant.SchedulerCollection)
	logger.Infof("Getting collection name %s", schedulerCollection)
	mgoDial := buildDial()
	logger.Infof("Connection was already set up %v", mgoDial)
	return run(mgoDial, schedulerID, getSchedulerEntity)
}

func insertSchedulerEntity(scheduler interface{}, c *mgo.Collection) (interface{}, error) {
	logger.Infof("Inserting scheduler %v into the mongo database", scheduler)
	return nil, c.Insert(scheduler)
}

func getSchedulerEntity(schedulerID string, c *mgo.Collection) (*Scheduler, error) {
	logger.Infof("Getting scheduler from", schedulerID)
	return nil, c.FindId(schedulerID)
}
