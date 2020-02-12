package main

import (
	"github.com/gin-gonic/gin"
	"github.com/reco_pool/src/infrastructure/datastore/db"
	"github.com/reco_pool/src/infrastructure/router"
)

func main() {
	r := gin.Default()
	router.RegisterRoutingGroup(r)
	db.Init()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	db.Close()
}
