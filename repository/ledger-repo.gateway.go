package repository

import (
	"bank/domain"
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepo struct {
	client  *mongo.Client
	db      string
	timeout time.Duration
}

func InitMongoClient(mongoURI string, timeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))

	if err != nil {
		return nil, err
	}

	return client, nil
}

func CreateMongoRepo(mongoURI, mongoDB string, timeout int) (domain.Repository, error) {
	mongoClient, err := InitMongoClient(mongoURI, timeout)

	repo := &MongoRepo{
		client:  mongoClient,
		db:      mongoDB,
		timeout: time.Duration(timeout) * time.Second,
	}

	if err != nil {
		return nil, errors.Wrap(err, "mongo client error")
	}

	return repo, nil
}

func (r *MongoRepo) Insert(ledger *domain.Ledger) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	collection := r.client.Database(r.db).Collection("ledger")

	_, err := collection.InsertOne(
		ctx,
		bson.M{
			"entityId":       ledger.EntityId,
			"timelineSerial": ledger.TimelineSerial,
			"amount":         ledger.Amount,
			"eventType":      ledger.EventType,
			"event":          ledger.Event,
			"timestamp":      ledger.Timestamp,
		})

	if err != nil {
		return errors.Wrap(err, "mongo client error")
	}

	return nil
}

func (r *MongoRepo) FindOne(entityId string) (*domain.Ledger, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	ledger := &domain.Ledger{}

	collection := r.client.Database(r.db).Collection("ledger")
	filter := bson.M{"entityId": entityId}

	err := collection.FindOne(ctx, filter).Decode(ledger)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("Error Finding a catalogue item")
		}
		return nil, errors.Wrap(err, "repository research")
	}

	return ledger, nil
}

func (r *MongoRepo) FindAll() ([]*domain.Ledger, error) {

	var items []*domain.Ledger

	collection := r.client.Database(r.db).Collection("ledger")
	cur, err := collection.Find(context.Background(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	for cur.Next(context.TODO()) {

		var item domain.Ledger
		if err := cur.Decode(&item); err != nil {
			log.Fatal(err)
			return nil, err
		}
		items = append(items, &item)
	}
	return items, nil

}
func (r *MongoRepo) Delete(code string) error {

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	filter := bson.M{"code": code}

	collection := r.client.Database(r.db).Collection("ledger")
	_, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	return nil

}
