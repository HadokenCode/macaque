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

func setUpMgoConn() *mgo.DialInfo {
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


func insert(dial *mgo.DialInfo, doc document, queryFn func(doc interface{}, c *mgo.Collection) (interface{}, error)) (interface{},error) {
	session, err := mgo.DialWithInfo(dial)
	if err != nil {
		logger.Errorf("\nError %+v\n", err)
		return nil,err
	}
	
	session.SetMode(mgo.Monotonic, true)
	coll := session.DB(dial.Database).C(doc.collection())
	if err != nil {
		return nil,err
	}
	defer session.Close()
	logger.Infof("\n Querying fn %+v\n", coll)
	res, err := queryFn(doc, coll)
	if err != nil {
		return nil,err
	}
	logger.Infof("\nSetting result %+v\n", res)
	return res,nil
}


//CheckStatus ensusres the mongodb status
func CheckStatus() error{
	dial:=setUpMgoConn()
	logger.Infof("\n - Dial details: %+v \n",dial)
	sess, err := mgo.DialWithInfo(dial)
	if err != nil {
		logger.Errorf("\nError: %s \n",err.Error())
		return err
	}
	defer sess.Close()
	err = sess.Ping()
	if err != nil {
		logger.Errorf("\nError: %s \n",err.Error())
		return err
	}
	logger.Infof("MongoDB server is healthy.")
	return nil
}