package repository

import (
	"time"

	"github.com/jordanhw34/GoWebStarter/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(res models.RoomRestriction) error
	AvailableByRoomIDAndDates(startDate, endDate time.Time, roomID int) (bool, error)
	AvailableByDates(startDate, endDate time.Time) ([]models.Room, error)
	GetRoomByID(id int) (models.Room, error)
	GetUserByID(id int) (models.User, error)
	UpdateUser(u models.User) error
	Authenticate(email, testPassword string) (int, string, error)
	GetAllReservations() ([]models.Reservation, error)
}
