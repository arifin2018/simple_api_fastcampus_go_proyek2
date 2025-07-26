package main

import (
	"log"

	"proyek3-catalog-music/internal/configs"
	membershipHandler "proyek3-catalog-music/internal/handler/memberships"
	"proyek3-catalog-music/internal/models/memberships"
	membershipRepo "proyek3-catalog-music/internal/repository/memberships"
	membershipService "proyek3-catalog-music/internal/services/memberships"
	"proyek3-catalog-music/pkg/internalsql"

	"github.com/gin-gonic/gin"
)

func main() {
	var cfg configs.Config

	err := configs.Init(
		configs.WithConfigFolder([]string{
			"./configs/",
			"./internal/configs/",
		}),
		configs.WithConfigFile("config.yaml"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatalf("Failed to initialize config: %v", err)
	}

	cfg = *configs.Get()
	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&memberships.User{}); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	r := gin.Default()

	memberRepository := membershipRepo.NewRepository(db)
	membershipService := membershipService.NewService(&cfg, memberRepository)
	membershipHandler := membershipHandler.NewHandler(r, membershipService)
	membershipHandler.RegisterRoutes()

	r.Run(cfg.Service.Port) // Start the server on the configured address
}
