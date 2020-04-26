package entity

import (
	go_errors "github.com/mohamed-abdelrhman/go-errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"macqueen_users/infrastructure/validations"
	"strings"
)

const (
	loginAction="login"
	updateAction="update"
	registerAction="register"
	deleteAction="delete"
)

type User struct{
	Id    primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	FirstName string `json:"first_name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
	Mobile string `json:"mobile" binding:"required"`
	Email string`json:"email" binding:"required"`
	BillingCountry string `json:"billing_country"`
	BillingCity string `json:"billing_city"`
	BillingState string `json:"billing_state"`
	BillingStreet string `json:"billing_street"`
	BillingPostalCode string `json:"billing_postal_code"`
	Password string `json:"password" binding:"required"`
	IsAdmin string `json:"is_admin"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
type Users []User


func (u *User)validate(action string)  go_errors.RestErr{
	var errCases []string
	switch strings.ToLower(action) {
	case "login":
		ok:=validations.ValidateEmail(u.Email)
		if !ok {
			errCases=append(errCases,"Invalid email Address")
		}
	case "register":
		ok:=validations.ValidateEmail(u.Email)
		if !ok {
			errCases=append(errCases,"Invalid email Address")
		}
	default:
	}
	if len(errCases)>0 {
		return go_errors.NewUnprocessableEntity("Please validate following details",errCases)
	}
	return nil
}



