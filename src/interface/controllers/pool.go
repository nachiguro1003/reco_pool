package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nachiguro1003/reco_pool/src/infrastructure/datastore/db"
	"github.com/nachiguro1003/reco_pool/src/infrastructure/datastore/repository/pool"
	"github.com/nachiguro1003/reco_pool/src/usecase/pool"
	"net/http"
)

func PoolCreate(c *gin.Context) {
	d := db.GetDB()
	pe := pool_repository.NewPoolRepository()

	if err := c.BindJSON(&pe); err != nil {
		c.JSON(http.StatusInternalServerError,err.Error())
		return
	}
	ps := pool_service.NewPoolService(pe)
	err := ps.CreatePool(d)
	if err != nil {
		c.JSON(http.StatusInternalServerError,err.Error())
		return
	}
	c.JSON(http.StatusOK,"pool created")
}
//
//func RegisterUser(c *gin.Context) {
//	d := db.GetDB()
//	slug := c.Param("slug")
//	user := user_service.NewUserService()
//	if err := c.BindJSON(&user); err != nil {
//		c.JSON(http.StatusInternalServerError,err.Error())
//	}
//
//	if err :.RegisterUserIntoPool(slug,user,d); err != nil {
//		c.JSON(http.StatusInternalServerError,err.Error())
//	}
//	c.JSON(http.StatusOK,"pool created")
//}
