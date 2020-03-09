package presenter

import (
	"github.com/gin-gonic/gin"
	"github.com/nachiguro1003/reco_pool/src/domain/model"
	"github.com/nachiguro1003/reco_pool/src/infrastructure/datastore/db"
	"github.com/nachiguro1003/reco_pool/src/infrastructure/datastore/repository/pool"
	"github.com/nachiguro1003/reco_pool/src/usecase/pool"
	"net/http"
)

func Pool(c *gin.Context) {
	d := db.GetDB()
	pe := pool_repository.NewPoolRepository()
	pe.Slug = c.Param("slug")
	ps := pool_service.NewPoolService(pe)

	_,err := ps.GetPool(d)
	if err != nil {
		c.JSON(http.StatusBadRequest,err.Error())
	}

	c.JSON(http.StatusOK,ps)
}

func Pools(c *gin.Context) {
	d := db.GetDB()
	pools := []*model.Pool{}
	d.Find(&pools)
	c.JSON(http.StatusOK,pools)
}
