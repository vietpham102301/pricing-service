package infra

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func InitMongoDB() (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://booking:booking@cluster0.c00rput.mongodb.net/pricing-service-db"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client, nil
}
