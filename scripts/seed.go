package main

import (
	"context"
	"log"

	"github.com/ricardoraposo/gohotel/data"
	"github.com/ricardoraposo/gohotel/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	roomStore  *db.MongoRoomStore
	hotelStore *db.MongoHotelStore
	userStore *db.MongoUserStore
	ctx        = context.Background()
)

func seedUser(fname, lname, email string) {
	user, err := data.NewUserFromRequest(data.CreateUserRequest{
		FirstName: fname,
		LastName:  lname,
		Email:     email,
		Password:  "superSecretFodase",
	})
	if err != nil {
		log.Fatal(err)
	}
    _, err = userStore.InsertUser(ctx, user)
    if err != nil {
		log.Fatal(err)
    }
}

func seedHotel(name string, location string, rating int) {
	hotel := data.Hotel{
		Name:     name,
		Location: location,
		Rooms:    []primitive.ObjectID{},
		Rating:   rating,
	}

	rooms := []data.Room{
		{
			Size:  "small",
			Price: 69.99,
		},
		{
			Size:  "kingsize",
			Price: 99.99,
		},
		{
			Size:  "medium",
			Price: 79.99,
		},
	}
	insertedHotel, err := hotelStore.InsertHotel(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}
	for _, room := range rooms {
		room.HotelId = insertedHotel.ID
		_, err := roomStore.InsertRoom(ctx, &room)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	Init()
	seedHotel("Bellucia", "France", 3)
	seedHotel("The Cozy Hotel", "Netherlands", 2)
	seedHotel("Ibis", "Brasil", 5)
    seedUser("Rick", "Renner", "rick@renner.com")
}

func Init() {
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Database(db.DBNAME).Drop(ctx); err != nil {
		log.Fatal(err)
	}

	hotelStore = db.NewMongoHotelStore(client)
	roomStore = db.NewMongoRoomStore(client, hotelStore)
	userStore = db.NewMongoUserStore(client)
}
