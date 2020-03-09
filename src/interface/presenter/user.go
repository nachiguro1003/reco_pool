package presenter

import (
	"github.com/gin-gonic/gin"
	"github.com/nachiguro1003/reco_pool/src/domain/model"
	"github.com/nachiguro1003/reco_pool/src/infrastructure/datastore/db"
	"net/http"
)

func Users(c *gin.Context) {
	d := db.GetDB()
	users := []*model.User{}
	d.Find(&users)
	c.JSON(http.StatusOK,users)
}
