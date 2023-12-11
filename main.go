package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ricardoraposo/gohotel/db"
	"github.com/ricardoraposo/gohotel/handlers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUri))
	if err != nil { log.Fatal(err) }

	userHandler := handlers.NewUserHandler(db.NewMongoUserStore(client))

	app := fiber.New(config)
	api := app.Group("/api")

	api.Get("/user", userHandler.GetUsers)
	api.Get("/user/:id", userHandler.GetUser)
	api.Post("/user", userHandler.PostUser)
    api.Delete("/user/:id", userHandler.DeleteUser)
    api.Put("/user/:id", userHandler.UpdateUser)

	app.Listen(listenAddr)
}
