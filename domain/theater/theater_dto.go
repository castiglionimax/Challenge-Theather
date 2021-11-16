package theater

type Theater struct {
	Name        string       `json:"name" bson:"name"`
	Auditoriums []Auditorium `json:"auditorium" bson:"auditorium"`
}

type Auditorium struct {
	Name  string `json:"name" bson:"name"`
	Seats int    `json:"seats" bson:"seats"`
}
