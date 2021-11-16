package performance

import (
	"time"

	"github.com/castiglionimax/MeliShows-Challenge/domain/show"
	"github.com/castiglionimax/MeliShows-Challenge/domain/theater"
)

type Performance struct {
	Show       show.Show          `json:"show" `
	Auditorium theater.Auditorium `json:"auditorium" bson:"auditorium"`
	Sections   []Section          `json:"sections"`
	Date       time.Time          `json:"date" bson:"date"`
}

type Section struct {
	ID          int    `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Seats       int    `json:"seats" bson:"seats"`
	Price       int    `json:"price" bson:"price"`
	Currency    string `json:"currency" bson:"currency"`
}
