package httphandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/santa512/monorepo-microservices/bookings/storage"
	mgo "gopkg.in/mgo.v2"
)

// Handler for HTTP Get - "/bookings"
// Get all Bookings
func GetBookings(s storage.Storage) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bookings := s.GetAll()
		// Create response data
		j, err := json.Marshal(BookingsResource{Data: bookings})
		if err != nil {
			panic(err)
		}
		// Send response back
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	})
}

// Handler for HTTP Post - "/bookings"
// Create a new Booking document
func CreateBooking(s storage.Storage) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var dataResource BookingResource
		// Decode the incoming Booking json
		err := json.NewDecoder(r.Body).Decode(&dataResource)
		if err != nil {
			panic(err)
		}
		booking := &dataResource.Data
		// Create Booking
		s.Create(booking)
		// Create response data
		j, err := json.Marshal(dataResource)
		if err != nil {
			panic(err)
		}
		// Send response back
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	})
}

// Handler for HTTP Delete - "/bookins"
// Delete a new Booking document
func DeleteBooking(s storage.Storage) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		// Remove showtime by id
		err := s.Delete(id)
		if err != nil {
			if err == mgo.ErrNotFound {
				w.WriteHeader(http.StatusNotFound)
				return
			} else {
				panic(err)
			}
		}

		// Send response back
		w.WriteHeader(http.StatusNoContent)
	})
}

func GetReadiness(s storage.Storage) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := s.Ping()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("error: %v", err)))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("ok"))
		}
	})
}

func GetLiveness() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
}
