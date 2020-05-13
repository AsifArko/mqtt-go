package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	conn "gitlab.com/stream/activity/connections"
	"gitlab.com/stream/activity/helper"
	"gitlab.com/stream/activity/models"
	"gitlab.com/stream/activity/server"
	"gitlab.com/stream/buffers/common"
	pb "gitlab.com/stream/buffers/profile"
	"log"
	"net/url"
	"os"
	"testing"
	"time"
)

var activityServer *server.ActivityServer

type mockConfig struct {
	cfg            *models.Config
	nosql          *conn.DBConnection
	bolt           *conn.BoltConnection
	activityServer *server.ActivityServer
}

func TestProfileService(t *testing.T) {

	// Setting up the configuration parameters for making a mock config for testing purposes
	cf := &models.Config{
		//Server Configuration
		ServerHost: "localhost",
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

	//Get nosql connection
	nosql, err := conn.GetDBConnection(cf)
	if err != nil {
		t.Fatal(err.Error())
	}

	//Get bolt connection
	bolt, err := conn.GetBoltConnection(cf)
	if err != nil {
		t.Fatal(err.Error())
	}

	// Get the profile server in testing environment
	server, err := server.GetActivityServer(cf, bolt, nosql)
	if err != nil {
		t.Fatal(err.Error())
	}

	//Making a mock configuration for server method testing
	var mock mockConfig
	mock.cfg = cf
	mock.nosql = nosql
	mock.bolt = bolt
	mock.activityServer = server

	//Test the server methods

	//Test Get profile
	//testGetProfile(t , &mock)

	// Test Insert profile
	//testInsertProfile(t, &mock)

	//Test Find User friends
	//testFindFriends(t, &mock)

	testInsertTravelPost(t, &mock)

}

func testInsertProfile(t *testing.T, mockCfg *mockConfig) {

	prof := helper.GenerateSampleProfile()

	profile, err := mockCfg.activityServer.InsertProfile(context.Background(), prof)
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Logf("%+v\n", profile)
}

func testInsertTravelPost(t *testing.T, mockCfg *mockConfig) {

	post := &pb.TravelPost{
		Uid:    "4878",
		PostId: "00002",
		Post:   "Awesome Sajek",
	}

	response, err := mockCfg.activityServer.InsertTravelPost(context.Background(), post)
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Logf("%+v\n", response)

	b, err := json.Marshal(response)

	os.Setenv("CLOUDMQTT_URL", "tcp://localhost:1883")
	uri, err := url.Parse(os.Getenv("CLOUDMQTT_URL"))
	if err != nil {
		log.Fatal(err)
	}
	topic := "test"

	c := connect("4878", uri)

	c.Publish(topic, 0, false, string(b))

}

func testGetProfile(t *testing.T, mockCfg *mockConfig) {
	request := &common.Request{Id: "4567", Type: "user"}

	response, err := mockCfg.activityServer.GetProfile(context.Background(), request)
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(response)
}

func testGetTravelPost(t *testing.T, mockCfg *mockConfig) {
	request := &common.Request{Id: "00001", Type: "post"}

	response, err := mockCfg.activityServer.GetTravelPost(context.Background(), request)
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(response)
}

func testFindFriends(t *testing.T, mockCfg *mockConfig) {
	uid := "4878"

	// Initiating the channel to be passed in FindUserFriends routine
	n4jRes := make(chan [][]models.Node)
	// Calling the go routine with the newly created channel
	go mockCfg.activityServer.FindUserFriends(uid, n4jRes)

	// Declarization of the channel type variable to catch the channel data
	var userFriendNodes [][]models.Node
	userFriendNodes = <-n4jRes

	// After getting the user id's of the friend of the following user getting the detailed information of the user profiles from Couchbase
	for _, eachNode := range userFriendNodes {
		friend, err := mockCfg.activityServer.GetProfile(context.Background(), &common.Request{Id: eachNode[0].Properties.Id, Type: "user"})
		if err != nil {
			t.Fatal(err.Error())
		}

		t.Logf("%+v\n", friend)
	}
}

func connect(clientId string, uri *url.URL) mqtt.Client {

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))

	opts.SetClientID(clientId)

	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}
