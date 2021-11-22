package booking

type Booking struct {
	PerformanceID int64   `json:"performanceID" bson:"performaceID"`
	Person        Person  `json:"person" bson:"person"`
	Sold          []Sold  `json:"sold" bson:"sold"`
	TotalPrice    float32 `json:"total_price" bson:"total_price"`
}

type Sold struct {
	Seat      int     `json:"seat" bson:"seat"`
	SectionID int64   `json:"sectionID" bson:"sectionID"`
	Price     float32 `json:"price" bson:"price"`
}

type Person struct {
	Dni      int    `json:"dni" bson:"dni"`
	FullName string `json:"fullname" bson:"fullname"`
}

func (s *Sold) SetPrice(price float32) {
	s.Price = price
}
