package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nachiguro1003/reco-pool/src/infrastructure/datastore/db/postgresql"
	"github.com/nachiguro1003/reco-pool/src/infrastructure/datastore/repository/user"
	"github.com/nachiguro1003/reco-pool/src/usecase/user"
	"log"
	"net/http"
)

func UserCreate(c *gin.Context) {
	d := postgresql.GetPostgresInstance()
	user := user_repository.NewUserRepository()

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError,err.Error())
		return
	}

	log.Print(user)

	us := user_service.NewUserService(user)
	if err := us.CreatePool(d.DB); err != nil {
		c.JSON(http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,"user created")
}
