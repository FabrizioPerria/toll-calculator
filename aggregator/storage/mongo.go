package storage

import (
	"context"
	"fmt"

	"github.com/fabrizioperria/toll/shared/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStorage struct {
	client     *mongo.Client
	collection *mongo.Collection
	context    context.Context
}

func NewMongoStorage(url string) Storer {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(url))
	if err != nil {
		panic(err)
	}
	dbName := "toll"
	distanceCollection := "distances"
	return &MongoStorage{
		client:     client,
		collection: client.Database(dbName).Collection(distanceCollection),
		context:    context.Background(),
	}
}

func (m *MongoStorage) Store(distance types.Distance) error {
	opts := options.Update().SetUpsert(true)
	_, err := m.collection.UpdateOne(context.Background(), bson.M{"obu_id": distance.ObuId}, bson.M{
		"$inc": bson.M{"distance": distance.Value},
		"$set": bson.M{"obu_id": distance.ObuId},
	}, opts)

	return err
}

type mongoDistance struct {
	ObuID    string  `bson:"obu_id"`
	Distance float64 `bson:"distance"`
}

func (m *MongoStorage) Get(obuID string) (float64, error) {
	filter := bson.M{"obu_id": obuID}
	result := m.collection.FindOne(m.context, filter)
	if result.Err() != nil {
		return -1, fmt.Errorf("obuID %s not found", obuID)
	}
	var distance mongoDistance
	err := result.Decode(&distance)
	if err != nil {
		return -1, err
	}

	return distance.Distance, nil
}
