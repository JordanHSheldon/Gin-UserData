package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	IsDisabled bool
	FirstName  string
	LastName   string
	Email      string
	UserName   string
	Password   string
	SettingsId int
}

func main() {
	r := gin.Default()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the MongoDB server to verify that the connection is active
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Get a handle to the "mydb" database and the "user" collection
	database := client.Database("mydb")
	collection := database.Collection("user")

	// Define a filter (empty in this case to get all documents)
	filter := bson.M{}

	// Define a context with a timeout (adjust as needed)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find documents in the collection
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	var results []User
	for cursor.Next(ctx) {
		var person User
		if err := cursor.Decode(&person); err != nil {
			log.Fatal(err)
		}
		results = append(results, person)
	}

	// Print the retrieved data
	fmt.Println("Retrieved data:")
	for _, result := range results {
		fmt.Printf("Name: %s, Email: %s\n", result.FirstName, result.Email)
	}

	// Disconnect from the MongoDB server
	err = client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// get the user, if they do not exist create them.
	r.POST("/getOrCreate", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"IsDisabled": "false",
			"username":   "nadroj",
			"password":   "password",
			"email":      "nadroj@gmail.com",
			"settingsid": "1",
		})
	})

	// update user
	r.POST("/update", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"IsDisabled": "false",
			"username":   "nadroj",
			"password":   "password",
			"email":      "nadroj@gmail.com",
			"settingsid": "1",
		})
	})
	r.Run()
}
