package models

type Config struct {

	//Serve configuration
	ServerHost string
	ServerPort string

	//Couchbase configuration
	DBHost         string
	DBPort         string
	NoSqlUser      string
	NoSqlPassword  string
	BucketName     string
	BucketPassword string

	//Bolt configuration
	BoltHost     string
	BoltPort     string
	BoltUser     string
	BoltPassword string

	//Redis configuration
	RedisHost     string
	RedisPort     string
	RedisUser     string
	RedisPassword string
}
