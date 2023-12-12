package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ricardoraposo/gohotel/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HotelHandler struct {
	store db.Store
}

func NewHotelHandler(store db.Store) *HotelHandler {
	return &HotelHandler{store}
}

func (h *HotelHandler) GetHotels(c *fiber.Ctx) error {
	hotels, err := h.store.Hotel.GetHotels(c.Context(), nil)
	if err != nil {
		return err
	}
	return c.JSON(hotels)
}

func (h *HotelHandler) GetHotel(c *fiber.Ctx) error {
	id := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
    hotel, err := h.store.Hotel.GetHotel(c.Context(), oid)
    if err != nil {
        return err
    }
    return c.JSON(hotel)
}

func (h *HotelHandler) GetRooms(c *fiber.Ctx) error {
	id := c.Params("id")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"hotelID": oid}
	rooms, err := h.store.Room.GetRooms(c.Context(), filter)
	if err != nil {
		return err
	}
	return c.JSON(rooms)
}
