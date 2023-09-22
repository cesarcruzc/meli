package core

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
)

func ConsumeHandler(database *mongo.Database, logger *log.Logger) (string, error) {

	// Inicia el proceso de consumo de datos.
	repository := NewRepository(database, logger)
	filePath := os.Getenv("FILE_PATH")
	fileType := os.Getenv("FILE_TYPE")

	// Obtiene el token de la API, este proceso lo separe a un servicio externo al cual me autentico mediante api key para darle mayor seguridad.
	apiToken, err := GetApiToken()
	if err != nil {
		fmt.Printf("Error obteniendo token: %v\n", err)
		return "", err
	}

	apiConfig := APIConfig{
		URL:   os.Getenv("API_URL"),
		Token: apiToken,
	}

	status, err := startProcessing(repository, filePath, fileType, apiConfig)

	if err != nil {
		return status, err
	}

	return status, nil
}

func GetItemsHandler(database *mongo.Database, logger *log.Logger, page, pageSize int) ([]Item, error) {
	// Obtiene los items de la base de datos.
	repository := NewRepository(database, logger)

	detailedItems, err := repository.GetAll(page, pageSize)
	if err != nil {
		fmt.Printf("Error obteniendo elementos paginados: %v\n", err)
		return nil, err
	}

	// Realiza el mapeo de DetailedItem a Item
	var items []Item
	for _, detailedItem := range detailedItems {
		item := Item{
			Site:        detailedItem.Site,
			ID:          detailedItem.ID,
			Price:       detailedItem.Price,
			StartTime:   detailedItem.StartTime,
			Name:        fmt.Sprintf("%v", detailedItem.CategoryID),
			Description: detailedItem.CurrencyID,
			Nickname:    detailedItem.SellerID,
		}
		items = append(items, item)
	}

	return items, nil
}
