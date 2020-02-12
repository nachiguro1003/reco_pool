package presenter

import (
	"github.com/gin-gonic/gin"
	"github.com/reco_pool/src/domain/model"
	"github.com/reco_pool/src/infrastructure/datastore/db"
	"github.com/reco_pool/src/usecase/pool"
	"net/http"
)

func Pool(c *gin.Context) {
	d := db.GetDB()
	ps := model.NewPoolService()
	ps.Slug = c.Param("slug")

	ps,err := pool_service.GetPool(ps,d)
	if err != nil {
		c.JSON(http.StatusBadRequest,err.Error())
	}

	c.JSON(http.StatusOK,ps)
}
