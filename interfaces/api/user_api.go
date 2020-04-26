package api

import (
	"macqueen_users/application"
	"macqueen_users/domain/service"
	"macqueen_users/infrastructure/persistence/db"
	"macqueen_users/interfaces/http"
)

func UserUrlMapping()  {
	//user routes
	users := http.NewUsers(application.NewUserApp(service.NewUserService(db.NewUserRepository())))
	Router.POST("/users", users.SaveUser)
	Router.GET("/users", users.GetUsers)
	Router.GET("/users/:user_id", users.GetUser)

}
