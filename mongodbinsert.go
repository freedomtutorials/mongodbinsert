package main

import (
    "context"
    "fmt"
    "log"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type Device struct 
{
	ID int
	Name string
}

func main() {
	
	// set the mongodb uri 
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	
	//connect to mongodb
    client, err := mongo.Connect(context.TODO(), clientOptions)

    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(context.TODO(), nil)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to MongoDB!")
	
	freedomtutorialsDB := client.Database("freedomtutorials")
	mydataCollection := freedomtutorialsDB.Collection("mydata")
	
	var device Device
	device.ID = 1
	device.Name = "Name"
	
	collInsertResult, err := mydataCollection.InsertOne(context.TODO(), device)
	
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted Single Document: ", collInsertResult.InsertedID)

}





