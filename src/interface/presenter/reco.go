package presenter

import (
	"github.com/gin-gonic/gin"
	"github.com/nachiguro1003/reco_pool/src/domain/model"
	"github.com/nachiguro1003/reco_pool/src/infrastructure/datastore/db"
	"net/http"
)

func Recos(c *gin.Context) {
	d := db.GetDB()
	recos := []*model.Reco{}
	d.Find(&recos)
	c.JSON(http.StatusOK,recos)
}
