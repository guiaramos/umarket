package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConfig defines a config options for MongoDB
type MongoConfig struct {
	Timeout  int
	DBname   string
	Username string
	Password string
	Host     string
	Port     string
}

func Connect(c MongoConfig) (*mongo.Database, error) {
	uriPattern := "mongodb://%v:%v@%v:%v"
	if c.Username == "" {
		uriPattern = "mongodb://%s%s%v:%v"
	}

	uri := fmt.Sprintf(uriPattern,
		c.Username,
		c.Password,
		c.Host,
		c.Port,
	)

	clientOpts := options.Client().ApplyURI(uri)

	client, err := mongo.NewClient(clientOpts)
	if err != nil {
		return nil, err
	}

	ctx, e := context.WithTimeout(context.Background(), time.Duration(c.Timeout)*time.Second)
	if e != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client.Database(c.DBname), err
}
