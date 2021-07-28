package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

func main(){
	//connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil{
		log.Fatalf("Cannot connect to mongo server", err)
	} else {
		err := client.Ping(ctx, readpref.Primary())
		if err != nil{
			log.Fatalf("Cannot ping server", err)
		} else {
			log.Print("Mongo server connected")
		}
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()


	//try inserting
	collection := client.Database("testing").Collection("numbers")
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, bson.D{{"name", "cos60"}, {"value", 0.5}})
	if err != nil {
		log.Fatal("error in inserting")
	} else {
		id := res.InsertedID
		log.Print("Inserted with Id: ", id)
	}

	//find
	ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil { log.Fatal(err) }
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil { log.Fatal(err) }
		log.Println("R:", result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	//find one
	var result struct {
		Value float64
	}
	filter := bson.D{{"name", "pi"}}
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = collection.FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		fmt.Println("record does not exist")
	} else if err != nil {
		log.Fatal(err)
	}
	log.Println("Found: ", result)
}