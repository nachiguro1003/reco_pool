package api

import (
"github.com/gin-gonic/gin"
"github.com/nachiguro1003/reco_pool/src/interface/controllers"
"github.com/nachiguro1003/reco_pool/src/interface/presenter"
)

func RegisterAPIRoutingGroup(r *gin.Engine) {
	Index(r)
	Pool(r)
	User(r)
	Reco(r)
}

func Index(r *gin.Engine) {
	p := r.Group("/api/v1")
	p.GET("",presenter.Index)
}

func Pool(r *gin.Engine) {
	p := r.Group("/api/v1/pools")
	p.POST("/new",controllers.PoolCreate)
	//p.POST("/user/:slug",controllers.RegisterUser)
}

func User(r *gin.Engine) {
	u := r.Group("/api/v1/users")
	u.POST("/")
}

func Reco(r *gin.Engine) {
	re := r.Group("/api/v1/recos")
	re.POST("/new",controllers.RecoCreate)
}

func Auth(r *gin.Engine) {
	a := r.Group("/api/v1/auth")
	a.POST("/sign_up",controllers.AuthSignUp)
}
