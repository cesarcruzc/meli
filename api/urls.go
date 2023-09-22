package api

import (
	v1 "github.com/cesarcruzc/meli/api/v1"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

// LoadUrls carga las rutas de la aplicación
func LoadUrls(database *mongo.Database, logger *log.Logger) *gin.Engine {
	url := gin.New()

	url.GET("/", v1.Root)
	url.GET("/health", v1.HealthCheck)
	apiV1 := url.Group("/api/v1")
	{
		apiV1.Use(v1.DBMiddleware(database), v1.LogMiddleware(logger))
		apiV1.POST("/process-file", v1.StartProcess)
		apiV1.GET("/items", v1.GetItems)

		apiV1.GET("/token", v1.GetApiToken)
		apiV1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return url
}
