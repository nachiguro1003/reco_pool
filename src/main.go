package main

import (
	"github.com/gin-gonic/gin"
	"github.com/reco_pool/src/infrastructure/datastore/db"
	"github.com/reco_pool/src/interface/router"
	"github.com/reco_pool/src/interface/router/api"
)

func main() {
	r := gin.Default()
	router.RegisterRoutingGroup(r)
	api.RegisterAPIRoutingGroup(r)
	db.Init()
	if err := r.Run(); err != nil {
		panic(err)
	}
	db.Close()
}
