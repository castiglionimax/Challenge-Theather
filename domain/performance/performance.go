package performance

import (
	"log"

	"github.com/castiglionimax/MeliShows-Challenge/domain/queries"
)

type Performance struct {
	DocumentID    *string   `json:"_id,omitempty"`
	PerformanceID int64     `json:"performanceID" bson:"performanceID"`
	ShowID        int64     `json:"showID" bson:"showID"`
	ShowName      string    `json:"showName" bson:"showName"`
	TheaterID     int64     `json:"theaterID" bson:"theaterID"`
	TheaterName   string    `json:"theaterName" bson:"theaterName"`
	City          string    `json:"city" bson:"city"`
	Auditorium    string    `json:"auditorium" bson:"auditorium"`
	Sections      []Section `json:"sections" bson:"sections"`
	DateTimeStamp *int64    `json:"date,omitempty" bson:"date"`
}

type Section struct {
	SectionID   int64   `json:"id" bson:"id"`
	Name        string  `json:"name" bson:"name"`
	Description string  `json:"description" bson:"description"`
	Seats       []int   `json:"seats" bson:"seats"`
	Price       float32 `json:"price" bson:"price"`
}

func (p *Performance) ValidatePrice(query queries.EsQuery) {

	if query.Range_price != nil {
		//10 30 && 100 30

		for index, v := range p.Sections {
			log.Print(v.Price)
			//log.Print(float64(query.Range_price.From))
			//	log.Print(float64(query.Range_price.To))

			if (v.Price < float32(query.Range_price.From)) || (v.Price > float32(query.Range_price.To)) {
				p.Sections = append(p.Sections[:index], p.Sections[index+1:]...)
			}
		}
	}
}

func (p *Performance) UpdateSeats(sectionId int64, numberSeat int) float32 {

	for a := 0; a < len(p.Sections); a++ {
		if p.Sections[a].SectionID == sectionId {
			for index, seat := range p.Sections[a].Seats {
				if seat == numberSeat {
					priceSeat := p.Sections[a].Price
					//					s.Seats = append(s.Seats[:index], s.Seats[index+1:]...)
					p.Sections[a].Seats = removeSeat(p.Sections[a].Seats, index)
					return priceSeat
				}
			}

		}
	}
	return 0
}
func removeSeat(s []int, i int) []int {
	s[i] = s[len(s)-1]
	nwarry := make([]int, len(s)-1)
	for a := 0; a < len(nwarry); a++ {
		nwarry[a] = s[a]
	}
	return nwarry
}
func (p *Performance) ValidateAvailabilitySeat(sectionId int64, numberSeat int) bool {

	for _, s := range p.Sections {
		if s.SectionID == sectionId {
			for _, seat := range s.Seats {
				if seat == numberSeat {
					return false
				}
			}

		}
	}

	return true
}
