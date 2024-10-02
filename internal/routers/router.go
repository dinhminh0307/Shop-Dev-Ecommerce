package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func NewRouter() *gin.Engine {
	r := gin.Default()

	// Group API to endpoint
	v1 := r.Group("test/api") 
	{
	  v1.GET("/ping/:id", greetingMessage)  // Corrected to get 'id' from the URL path
	  v1.PUT("/ping", greetingMessage)
	  v1.POST("/ping", greetingMessage)
	  v1.DELETE("/ping", greetingMessage)
	}
	
	return r
	
}

// API or router handler
func greetingMessage(c *gin.Context) {
	// Get the 'name' query parameter from the URL (e.g., /ping/123?name=minh)
	name := c.DefaultQuery("name", "hello") // if dont find name querry
  
	// Get the 'id' parameter from the URL path (e.g., /ping/123)
	id := c.Param("id")
  
	// Return the response as JSON with a greeting message and id
	c.JSON(http.StatusOK, gin.H{
	  "message": "Hello " + name,
	  "id":      id,
	})
  }