package db

const (
	DBNAME      = "hotel-reservation"
	DBNAME_TEST = "hotel-reservation-test"
	DBURI       = "mongodb://localhost:27017"
)

type Store struct {
	User  UserStore
	Hotel HotelStore
	Room  RoomStore
}
