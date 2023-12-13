package db

import (
	"context"

	"github.com/ricardoraposo/gohotel/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const bookingColl = "bookings"

type BookingStore interface {
	InsertBooking(context.Context, *data.Booking) (*data.Booking, error)
	GetBookings(ctx context.Context, filter bson.M) ([]*data.Booking, error)
}

type MongoBookingStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoBookingStore(client *mongo.Client) *MongoBookingStore {
	return &MongoBookingStore{
		client: client,
		coll:   client.Database(DBNAME).Collection(bookingColl),
	}
}

func (s *MongoBookingStore) GetBookings(ctx context.Context, filter bson.M) ([]*data.Booking, error) {
    res, err := s.coll.Find(ctx, filter)
    if err != nil {
        return nil, err
    }

    var bookings []*data.Booking
    if err := res.All(ctx, &bookings); err != nil {
        return nil, err
    }
    return bookings, nil
}

func (s *MongoBookingStore) InsertBooking(ctx context.Context, booking *data.Booking) (*data.Booking, error) {
	res, err := s.coll.InsertOne(ctx, booking)
	if err != nil {
		return nil, err
	}
	booking.ID = res.InsertedID.(primitive.ObjectID)
	return booking, nil
}
