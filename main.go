package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ricardoraposo/gohotel/db"
	"github.com/ricardoraposo/gohotel/handlers"
	"github.com/ricardoraposo/gohotel/middleware"
)

func main() {
	client := db.NewMongoClient()

	userStore := db.NewMongoUserStore(client)
	hotelStore := db.NewMongoHotelStore(client)
	roomStore := db.NewMongoRoomStore(client, hotelStore)
	store := db.Store{
		Hotel: hotelStore,
		Room:  roomStore,
		User:  userStore,
	}

	userHandler := handlers.NewUserHandler(userStore)
	authHandler := handlers.NewAuthHandler(userStore)
	hotelHandler := handlers.NewHotelHandler(store)

	app := fiber.New(config)
	auth := app.Group("/auth")
	api := app.Group("/api", middleware.JWTAuth)

	// auth
	auth.Post("/", authHandler.Authenticate)

	// userRoutes
	api.Get("/user", userHandler.GetUsers)
	api.Get("/user/:id", userHandler.GetUser)
	api.Post("/user", userHandler.PostUser)
	api.Delete("/user/:id", userHandler.DeleteUser)
	api.Put("/user/:id", userHandler.UpdateUser)

	// hotelRoutes
	api.Get("/hotel", hotelHandler.GetHotels)
	api.Get("/hotel/:id", hotelHandler.GetHotel)
	api.Get("/hotel/:id/rooms", hotelHandler.GetRooms)

	app.Listen(listenAddr)
}
