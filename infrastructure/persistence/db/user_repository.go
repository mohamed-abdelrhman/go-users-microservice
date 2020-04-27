package db

import (
	"context"
	go_errors "github.com/mohamed-abdelrhman/go-errors"
	go_logger "github.com/mohamed-abdelrhman/go-logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
	"macqueen_users/domain/entity"
	"macqueen_users/infrastructure/clients"
)

const (
	usersCollection="users"
)

type UserRepositoryInterface interface {
	SaveUser(*entity.User) (*entity.User, go_errors.RestErr)
	GetUser(string) (*entity.User, go_errors.RestErr)
	GetUsers() ([]entity.User, go_errors.RestErr)
	GetUserByEmailAndPassword(*entity.User) (*entity.User, go_errors.RestErr)
}

type userRepository struct {
}

func NewUserRepository() UserRepositoryInterface {
	return &userRepository{}
}

func (r *userRepository)SaveUser(user *entity.User) (*entity.User, go_errors.RestErr){
	user.Id=primitive.NewObjectID()
	clients.GetMongoClient().Collection(usersCollection)
	_, err := clients.GetMongoClient().Collection(usersCollection).InsertOne(context.TODO(),user)
	if err != nil {
		go_logger.Error("error inserting user",err)
		return nil, go_errors.NewInternalServerError("Database error")
	}
	return user,nil
}


func (r *userRepository)GetUsers() ([]entity.User, go_errors.RestErr){
	return nil,nil
}

func (r *userRepository)GetUser( docId string) (*entity.User, go_errors.RestErr){
	var user entity.User
	docID, objectErr := primitive.ObjectIDFromHex(docId)
	if objectErr != nil {
		go_logger.Error("user repository,Get user,Getting Object Id",objectErr)
		return nil,go_errors.NewInternalServerError("Error Getting User")
	}
	result:=clients.GetMongoClient().Collection("users").FindOne(context.TODO(),bson.M{"_id": docID})
	err:=result.Decode(&user)
	if err != nil {
		go_logger.Error("user repository,Get user , decoding user",err)
		return nil,go_errors.NewInternalServerError("Error Decoding user")
	}
	return &user,nil
}
func (r *userRepository)GetUserByEmailAndPassword(user *entity.User) (*entity.User, go_errors.RestErr){
	return nil,nil
}
