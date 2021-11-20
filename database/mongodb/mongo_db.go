package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	ULR = "mongodb+srv://dbAlexis:RENccHo22n8RuPDi@cluster0.rsjmz.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
)

var (
	client *mongo.Client
	//password = os.Getenv("MongodbMeliShow")
)

func init() {
	//	var b bytes.Buffer
	var err error

	//	b.WriteString("mongodb+srv://dbAlexis:")
	//	b.WriteString("RENccHo22n8RuPDi")
	//	b.WriteString("@cluster0.rsjmz.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

	//client, err = mongo.NewClient(options.Client().ApplyURI(b.String()))
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb+srv://dbAlexis:RENccHo22n8RuPDi@cluster0.rsjmz.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

}

func GetSession() *mongo.Client {
	return client
}
