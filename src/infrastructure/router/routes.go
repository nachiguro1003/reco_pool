package router

import (
	"github.com/gin-gonic/gin"
	"github.com/reco_pool/src/interface/controllers"
)

func RegisterRoutingGroup(r *gin.Engine) {
	Pool(r)
}

func Pool(r *gin.Engine) {
	p := r.Group("/pool")
	p.GET("/:slug",)
	p.POST("/new",controllers.NewPool)
}
