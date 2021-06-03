package aws

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"sync"
)

var initLock sync.Mutex
var conn *mongo.Client

func initDb(config *ObjectStoreConfig) *mongo.Client {
	defer initLock.Unlock()

	initLock.Lock()

	credentials := options.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		Username:      config.DB.User,
		Password:      config.DB.Password,
	}

	var err error
	conn, err = mongo.Connect(
		context.TODO(),
		options.Client().SetAuth(credentials).ApplyURI(config.DB.URI),
	)
	if err != nil {
		log.Fatalf("unable to connect to object store database: %v", err)
	}

	err = conn.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatalf("unable to connect to object store database: %v", err)
	}

	log.Infof("succesfully connected to object store database %s", config.DB.DatabaseName)

	return conn
}
