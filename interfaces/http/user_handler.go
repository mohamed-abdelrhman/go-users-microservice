package http

import (
	"github.com/gin-gonic/gin"
	go_errors "github.com/mohamed-abdelrhman/go-errors"
	"macqueen_users/application"
	"macqueen_users/domain/entity"
	"net/http"
	"strconv"
)

type UserHandlerInterface interface {
	SaveUser(c *gin.Context)
	GetUsers(c *gin.Context)
	GetUser(c *gin.Context)
}

type userHandler struct {
	ua application.UserAppInterface
}

//Users constructor
func NewUsers(ua application.UserAppInterface ) UserHandlerInterface {
	return &userHandler{
		ua: ua,
	}

}

func (uh *userHandler) SaveUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr:=go_errors.NewBadRequestError(err.Error())
		c.JSON(restErr.Status(),restErr)
		return
	}
	newUser, err := uh.ua.SaveUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, newUser)
}

func (uh *userHandler) GetUsers(c *gin.Context) {
	users := entity.Users{} //customize user
	var err error
	users, err = uh.ua.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, users)
}

func (uh *userHandler) GetUser(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	user, err := uh.ua.GetUser(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}