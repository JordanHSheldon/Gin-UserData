package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	IsDisabled bool // hidden
	FirstName  string
	LastName   string
	Email      string
	UserName   string
	Password   string
	SettingsId int // hidden
	GameDataId int // hidden
}

func getUserData(id string) []User {
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
	return results
}

func main() {
	router := gin.Default()

	// get the user
	router.POST("/GetUserByName/:id", func(c *gin.Context) {
		id := c.Param("id")
		data := getUserData(id)
		c.JSON(200, data)
	})

	// Create a user
	router.POST("/RegisterUser/:test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": c.Param("test"),
		})
	})

	// update user
	router.POST("/update", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": "false",
		})
	})

	router.POST("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"result": "pong!"})
	})

	log.Fatal(router.Run(":8000"))
}
