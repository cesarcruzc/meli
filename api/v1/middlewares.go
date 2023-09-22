package v1

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func DBMiddleware(database *mongo.Database) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("database", database)
		ctx.Next()
	}
}

func LogMiddleware(logger *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("logger", logger)
		ctx.Next()
	}
}
