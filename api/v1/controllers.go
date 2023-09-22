package v1

import (
	"github.com/cesarcruzc/meli/api/core"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"strconv"
)

// Root godoc
// @Summary Root
// @Description Entry point of the API
// @Tags meta
// @Produce  json
// @Router / [get]
func Root(ctx *gin.Context) {
	ctx.JSONP(
		http.StatusOK,
		gin.H{
			"message": "Meli Microservice",
			"status":  http.StatusOK,
		})
}

// HealthCheck godoc
// @Summary Health Check
// @Description Health check of the API
// @Tags meta
// @Produce  json
// @Router /health [get]
func HealthCheck(ctx *gin.Context) {
	ctx.JSONP(
		http.StatusOK,
		gin.H{
			"service": "ok",
			"status":  http.StatusOK,
		})
}

// StartProcess godoc
// @Summary Start Process
// @Description Start the process of consuming data from Mercado Libre and storing it in a database.
// @Tags process
// @Produce  json
// @Router /api/v1/process-file [post]
func StartProcess(ctx *gin.Context) {
	database := ctx.MustGet("database").(*mongo.Database)
	logger := ctx.MustGet("logger").(*log.Logger)

	// Inicia el proceso de consumo de datos.
	status, err := core.ConsumeHandler(database, logger)

	if err != nil {
		ctx.JSONP(
			http.StatusBadRequest,
			gin.H{
				"message": status,
				"status":  http.StatusBadRequest,
			})
		return
	}

	ctx.JSONP(
		http.StatusOK,
		gin.H{
			"message": status,
			"status":  http.StatusOK,
		})
}

// GetItems godoc
// @Summary Get Items
// @Description Get items from the database.
// @Tags items
// @Produce  json
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Page size" default(10)
// @Success 200 {object} []core.Item
// @Router /api/v1/items [get]
func GetItems(ctx *gin.Context) {
	database := ctx.MustGet("database").(*mongo.Database)
	logger := ctx.MustGet("logger").(*log.Logger)

	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil {
		// Manejar el error de conversión si es necesario
		ctx.JSON(400, gin.H{"error": "page debe ser un número entero válido"})
		return
	}

	pageSize, err := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	if err != nil {
		// Manejar el error de conversión si es necesario
		ctx.JSON(400, gin.H{"error": "pageSize debe ser un número entero válido"})
		return
	}

	// Obtiene los items de la base de datos.
	items, err := core.GetItemsHandler(database, logger, page, pageSize)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error":   err.Error(),
				"message": "Error obteniendo elementos",
			})
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"message": items,
			"status":  http.StatusOK,
		})
}

// GetApiToken godoc
// @Summary Get API Token
// @Description Get the API token to consume data from Mercado Libre.
// @Tags process
// @Produce  json
// @Router /api/v1/token [get]
func GetApiToken(ctx *gin.Context) {
	apiToken, err := core.GetApiToken()
	if err != nil {
		ctx.JSONP(
			http.StatusBadRequest,
			gin.H{
				"message": "Error obteniendo token",
				"status":  http.StatusBadRequest,
			})
		return
	}

	ctx.JSONP(
		http.StatusOK,
		gin.H{
			"message": apiToken,
			"status":  http.StatusOK,
		})
}
