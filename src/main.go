package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nachiguro1003/reco-pool/src/infrastructure/datastore/db"
	"github.com/nachiguro1003/reco-pool/src/interface/router"
	"github.com/nachiguro1003/reco-pool/src/interface/router/api"
)

var (
	RecoPoolDBManager *db.RecoPoolDBManager
)

func run() error {
	r := gin.Default()
	router.RegisterRoutingGroup(r)
	api.RegisterAPIRoutingGroup(r)

	var err error
	RecoPoolDBManager,err = db.Init()
	if err != nil {
		return err
	}
	if err := r.Run(); err != nil {
		return err
	}
	RecoPoolDBManager.Close()
	return nil
}

func main() {
	if err := run();err != nil {
		panic(err)
	}
}
