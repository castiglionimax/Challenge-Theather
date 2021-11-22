package booking

import (
	"context"
	"encoding/json"
	"log"

	"github.com/castiglionimax/MeliShows-Challenge/database/mongodb"
	"github.com/castiglionimax/MeliShows-Challenge/domain/performance"
	"github.com/castiglionimax/MeliShows-Challenge/utils/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

const (
	DATABASE_NAME       = "MeliShows"
	SCHEMA_BOOKINGS     = "bookings"
	SCHEMA_PERFORMANCES = "performances"
)

func (v *Booking) SaveMongo() *errors.RestErr {

	client := mongodb.GetSession()

	wc := writeconcern.New(writeconcern.WMajority())
	rc := readconcern.Snapshot()
	txnOpts := options.Transaction().SetWriteConcern(wc).SetReadConcern(rc)

	session, err := client.StartSession()
	if err != nil {
		panic(err)
	}
	defer session.EndSession(context.Background())

	quickstartDatabase := client.Database(DATABASE_NAME)
	bookingsCollection := quickstartDatabase.Collection(SCHEMA_BOOKINGS)
	performancesCollection := quickstartDatabase.Collection(SCHEMA_PERFORMANCES)

	callback := func(sessionContext mongo.SessionContext) (interface{}, error) {

		_, err := bookingsCollection.InsertOne(
			sessionContext,
			v,
		)
		if err != nil {
			return nil, err
		}

		//	performance := &performance.Performance{}

		var performance performance.Performance
		var z bson.M

		err = performancesCollection.FindOne(sessionContext, bson.M{"performanceID": v.PerformanceID}).Decode(&z)
		if err != nil {
			return nil, err
		}

		arrayByte, _ := json.Marshal(z)

		err = json.Unmarshal(arrayByte, &performance)
		if err != nil {
			log.Print(err.Error())
			return nil, err
		}

		for _, s := range v.Sold {
			performance.UpdateSeats(s.SectionID, s.Seat)
		}

		asd, err := performancesCollection.ReplaceOne(
			sessionContext,
			bson.M{"performanceID": v.PerformanceID},
			performance,
		)

		if err != nil {
			return nil, err
		}
		log.Print(asd)

		return nil, nil
	}

	_, err = session.WithTransaction(context.Background(), callback, txnOpts)
	if err != nil {
		panic(err)
	}

	return nil

}
