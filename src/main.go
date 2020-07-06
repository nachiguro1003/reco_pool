package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nachiguro1003/reco_pool/src/infrastructure/datastore/db"
	"github.com/nachiguro1003/reco_pool/src/interface/router"
	"github.com/nachiguro1003/reco_pool/src/interface/router/api"
)

func run() error {
	r := gin.Default()
	router.RegisterRoutingGroup(r)
	api.RegisterAPIRoutingGroup(r)
	db.Init()
	if err := r.Run(); err != nil {
		return err
	}
	db.Close()
	return nil
}

func main() {
	if err := run();err != nil {
		panic(err)
	}
}
