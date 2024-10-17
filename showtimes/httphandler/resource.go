package httphandler

import "github.com/santa512/monorepo-microservices/showtimes/model"

type (
	// For Get - /showtimes
	ShowTimesResource struct {
		Data []model.ShowTime `json:"data"`
	}
	// For Post/Put - /showtimes
	ShowTimeResource struct {
		Data model.ShowTime `json:"data"`
	}
)
