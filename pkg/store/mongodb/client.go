package mongodb

import (
	"time"

	"github.com/wildnature/macaque/pkg/configuration"
	constant "github.com/wildnature/macaque/pkg/constant/store/mongodb"
	"github.com/wildnature/macaque/pkg/logger"
	mgo "gopkg.in/mgo.v2"
)

type document interface {
	collection() string
}

func getConfigValueOrPanic(key string) string {
	val, err := configuration.GetConfiguration().Get(key)
	if err != nil {
		panic(err)
	}
	return val
}

func buildDial() *mgo.DialInfo {
	db := getConfigValueOrPanic(constant.Database)
	address := getConfigValueOrPanic(constant.Address)
	db, err := configuration.GetConfiguration().Get(constant.Database)
	if err != nil {
		panic(err)
	}
	return &mgo.DialInfo{
		Addrs:    []string{address},
		Timeout:  5 * time.Second,
		Database: db,
	}
}

func findOne(dial *mgo.DialInfo, collection string, docID interface{}) (interface{}, error) {
	session, err := mgo.DialWithInfo(dial)
	if err != nil {
		logger.Errorf("\nError %+v\n", err)
		return nil, err
	}
	session.SetMode(mgo.Monotonic, true)
	coll := session.DB(dial.Database).C(collection)
	if err != nil {
		return nil, err
	}
	defer session.Close()
	var res interface{}
	err = coll.FindId(docID).One(res)
	return res, err
}

func insert(dial *mgo.DialInfo, doc document) error {
	session, err := mgo.DialWithInfo(dial)
	if err != nil {
		logger.Errorf("\nError %+v\n", err)
		return err
	}
	session.SetMode(mgo.Monotonic, true)
	coll := session.DB(dial.Database).C(doc.collection())
	if err != nil {
		return err
	}
	defer session.Close()
	return coll.Insert(doc)
}

//CheckStatus ensusres the mongodb status
func CheckStatus() error {
	dial := buildDial()
	sess, err := mgo.DialWithInfo(dial)
	if err != nil {
		logger.Errorf("\nError: %s \n", err.Error())
		return err
	}
	defer sess.Close()
	err = sess.Ping()
	if err != nil {
		logger.Errorf("\nError: %s \n", err.Error())
		return err
	}
	logger.Infof("MongoDB server is healthy.")
	return nil
}
