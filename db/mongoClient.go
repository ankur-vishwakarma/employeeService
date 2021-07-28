package db
// change package to main for testing main func(also uncomment it)

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

func Connect(uri string)(context.Context, context.CancelFunc, *mongo.Client, error){
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Cannot Connect to db: ", err)
	}
	log.Println("Connnected to server!")
	return ctx, cancel, client, err
}

func Close(cancel context.CancelFunc, client *mongo.Client, ctx context.Context)(){
	defer cancel()
	defer func(){
		if err := client.Disconnect(ctx); err != nil{
			panic(err)
		}
	}()
	log.Println("connection closed")
}

func CheckPing(client *mongo.Client, ctx context.Context)(){
	err := client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Pinged successfully")
}

func GetDatabases(ctx context.Context, client *mongo.Client)([]string, error){
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal("error while getting db: ", err)
	}
	return databases, err
}

//can take document here as an interface eg: document interface{}
func InsertOne(client *mongo.Client, ctx context.Context, databaseName string, collectionName string, document bson.D)(*mongo.InsertOneResult, error){
	db := client.Database(databaseName)
	collection := db.Collection(collectionName)
	result, err := collection.InsertOne(ctx, document)
	return result, err
}

func GetOne(client *mongo.Client, ctx context.Context, databaseName string, collectionName string, filter bson.M)([]bson.M){
	db := client.Database(databaseName)
	collection := db.Collection(collectionName)
	filterCursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	var episodesFiltered []bson.M
	err = filterCursor.All(ctx, &episodesFiltered)
	if err != nil{
		log.Fatal("Error in getting: ", err)
	}
	return episodesFiltered
}
/*
func main() {
	//Connect
	ctx, cancel, client, err := Connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	//Close
	defer Close(cancel, client, ctx)

	//try pinging
	CheckPing(client, ctx)

	//try checking db list
	databases, err := GetDatabases(ctx, client)
	if err != nil {
		panic(err)
	}

	log.Println("database list: ", databases)

	//try inserting
	document := bson.D{
		{"name", "sin90"},
		{"value" , 1},
	}
	result, err := InsertOne(client, ctx, "testing", "numbers", document)

	if err != nil {
		log.Fatal("error while inserting: ", err)
		panic(err)
	}

	log.Println("Inserted Successfully: ", result)

	//Query
	queryResult := GetOne(client, ctx, "testing", "numbers", bson.M{"name": "sin90"})
	log.Println("Get Result is: ", queryResult)
}*/