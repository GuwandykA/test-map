package handlermanager

import (
	"bd-backend/internal/admin/category"
	categorydb "bd-backend/internal/admin/category/db"
	"bd-backend/pkg/logging"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	testURL = "/api/timetables"
)

func Manager(client *pgxpool.Pool, logger *logging.Logger) *gin.Engine {

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	},
	))
	categoryRouterManager := router.Group(testURL)
	categoryRepository := categorydb.NewRepository(client, logger)
	categoryRouterHandler := category.NewHandler(categoryRepository, logger)
	categoryRouterHandler.Register(categoryRouterManager)

	return router
}
