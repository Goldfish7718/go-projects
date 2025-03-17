package blogs

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func RegisterRoutes(router *gin.Engine, client *mongo.Client) {
	collection := client.Database("blog_db").Collection("blogs")

	// Assign the handler with the collection
	router.GET("/blogs", func(c *gin.Context) {
		getBlogs(c, collection)
	})

	router.POST("/blogs", func(c *gin.Context) {
		createBlog(c, collection)
	})
}
