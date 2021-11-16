package performance

import (
	"fmt"
	"log"

	"github.com/castiglionimax/MeliShows-Challenge/database/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

func (p *Performance) GetAllbyShowName() ([]*Performance, error) {
	fmt.Printf("aca estoy buen")
	client, ctx := mongodb.GetSession()

	meliDatabase := client.Database("MeliShows")
	performacesCollection := meliDatabase.Collection("performace")
	cursor, err := performacesCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var performaces []bson.M
	if err = cursor.All(ctx, &performaces); err != nil {
		log.Fatal(err)
	}
	fmt.Println(performaces)

	return nil, nil
}

//{createdAt:{$gte:ISODate("2021-01-01"),$lt:ISODate("2022-05-01"}}
