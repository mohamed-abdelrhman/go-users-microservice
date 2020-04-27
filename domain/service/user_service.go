package service

import (
	go_errors "github.com/mohamed-abdelrhman/go-errors"
	go_logger "github.com/mohamed-abdelrhman/go-logger"
	"macqueen_users/domain/entity"
	"macqueen_users/infrastructure/persistence/db"
	"macqueen_users/infrastructure/security"
)

type UserServiceInterface interface {
	SaveUser(*entity.User) (*entity.User, go_errors.RestErr)
	GetUser(string) (*entity.User, go_errors.RestErr)
	GetUsers() ([]entity.User, go_errors.RestErr)
	GetUserByEmailAndPassword(*entity.User) (*entity.User, go_errors.RestErr)
}


type userService struct {
	ur db.UserRepositoryInterface
}

//Users constructor
func NewUserService(ur db.UserRepositoryInterface ) UserServiceInterface {
	return &userService{
		ur: ur,
	}

}

func (s *userService)SaveUser(user *entity.User) (*entity.User, go_errors.RestErr){
	hashedPassword,hashErr:=security.Hash(user.Password)
	if hashErr != nil {
		go_logger.Error("error hashing password",hashErr)
	}
	user.Password=string(hashedPassword)

	return s.ur.SaveUser(user)
}
func (s *userService)GetUser( userId string) (*entity.User, go_errors.RestErr){
	return s.ur.GetUser(userId)
}
func (s *userService)GetUsers() ([]entity.User, go_errors.RestErr){
	return s.ur.GetUsers()
}
func (s *userService)GetUserByEmailAndPassword(user *entity.User) (*entity.User, go_errors.RestErr){
	return s.ur.GetUserByEmailAndPassword(user)
}