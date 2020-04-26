package application

import (
	"github.com/mohamed-abdelrhman/go-errors"
	"macqueen_users/domain/entity"
	"macqueen_users/domain/service"
)

type userApp struct {
	us service.UserServiceInterface
	//uR uR.RestRepository  // to implement rest from the application
}
//UserApp implements the UserAppInterface
var _ UserAppInterface = &userApp{}


type UserAppInterface interface {
SaveUser(*entity.User) (*entity.User, go_errors.RestErr)
GetUsers() ([]entity.User, go_errors.RestErr)
GetUser(uint64) (*entity.User, go_errors.RestErr)
GetUserByEmailAndPassword(*entity.User) (*entity.User, go_errors.RestErr)
}
//type UserAppInterface repository.UserRepositoryInterface

//Users constructor
func NewUserApp(us service.UserServiceInterface ) UserAppInterface {
	return &userApp{
		us: us,
	}
}

func (u *userApp) SaveUser(user *entity.User) (*entity.User, go_errors.RestErr) {
	return u.us.SaveUser(user)
}

func (u *userApp) GetUser(userId uint64) (*entity.User, go_errors.RestErr) {
	return u.us.GetUser(userId)
}

func (u *userApp) GetUsers() ([]entity.User, go_errors.RestErr) {
	return u.us.GetUsers()
}

func (u *userApp) GetUserByEmailAndPassword(user *entity.User) (*entity.User, go_errors.RestErr) {
	return u.us.GetUserByEmailAndPassword(user)
}