package presenter

import (
	"github.com/gin-gonic/gin"
	"github.com/nachiguro1003/reco-pool/src/domain/model"
	"github.com/nachiguro1003/reco-pool/src/infrastructure/datastore/db/postgresql"
	"net/http"
)

func Users(c *gin.Context) {
	d := postgresql.GetPostgresInstance()
	users := []*model.User{}
	d.DB.Find(&users)
	c.JSON(http.StatusOK,users)
}
