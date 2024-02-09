package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var uri = "mongodb+srv://mostafa:PawGp90emr3NYAra@gocluster.ch0mfy8.mongodb.net/?retryWrites=true&w=majority"
var mongoClient *mongo.Client

func init() {
	if err := connect_to_mongodb(); err != nil {
		log.Fatal("Could not connect to MongoDB")
	}
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world",
		})
	})

	r.GET("/movies/:id", getMovie)
	r.GET("/movies", getMovies)
	r.POST("/movies", addMovie)
	r.Run()

	// Run the Gin server
	if err := r.Run(); err != nil {
		panic(err)
	}
}

func connect_to_mongodb() error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err == nil {
		mongoClient = client
		return nil
	}
	return err
}

func getMovie(c *gin.Context) {
	var movie bson.M
	id := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	err = mongoClient.Database("sample_mflix").Collection("movies").FindOne(context.TODO(), bson.D{{"_id", objectId}}).Decode(&movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		c.AbortWithStatus(404)
		return
	}
	c.JSON(http.StatusOK, movie)
}

func getMovies(c *gin.Context) {
	cursor, err := mongoClient.Database("sample_mflix").Collection("movies").Find(context.TODO(), bson.D{{}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Map results
	var movies []bson.M
	if err = cursor.All(context.TODO(), &movies); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return movies
	c.JSON(http.StatusOK, movies)
}

func addMovie(c *gin.Context) {
	// var movie Movie

	// c.Status(http.StatusOK)
}
