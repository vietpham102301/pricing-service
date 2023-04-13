package repos

import (
	"booking-service/internal/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type PricingNoSQLRepo struct {
	mongoClient *mongo.Client
}

func NewPricingNoSQLRepo(mongoClient *mongo.Client) IPricingRepo {
	return &PricingNoSQLRepo{
		mongoClient: mongoClient,
	}
}

func (p PricingNoSQLRepo) GetPrice(jobType int) (*models.Price, error) {
	priceCollection := p.mongoClient.Database("pricing-service-db").Collection("prices")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	price := &models.Price{}
	err := priceCollection.FindOne(ctx, bson.M{"id": jobType}).Decode(price)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		log.Printf("getting price fail with err: %v", err.Error())
		return nil, err
	}

	return price, nil
}
