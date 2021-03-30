package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"test.com/mutant-checker/infrastructure/repositories_impl"
)

var this struct {
	db *mongo.Database
}

var connect = connectMongo

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func connectMongo() {
	var err error

	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("DB_URL")).SetMaxPoolSize(10))
	check(err)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	check(err)

	err = client.Ping(ctx, readpref.Primary())
	check(err)

	this.db = client.Database("mutant")
	repositories_impl.Setup(this.db)
}

// Initialize is a manual init func to connect to the db
func Initialize() {
	connect()
	fmt.Println("We are connected!")
}
