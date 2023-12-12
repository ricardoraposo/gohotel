package db

import (
	"context"

	"github.com/ricardoraposo/gohotel/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const hotelColl = "hotels"

type HotelStore interface {
	InsertHotel(context.Context, *data.Hotel) (*data.Hotel, error)
	UpdateHotel(context.Context, primitive.ObjectID, bson.M) error
	GetHotels(context.Context, bson.M) ([]*data.Hotel, error)
	GetHotel(context.Context, primitive.ObjectID) (*data.Hotel, error)
}

type MongoHotelStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoHotelStore(client *mongo.Client) *MongoHotelStore {
	return &MongoHotelStore{
		client: client,
		coll:   client.Database(DBNAME).Collection(hotelColl),
	}
}

func (s *MongoHotelStore) GetHotels(ctx context.Context, filter bson.M) ([]*data.Hotel, error) {
	cur, err := s.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var hotels []*data.Hotel
	if err := cur.All(ctx, &hotels); err != nil {
		return nil, err
	}
	return hotels, nil
}

func (s *MongoHotelStore) GetHotel(ctx context.Context, id primitive.ObjectID) (*data.Hotel, error) {
	var hotel *data.Hotel
	if err := s.coll.FindOne(ctx, bson.M{"_id": id}).Decode(&hotel); err != nil {
		return nil, err
	}
	return hotel, nil
}

func (s *MongoHotelStore) InsertHotel(ctx context.Context, hotel *data.Hotel) (*data.Hotel, error) {
	res, err := s.coll.InsertOne(ctx, hotel)
	if err != nil {
		return nil, err
	}

	hotel.ID = res.InsertedID.(primitive.ObjectID)
	return hotel, nil
}

func (s *MongoHotelStore) UpdateHotel(ctx context.Context, id primitive.ObjectID, update bson.M) error {
	_, err := s.coll.UpdateOne(ctx, bson.M{"_id": id}, update)
	return err
}
