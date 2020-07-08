package presenter

import (
	"github.com/gin-gonic/gin"
	"github.com/nachiguro1003/reco-pool/src/domain/model"
	"github.com/nachiguro1003/reco-pool/src/infrastructure/datastore/db/postgresql"
	"net/http"
)

func Recos(c *gin.Context) {
	d := postgresql.GetPostgresInstance()
	recos := []*model.Reco{}
	d.DB.Find(&recos)
	c.JSON(http.StatusOK,recos)
}
