package repo

type UserRepo struct {}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}
//start with a lowercase letter are unexported
func (ur *UserRepo) GetUserInfo() string {
	return "Minh";
}