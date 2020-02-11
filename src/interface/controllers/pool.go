package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/reco_pool/src/domain/model"
	"github.com/reco_pool/src/infrastructure/datastore/db"
	"net/http"
)

func NewPool(c *gin.Context) {
	d := db.GetDB()
	ps := model.NewPoolService()
	err := ps.CreateNewPool(d)
	if err != nil {
		c.JSON(http.StatusInternalServerError,"failed to create pool")
	}
	c.JSON(http.StatusOK,"pool created")
}
