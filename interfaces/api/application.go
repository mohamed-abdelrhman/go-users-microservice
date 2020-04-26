package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"macqueen_users/infrastructure/clients"
)

var(
	Router=gin.Default()
)

func StartApplication() {
	clients.GetMongoClient()
	UserUrlMapping()
	log.Fatal(Router.Run(":8000"))
}
