package server

import (
	"context"
	"encoding/json"
	"fmt"
	conn "gitlab.com/stream/activity/connections"
	"gitlab.com/stream/activity/models"
	"gitlab.com/stream/buffers/common"
	pb "gitlab.com/stream/buffers/profile"
	"gopkg.in/couchbase/gocb.v1"
)

const ProfileType = "user"
const PostType = "post"

type ActivityServer struct {
	cfg       *models.Config
	bolt      *conn.BoltConnection
	couchbase *conn.DBConnection
}

func GetActivityServer(cfg *models.Config, n4j *conn.BoltConnection, cb *conn.DBConnection) (*ActivityServer, error) {
	return &ActivityServer{
		cfg:       cfg,
		bolt:      n4j,
		couchbase: cb,
	}, nil
}

func (s *ActivityServer) InsertProfile(ctx context.Context, in *pb.ProfileInfo) (*pb.ProfileInfo, error) {
	/*
		Before insert we need to GetProfile once to check whether the user exists or not
		If the user is not found then gocb will return ERROR_KEY_NOT_FOUND . Using that error
		we'll insert the profile and after inserting it return it by doing a GetProfile.
	*/
	profile, err := s.GetProfile(ctx, &common.Request{Id: in.Id, Type: ProfileType})
	if err != nil {
		if err == gocb.ErrKeyNotFound {
			key := fmt.Sprintf("%s::%s", ProfileType, in.Id)
			_, err = s.couchbase.Insert(key, in, 0)
			if err != nil {
				return nil, errorMessage(err, "Insert")
			}
			return s.GetProfile(ctx, &common.Request{Id: in.Id, Type: ProfileType})
		}
	}

	// If the profile is already in the database return it with error nil
	if profile.Id != "" {
		return profile, nil
	}

	//Otherwise return an empty profile
	return &pb.ProfileInfo{}, nil
}

func (s *ActivityServer) GetProfile(ctx context.Context, in *common.Request) (*pb.ProfileInfo, error) {

	//Generating the key of the document
	key := fmt.Sprintf("%s::%s", ProfileType, in.Id)

	// Creating the profile type variable where the profile data will be stored after pulling from couchbase
	var data pb.ProfileInfo
	_, err := s.couchbase.Get(key, &data)
	if err != nil {
		return nil, errorMessage(err, "GET")
	}

	return &data, nil
}

func (s *ActivityServer) UpdateProfile(ctx context.Context, in *common.Request) (*pb.ProfileInfo, error) {
	return nil, nil
}

func (s *ActivityServer) FindUserFriends(uid string, ch chan [][]models.Node) {

	//Generating the query to find friends of the associated user id
	query := fmt.Sprintf("MATCH (n:User)-[:FRIENDS]-(m:User{id:\"%s\"}) return n", uid)

	// Execute the query in neo4j
	result, _, _, err := s.bolt.Connection.QueryNeoAll(query, nil)
	if err != nil {
		errorMessage(err, "Error getting data from neo4j server")
	}

	// Serialize the result gotten from neo4j
	b, err := json.Marshal(result)
	if err != nil {
		errorMessage(err, "Error Marshalling data into bytes")
	}

	//Deserialize bytes array into our temporary friends model
	var friends [][]models.Node
	err = json.Unmarshal(b, &friends)
	if err != nil {
		errorMessage(err, "Error Marshalling data into bytes")
	}

	ch <- friends
}

func (s *ActivityServer) InsertTravelPost(ctx context.Context, in *pb.TravelPost) (*pb.TravelPost, error) {
	/*
		Before insert we need to GetProfile once to check whether the user exists or not
		If the user is not found then gocb will return ERROR_KEY_NOT_FOUND . Using that error
		we'll insert the profile and after inserting it return it by doing a GetProfile.
	*/
	post, err := s.GetTravelPost(ctx, &common.Request{Id: in.PostId, Type: PostType})
	if err != nil {
		if err == gocb.ErrKeyNotFound {
			key := fmt.Sprintf("%s::%s", PostType, in.PostId)
			_, err = s.couchbase.Insert(key, in, 0)
			if err != nil {
				return nil, errorMessage(err, "Insert")
			}
			return s.GetTravelPost(ctx, &common.Request{Id: in.PostId, Type: PostType})
		}
	}

	// If the profile is already in the database return it with error nil
	if post.PostId != "" {
		return post, nil
	}

	//Otherwise return an empty profile
	return &pb.TravelPost{}, nil
}

func (s *ActivityServer) GetTravelPost(ctx context.Context, in *common.Request) (*pb.TravelPost, error) {
	//Generating the key of the document
	key := fmt.Sprintf("%s::%s", "post", in.Id)

	// Creating the profile type variable where the profile data will be stored after pulling from couchbase
	var data pb.TravelPost
	_, err := s.couchbase.Get(key, &data)
	if err != nil {
		return nil, errorMessage(err, "GET")
	}

	return &data, nil
}

func errorMessage(err error, context string) error {
	fmt.Printf("Error :: %s :: %s\n", context, err.Error())
	return err
}
