package main

import (
	"blog-app/blogs"
	"context"
	"fmt"

	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	// Load MongoDB URI from environment or fallback to localhost
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	// Connect to MongoDB
	client, err := mongo.Connect(options.Client().ApplyURI(mongoURI))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Create a Gin router
	router := gin.Default()

	// Pass MongoDB client to handlers
	blogs.RegisterRoutes(router, client)

	// Start the server
	fmt.Println("Server running on :8080")
	router.Run(":3000")
}
