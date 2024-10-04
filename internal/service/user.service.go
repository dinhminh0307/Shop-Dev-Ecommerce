package service

import "github.com/dinhminh0307/Shop-Dev-Ecommerce/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo // assign attribute of UserRepo addres
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(), // assign new object to the attribute
	}
}

func (us *UserService) GetUserInfoService() string {
	return us.userRepo.GetUserInfo()
}