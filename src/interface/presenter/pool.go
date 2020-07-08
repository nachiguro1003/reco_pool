package presenter

import (
	"github.com/gin-gonic/gin"
	"github.com/nachiguro1003/reco-pool/src/domain/model"
	"github.com/nachiguro1003/reco-pool/src/infrastructure/datastore/db/postgresql"
	"github.com/nachiguro1003/reco-pool/src/infrastructure/datastore/repository/pool"
	"github.com/nachiguro1003/reco-pool/src/usecase/pool"
	"net/http"
)

func Pool(c *gin.Context) {
	d := postgresql.GetPostgresInstance()
	pe := pool_repository.NewPoolRepository()
	pe.Slug = c.Param("slug")
	ps := pool_service.NewPoolService(pe)

	_,err := ps.GetPool(d.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest,err.Error())
	}

	c.JSON(http.StatusOK,ps)
}

func Pools(c *gin.Context) {
	d := postgresql.GetPostgresInstance()
	pools := []*model.Pool{}
	d.DB.Find(&pools)
	c.JSON(http.StatusOK,pools)
}
