package booking

type Booking struct {
	PerformanceID int64  `json:"performaceID" bson:"performaceID"`
	Person        Person `json:"person" bson:"person"`
	Sold          []Sold `json:"sold" bson:"sold"`
}

type Sold struct {
	Seat      int   `json:"seat" bson:"seat"`
	SectionID int64 `json:"sectionID" bson:"sectionID"`
}

type Person struct {
	Dni      int    `json:"dni" bson:"dni"`
	FullName string `json:"fullname" bson:"fullname"`
}
