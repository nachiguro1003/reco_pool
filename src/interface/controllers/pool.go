package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/reco_pool/src/domain/model"
	"github.com/reco_pool/src/infrastructure/datastore/db"
	"github.com/reco_pool/src/usecase/pool"
	"net/http"
)

func NewPool(c *gin.Context) {
	d := db.GetDB()
	ps := model.NewPoolService()
	if err := c.BindJSON(&ps); err != nil {
		c.JSON(http.StatusInternalServerError,err.Error())
		return
	}
	err := pool_service.Create(ps,d)
	if err != nil {
		c.JSON(http.StatusInternalServerError,err.Error())
		return
	}
	c.JSON(http.StatusOK,"pool created")
}
