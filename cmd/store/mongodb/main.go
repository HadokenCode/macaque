package main

import (
	"net"

	"github.com/wildnature/macaque/pkg/configuration"
	constant "github.com/wildnature/macaque/pkg/constant/store/mongodb"
	"github.com/wildnature/macaque/pkg/env"
	"github.com/wildnature/macaque/pkg/logger"
	pbStore "github.com/wildnature/macaque/pkg/pb/store"
	"github.com/wildnature/macaque/pkg/server/store/mongodb"
	grpc "google.golang.org/grpc"
	mgo "gopkg.in/mgo.v2"
)

const (
	protocol                 string = "tcp"
	address                         = ":5051"
	mongoAddress                    = "localhost:27017"
	mongoDatabase                   = "macaque"
	mongoSchedulerCollection        = "schedulers"
)

func init() {
	c := configuration.GetConfiguration()
	c.Set(constant.Address, env.GetEnv(constant.Address, mongoAddress))
	c.Set(constant.Database, env.GetEnv(constant.Database, mongoDatabase))
	c.Set(constant.SchedulerCollection, env.GetEnv(constant.SchedulerCollection, mongoSchedulerCollection))
	logger.Infof("Configuration: '%+v' \n", c)
}

func getGrpcServerConfiguration() []grpc.ServerOption {
	return []grpc.ServerOption{}
}

func main() {
	mgo.SetDebug(true)
	go mongodb.RunHealthServer()
	conn, err := net.Listen(protocol, address)
	if err != nil {
		logger.Errorf("Error: '%s' \n", err.Error())
		return
	}
	grpcServer := grpc.NewServer(getGrpcServerConfiguration()...)
	pbStore.RegisterStoreServiceServer(grpcServer, &mongodb.Server{})
	logger.Infof("\nRunning server on %s", address)
	grpcServer.Serve(conn)
}
