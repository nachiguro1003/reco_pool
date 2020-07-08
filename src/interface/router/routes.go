package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nachiguro1003/reco-pool/src/interface/controllers"
	"github.com/nachiguro1003/reco-pool/src/interface/presenter"
)

func RegisterRoutingGroup(r *gin.Engine) {
	Index(r)
	Pool(r)
	User(r)
	Reco(r)
}

func Index(r *gin.Engine) {
	p := r.Group("/")
	p.GET("",presenter.Index)
}

func Pool(r *gin.Engine) {
	p := r.Group("/pools")
	p.GET("/:slug",presenter.Pool)
	p.GET("/",presenter.Pools)
	p.POST("/new",controllers.PoolCreate)
	//p.POST("/user/:slug",controllers.RegisterUser)
}

func User(r *gin.Engine) {
	u := r.Group("/users")
	u.POST("/new",controllers.UserCreate)
	u.GET("/",presenter.Users)
}

func Reco(r *gin.Engine) {
	re := r.Group("/recos")
	re.POST("/new",controllers.RecoCreate)
	re.GET("/",presenter.Recos)
}
