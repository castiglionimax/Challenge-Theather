package performance

import "time"

type Performance struct {
	PerformanceID int64     `json:"performanceID" bson:"performanceID"`
	ShowID        int64     `json:"showID" bson:"showID"`
	ShowName      string    `json:"showName" bson:"showName"`
	TheaterID     int64     `json:"theaterID" bson:"theaterID"`
	City          string    `json:"city" `
	TheaterName   string    `json:"theaterName" bson:"theaterName"`
	Auditorium    string    `json:"auditorium" bson:"auditorium"`
	Sections      []Section `json:"sections" bson:"sections"`
	DateTimeStamp *int64    `json:"date,omitempty" bson:"sections"`
	Date          time.Time `json:"dateShow"`
}

type Section struct {
	SeactionID  int64   `json:"id" bson:"id"`
	Name        string  `json:"name" bson:"name"`
	Description string  `json:"description" bson:"description"`
	Seats       []int   `json:"seats" bson:"seats"`
	Price       float64 `json:"price" bson:"price"`
	Currency    string  `json:"currency" bson:"currency"`
}
