package main

import (
	"fmt"
	"log"

	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/internal/configs"
	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/internal/handlers/memperships"
	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/internal/handlers/posts"
	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/internal/repositories/memberships"

	postRepository "github.com/arifin2018/simple_api_fastcampus_go_proyek2/internal/repositories/posts"
	membershipService "github.com/arifin2018/simple_api_fastcampus_go_proyek2/internal/services/memberships"
	postService "github.com/arifin2018/simple_api_fastcampus_go_proyek2/internal/services/posts"
	"github.com/arifin2018/simple_api_fastcampus_go_proyek2/pkg/internalsql"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
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

	dataDBsourcename := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.DB)
	db, err := internalsql.Connect(cfg.Database.Drivername, dataDBsourcename)
	if err != nil {
		log.Fatal("gagal inisiasi config ", err)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	postRepo := postRepository.NewRepository(db)
	postService := postService.NewService(cfg, postRepo)
	postHandler := posts.NewHandler(r, postService)
	postHandler.RegisterRoute()

	membershipRepository := memberships.NewRepository(db)
	membershipService := membershipService.NewService(cfg, membershipRepository)
	membershipsHandler := memperships.NewHandler(r, membershipService)
	membershipsHandler.RegisterRoute()

	// r.Run(":3000")
	r.Run(cfg.Service.Port)
}
