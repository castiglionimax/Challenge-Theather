package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	URL          = "mongodb+srv://dbAlexis:RENccHo22n8RuPDi@cluster0.rsjmz.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
	IndexBooking = "bookings"
)

var (
	client *mongo.Client
	//password = os.Getenv("MongodbMeliShow")
)

func init() {
	var err error

	client, err = mongo.NewClient(options.Client().ApplyURI(URL))

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

}

func GetSession() *mongo.Client {
	return client
}
