package main

import (
	"encoding/json"
	"fmt"
	conn "gitlab.com/stream/activity/connections"
	"gitlab.com/stream/activity/models"
	"gitlab.com/stream/activity/server"
	pb "gitlab.com/stream/buffers/profile"
	"google.golang.org/grpc"
	"net"
)

var cfg *models.Config

func init() {
	//Initialize all the configuraion
	cfg = &models.Config{

		//Server configuration
		ServerPort: "8000",

		//Couchbase configuration
		DBHost:         "localhost",
		DBPort:         "8091",
		NoSqlUser:      "Administrator",
		NoSqlPassword:  "nirfapurba",
		BucketName:     "socrates",
		BucketPassword: "",

		//Bolt configuration
		BoltHost:     "localhost",
		BoltPort:     "7687",
		BoltUser:     "neo4j",
		BoltPassword: "n4j",
	}
}

func main() {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.ServerPort))
	if err != nil {
		errorMessage(err, "Listenning")
	}

	nosql, err := conn.GetDBConnection(cfg)
	if err != nil {
		errorMessage(err, "Couchbase Init")
	}

	bolt, err := conn.GetBoltConnection(cfg)
	if err != nil {
		errorMessage(err, "Bolt Init")
	}

	activityServer, err := server.GetActivityServer(cfg, bolt, nosql)
	if err != nil {
		errorMessage(err, "Get Profile Server")
	}

	rpcServer := grpc.NewServer()
	pb.RegisterProfileServiceServer(rpcServer, activityServer)

	fmt.Println(string(fmtConfig(cfg)))

	err = rpcServer.Serve(listener)
	if err != nil {
		errorMessage(err, "Serve")
	}
}

func errorMessage(err error, context string) error {
	message := fmt.Sprintf("Service :: Profile :: %s :: %s", context, err.Error())
	fmt.Println(message)
	return err
}

func fmtConfig(cfg *models.Config) []byte {
	str, err := json.MarshalIndent(cfg, " ", " ")
	if err != nil {
		return nil
	}
	return str
}
