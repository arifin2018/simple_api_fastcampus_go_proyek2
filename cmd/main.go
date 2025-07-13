package main

import (
	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/internal/handlers/memperships"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	membershipsHandler := memperships.NewHandler(r)
	membershipsHandler.RegisterRoute()

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":3000")
}
