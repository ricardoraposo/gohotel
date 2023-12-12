package db

import (
	"context"

	"github.com/ricardoraposo/gohotel/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const roomColl = "rooms"

type RoomStore interface {
	InsertRoom(context.Context, *data.Room) (*data.Room, error)
	GetRooms(context.Context, bson.M) ([]*data.Room, error)
}

type MongoRoomStore struct {
	client *mongo.Client
	coll   *mongo.Collection

	HotelStore
}

func NewMongoRoomStore(client *mongo.Client, hotelStore HotelStore) *MongoRoomStore {
	return &MongoRoomStore{
		client:     client,
		coll:       client.Database(DBNAME).Collection(roomColl),
		HotelStore: hotelStore,
	}
}

func (s *MongoRoomStore) GetRooms(ctx context.Context, filter bson.M) ([]*data.Room, error) {
	res, err := s.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var rooms []*data.Room
	if err := res.All(ctx, &rooms); err != nil {
		return nil, err
	}
	return rooms, err
}

func (s *MongoRoomStore) InsertRoom(ctx context.Context, room *data.Room) (*data.Room, error) {
	res, err := s.coll.InsertOne(ctx, room)
	if err != nil {
		return nil, err
	}

	room.ID = res.InsertedID.(primitive.ObjectID)

	// update the hotel with this room id
	update := bson.M{"$push": bson.M{"rooms": room.ID}}

	if err := s.HotelStore.UpdateHotel(ctx, room.HotelId, update); err != nil {
		return nil, err
	}

	return room, nil
}
