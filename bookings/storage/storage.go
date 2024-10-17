package storage

import "github.com/santa512/monorepo-microservices/bookings/model"

type Storage interface {
	GetAll() []model.Booking
	Create(booking *model.Booking) error
	Delete(id string) error
	Ping() error
}
