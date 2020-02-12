package router

import (
	"github.com/gin-gonic/gin"
	"github.com/reco_pool/src/interface/controllers"
	"github.com/reco_pool/src/interface/presenter"
)

func RegisterRoutingGroup(r *gin.Engine) {
	Pool(r)
}

func Pool(r *gin.Engine) {
	p := r.Group("/pool")
	p.GET("/:slug",presenter.Pool)
	p.POST("/new",controllers.NewPool)
}
