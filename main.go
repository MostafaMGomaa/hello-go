package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var uri = "mongodb://localhost:27017/go_demo"
var mongoClient *mongo.Client

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world",
		})
	})

	r.Run()
}

func connectToMongoDB() error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err == nil {
		mongoClient = client
		return nil
	}
	return err
}
