package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type MongoDB struct {

}

func NewMongoClient(config *Config) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//TODO
	client, err := mongo.Connect(ctx, options.Client().SetHosts([]string{config.Mongodb.Host}))
	if err != nil{
		log.Printf("Connect Database err #%v", err)
		return nil, err
	}
	return client.Database("economic_dev"), nil
}

