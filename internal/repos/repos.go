package repos

import "go.mongodb.org/mongo-driver/mongo"

type MongoRepo struct {
	mgClient *mongo.Client
}

func NewMongoDBRepo(mgClient *mongo.Client) IRepo {
	return &MongoRepo{
		mgClient: mgClient,
	}
}

func (m *MongoRepo) Pricing() IPricingRepo {
	return NewPricingNoSQLRepo(m.mgClient)
}
