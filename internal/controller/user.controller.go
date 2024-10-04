package controller

import (
	"fmt"

	"github.com/dinhminh0307/Shop-Dev-Ecommerce/internal/service"
)

// UserController struct
type UserController struct {
	// call the service
	userService *service.UserService
}


// Default constructor
func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// Method to get the user by name
func (uc *UserController) GetUserByName() string {
	fmt.Println("called");
	// Get the 'name' query parameter from the URL (e.g., /ping/123?name=minh)
	name := uc.userService.GetUserInfoService()
  
	return name
}


