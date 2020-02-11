package presenter

import (
	"github.com/gin-gonic/gin"
	"github.com/reco_pool/src/domain/model"
	"github.com/reco_pool/src/infrastructure/datastore/db"
	"github.com/reco_pool/src/usecase/pool"
	"net/http"
)

func GetPools(c *gin.Context) {
	d := db.GetDB()
	ps := model.NewPoolService()
	if err := c.ShouldBind(&ps); err != nil {
		c.JSON(http.StatusBadRequest,"err")
	}

}
