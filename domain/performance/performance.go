package performance

import (
	"log"
	"time"

	"github.com/castiglionimax/MeliShows-Challenge/domain/queries"
)

type Performance struct {
	DocumentID    string    `json:"_id", omitempty`
	PerformanceID int64     `json:"performanceID" bson:"performanceID"`
	ShowID        int64     `json:"showID" bson:"showID"`
	ShowName      string    `json:"showName" bson:"showName"`
	TheaterID     int64     `json:"theaterID" bson:"theaterID"`
	City          string    `json:"city" bson:"city"`
	TheaterName   string    `json:"theaterName" bson:"theaterName"`
	Auditorium    string    `json:"auditorium" bson:"auditorium"`
	Sections      []Section `json:"sections" bson:"sections"`
	DateTimeStamp *int64    `json:"date,omitempty" bson:"date"`
	Date          time.Time `json:"dateShow"`
}

type Section struct {
	SectionID   int64   `json:"id" bson:"id"`
	Name        string  `json:"name" bson:"name"`
	Description string  `json:"description" bson:"description"`
	Seats       []int   `json:"seats" bson:"seats"`
	Price       float64 `json:"price" bson:"price"`
	Currency    string  `json:"currency" bson:"currency"`
}

func (p Performance) ValidatePrice(query queries.EsQuery) {

	if query.Range_price != nil {
		//10 30 && 100 30

		for index, v := range p.Sections {
			log.Print(v.Price)
			//log.Print(float64(query.Range_price.From))
			//	log.Print(float64(query.Range_price.To))

			if (v.Price < float64(query.Range_price.From)) || (v.Price > float64(query.Range_price.To)) {
				p.Sections = append(p.Sections[:index], p.Sections[index+1:]...)
			}
		}
	}
}

func (p Performance) UpdateSeats(sectionId int64, numberSeat int) {

	for _, s := range p.Sections {
		if s.SectionID == sectionId {
			for index, seat := range s.Seats {
				if seat == numberSeat {
					s.Seats = append(s.Seats[:index], s.Seats[index+1:]...)
				}
			}

		}
	}

}
