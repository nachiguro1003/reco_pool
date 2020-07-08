package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nachiguro1003/reco-pool/src/infrastructure/datastore/db/postgresql"
	reco_repository "github.com/nachiguro1003/reco-pool/src/infrastructure/datastore/repository/reco"
	reco_service "github.com/nachiguro1003/reco-pool/src/usecase/reco"
	"log"
	"net/http"
)

func RecoCreate(c *gin.Context) {
	d := postgresql.GetPostgresInstance()
	re := reco_repository.NewRecoRepository()

	if err := c.BindJSON(&re); err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError,err.Error())
		return
	}
	ps := reco_service.NewRecoService(re)
	err := ps.CreateReco(d.DB)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError,err.Error())
		return
	}
	c.JSON(http.StatusOK,"reco created")
}
