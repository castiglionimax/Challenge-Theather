package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	ctx    context.Context
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

	if err != nil {
		fmt.Print("aca salta")

		log.Fatal(err)
	}
	fmt.Print("hasta aca vengo bien")
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

}

func GetSession() (*mongo.Client, context.Context) {
	return client, ctx
}
