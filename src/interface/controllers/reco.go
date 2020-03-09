package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nachiguro1003/reco_pool/src/infrastructure/datastore/db"
	reco_repository "github.com/nachiguro1003/reco_pool/src/infrastructure/datastore/repository/reco"
	reco_service "github.com/nachiguro1003/reco_pool/src/usecase/reco"
	"net/http"
)

func RecoCreate(c *gin.Context) {
	d := db.GetDB()
	re := reco_repository.NewRecoRepository()

	if err := c.BindJSON(&re); err != nil {
		c.JSON(http.StatusInternalServerError,err.Error())
		return
	}
	ps := reco_service.NewRecoService(re)
	err := ps.CreateReco(d)
	if err != nil {
		c.JSON(http.StatusInternalServerError,err.Error())
		return
	}
	c.JSON(http.StatusOK,"reco created")
}
