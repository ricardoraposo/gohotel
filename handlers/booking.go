package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ricardoraposo/gohotel/data"
	"github.com/ricardoraposo/gohotel/db"
	"go.mongodb.org/mongo-driver/bson"
)

type BookingHandler struct {
	store db.Store
}

func NewBookingHandler(store db.Store) *BookingHandler {
	return &BookingHandler{store}
}

// TODO: This has to be admin authorized only
func (h *BookingHandler) GetBookings(c *fiber.Ctx) error {
	bookings, err := h.store.Booking.GetBookings(c.Context(), bson.M{})
    if err != nil {
        return err
    }
	return c.JSON(bookings)
}

// TODO: this has to be user authorized
func (h *BookingHandler) GetBooking(c *fiber.Ctx) error {
    id := c.Params("id")
    booking, err := h.store.Booking.GetBookingByID(c.Context(), id)
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(data.StdError{Msg: "Shit doesn't exist"})
    }
	return c.JSON(booking)
}

func (h *BookingHandler) AdminThing(c *fiber.Ctx) error {
    return c.JSON("I said hey, what's going on")
}
