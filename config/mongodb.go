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

func NewMongoDatabase(config *Config) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	credential := options.Credential{
		Username: config.Mongodb.UserName,
		Password: config.Mongodb.Password,
		AuthSource: config.Mongodb.AuthenticationDatabase,
	}
	mongoOptions := options.Client().SetHosts([]string{config.Mongodb.Host}).SetAuth(credential)

	defer cancel()
	client, err := mongo.Connect(ctx, mongoOptions)
	if err != nil{
		log.Printf("Connect Database err #%v", err)
		return nil, err
	}
	return client.Database(config.Mongodb.Database), nil
}

