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
	bookingStore := db.NewMongoBookingStore(client)
	store := db.Store{
		Hotel:   hotelStore,
		Room:    roomStore,
		User:    userStore,
		Booking: bookingStore,
	}

	userHandler := handlers.NewUserHandler(userStore)
	authHandler := handlers.NewAuthHandler(userStore)
	hotelHandler := handlers.NewHotelHandler(store)
	roomHandler := handlers.NewRoomHandler(store)
	bookingHandler := handlers.NewBookingHandler(store)

	app := fiber.New(config)
	auth := app.Group("/auth")
	api := app.Group("/api", middleware.JWTAuth(userStore))
	admin := api.Group("/admin")

	// auth
	auth.Post("/", authHandler.Authenticate)

    // admin
    admin.Get("/booking", middleware.AdminAuth, bookingHandler.AdminThing)

	// user handlers
	api.Get("/user", userHandler.GetUsers)
	api.Get("/user/:id", userHandler.GetUser)
	api.Post("/user", userHandler.PostUser)
	api.Delete("/user/:id", userHandler.DeleteUser)
	api.Put("/user/:id", userHandler.UpdateUser)

	// hotel handlers
	api.Get("/hotel", hotelHandler.GetHotels)
	api.Get("/hotel/:id", hotelHandler.GetHotel)
	api.Get("/hotel/:id/rooms", hotelHandler.GetRooms)

    // room handlers
	api.Post("/room/:id/book", roomHandler.BookRoom)

    // booking handlers
    api.Get("/booking/:id", bookingHandler.GetBooking)
    //TODO: Cancel a booking

	app.Listen(listenAddr)
}
