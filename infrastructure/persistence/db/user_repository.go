package db

import (
	"context"
	go_errors "github.com/mohamed-abdelrhman/go-errors"
	go_logger "github.com/mohamed-abdelrhman/go-logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"macqueen_users/domain/entity"
	"macqueen_users/infrastructure/clients"
)

type UserRepositoryInterface interface {
	SaveUser(*entity.User) (*entity.User, go_errors.RestErr)
	GetUser(uint64) (*entity.User, go_errors.RestErr)
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
	_, err := clients.GetMongoClient().Collection("users").InsertOne(context.TODO(),user)
	if err != nil {
		go_logger.Error("error inserting user",err)
		return nil, go_errors.NewInternalServerError("Database error")
	}
	//Id:=result.InsertedID
	//user.Id=Id


	return user,nil
}


func (r *userRepository)GetUsers() ([]entity.User, go_errors.RestErr){
	return nil,nil
}
func (r *userRepository)GetUser(uint64) (*entity.User, go_errors.RestErr){
	//docID, _ := primitive.ObjectIDFromHex(user.Id)
	//
	//results:=clients.GetMongoClient().Collection("users").FindOne(context.TODO(),bson.M{"_id": docID})
	//
	//fmt.Println(results.DecodeBytes())

	return nil,nil
}
func (r *userRepository)GetUserByEmailAndPassword(user *entity.User) (*entity.User, go_errors.RestErr){
	return nil,nil
}
