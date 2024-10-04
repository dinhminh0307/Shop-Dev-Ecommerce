package routers

import (
	"net/http"
	"github.com/dinhminh0307/Shop-Dev-Ecommerce/internal/controller"
	"github.com/gin-gonic/gin"
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
//flow:
// router -> controller -> service -> repo -> model -> db
// API or router handler
func greetingMessage(c *gin.Context) {
	// Get the 'name' query parameter from the URL (e.g., /ping/123?name=minh)
	name := controller.NewUserController().GetUserByName(); // call the method of 
  
	// Get the 'id' parameter from the URL path (e.g., /ping/123)
	id := c.Param("id")
	
	// Return the response as JSON with a greeting message and id
	c.JSON(http.StatusOK, gin.H{
	  "message": "Hello " + name,
	  "id":      id,
	})
  }