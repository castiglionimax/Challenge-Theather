package mongodb

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	IndexBooking = "bookings"
)

var (
	client *mongo.Client
)

func init() {
	var err error

	errxa := godotenv.Load()
	if errxa != nil {
		log.Fatal("Error loading .env file")
	}

	URL := os.Getenv("URL_MONGODB")
	fmt.Println(URL)

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
