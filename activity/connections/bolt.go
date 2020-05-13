package connectiors

import (
	"fmt"
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"github.com/pkg/errors"
	"gitlab.com/stream/activity/models"
)

type BoltConnection struct {
	Connection bolt.Conn
}

func GetBoltConnection(cfg *models.Config) (*BoltConnection, error) {

	driver := bolt.NewDriver()
	conn, err := driver.OpenNeo(fmt.Sprintf("bolt://%s:%s@localhost:%s", cfg.BoltUser, cfg.BoltPassword, cfg.BoltPort))
	if err != nil {
		return nil, errors.New("Neo4j Connection not established .")
	}

	//defer conn.Close()
	return &BoltConnection{conn}, nil
}
