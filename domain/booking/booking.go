package booking

import "github.com/castiglionimax/MeliShows-Challenge/domain/performance"

type Booking struct {
	Performace performance.Performance `json:"performace"`
	Person     Person                  `json:"person"`
	Sold       []Sold                  `json:"sold"`
}

type Sold struct {
	Seat    int                 `json:"seat"`
	Section performance.Section `json:"section"sectionId"`
}

type Person struct {
	Dni      int    `json:"dni"`
	Fullname string `json:"fullname"`
}
