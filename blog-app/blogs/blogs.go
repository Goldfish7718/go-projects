package blogs

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Blog struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func getBlogs(c *gin.Context, collection *mongo.Collection) {
	var blogs []Blog

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error finding documents",
		})

		return
	}

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var blog Blog
		if err := cursor.Decode(&blog); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding document"})
			return
		}
		blogs = append(blogs, blog)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cursor error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"blogs": blogs})
}

func createBlog(c *gin.Context, collection *mongo.Collection) {
	var newBlog Blog

	if err := c.ShouldBindBodyWithJSON(&newBlog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	doc := bson.M{
		"title":    newBlog.Title,
		"content":  newBlog.Content,
		"createAt": time.Now(),
	}

	result, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error inserting document",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Blog created",
		"id":      result.InsertedID,
	})
}
