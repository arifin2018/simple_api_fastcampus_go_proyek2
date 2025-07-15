package main

import (
	"fmt"
	"log"

	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/internal/configs"
	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/internal/handlers/memperships"
	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/internal/repositories/memberships"
	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/pkg/internalsql"
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

	err := configs.Init(
		configs.WithConfigFolder([]string{
			"./internal/configs",
		}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yml"),
	)

	if err != nil {
		log.Fatal("gagal inisiasi config ", err)
	}

	cfg := configs.Get()
	fmt.Println("cfg.Service")
	fmt.Println(cfg)

	dataDBsourcename := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.DB)
	db, err := internalsql.Connect(cfg.Database.Drivername, dataDBsourcename)
	if err != nil {
		log.Fatal("gagal inisiasi config ", err)
	}
	memberships.NewRepository(db)
	// r.Run(":3000")
	r.Run(cfg.Service.Port)
}
