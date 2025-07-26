package main

import (
	"log"

	"proyek3-catalog-music/internal/configs"
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

	r := gin.Default()
	r.Run(cfg.Service.Port) // Start the server on the configured address
}
