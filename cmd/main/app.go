package main

import (
	"bd-backend/internal/config"
	handlermanager "bd-backend/internal/handlers/manager"
	"bd-backend/pkg/client/postgresql"
	"bd-backend/pkg/logging"
	repeatable "bd-backend/pkg/utils"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {

	cfg := config.GetConfig()
	logger := logging.GetLogger()

	err := repeatable.CrateDir()
	if err != nil {
		logger.Fatalf("%v", err)
	}

	postgresSQLClient := startPostgresql(cfg, logger)

	start(handlermanager.Manager(postgresSQLClient, logger), cfg)
}

func startPostgresql(cfg *config.Config, logger *logging.Logger) *pgxpool.Pool {
	postgresSQLClient, err := postgresql.NewClient(context.TODO(), 3, cfg.Storage)
	if err != nil {
		return nil

	}
	return postgresSQLClient
}

func start(router *gin.Engine, cfg *config.Config) {
	logger := logging.GetLogger()
	logger.Info("start application")

	router.StaticFS("/public", gin.Dir(cfg.PublicFilePath, false))
	router.Run(cfg.Listen.BindIP + ":" + cfg.Listen.Port)
}
