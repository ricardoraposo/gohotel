package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ricardoraposo/gohotel/data"
	"github.com/ricardoraposo/gohotel/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookRoomParams struct {
	FromDate   time.Time `json:"fromDate"`
	TillDate   time.Time `json:"tillDate"`
	NumPersons int       `json:"numPersons"`
}

func (p BookRoomParams) validate() error {
	now := time.Now()
	if now.After(p.FromDate) || now.After(p.TillDate) {
		return fmt.Errorf("Cannot book rooms in the past")
	}
	return nil
}

type RoomHandler struct {
	store db.Store
}

func NewRoomHandler(store db.Store) *RoomHandler {
	return &RoomHandler{store}
}

func (h *RoomHandler) BookRoom(c *fiber.Ctx) error {
	var params BookRoomParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	if err := params.validate(); err != nil {
		return err
	}

	roomID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return err
	}

	user, ok := c.Context().Value("user").(data.User)
	if !ok {
		return c.Status(http.StatusInternalServerError).JSON(data.StdError{Msg: "Internal server error"})
	}

    ok, err = h.isRoomAvailableForBooking(c.Context(), roomID, params)
    if err != nil {
        return err
    }
    if !ok {
        return c.Status(http.StatusBadRequest).JSON(data.StdError{
            Msg: fmt.Sprintf("Room %s already booked", roomID.String()),
        })
    }


	booking := data.Booking{
		UserID:     user.ID,
		RoomID:     roomID,
		FromDate:   params.FromDate,
		TillDate:   params.TillDate,
		NumPersons: params.NumPersons,
	}

	inserted, err := h.store.Booking.InsertBooking(c.Context(), &booking)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(inserted)
}

func (h *RoomHandler) isRoomAvailableForBooking(ctx context.Context, roomID primitive.ObjectID, params BookRoomParams) (bool, error) {
	where := bson.M{
		"roomID": roomID,
		"fromDate": bson.M{
			"$gte": params.FromDate,
		},
		"tillDate": bson.M{
			"$lte": params.TillDate,
		},
	}
	bookings, err := h.store.Booking.GetBookings(ctx, where)
	if err != nil {
		return false, err
	}
    ok := len(bookings) == 0 
    return ok, nil
}
