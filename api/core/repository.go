package core

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Repository struct {
	database *mongo.Database
	logger   *log.Logger
}

func NewRepository(database *mongo.Database, logger *log.Logger) *Repository {
	return &Repository{
		database: database,
		logger:   logger,
	}
}

func (repo *Repository) CreateMany(items []DetailedItem) error {
	// Crear una lista de operaciones de escritura para cada elemento.
	var operations []mongo.WriteModel

	for _, item := range items {
		// Filtro para buscar elementos existentes por su ID.
		filter := bson.D{{"id", item.ID}}

		// Actualiza el elemento si existe, o lo crea si no existe.
		update := bson.D{{"$set", item}}
		upsert := true

		operation := mongo.NewUpdateOneModel()
		operation.Filter = filter
		operation.Update = update
		operation.Upsert = &upsert

		// Agrega la operación a la lista de operaciones.
		operations = append(operations, operation)
	}

	// Realiza las operaciones de escritura en la base de datos.
	_, err := repo.database.Collection("items").BulkWrite(context.TODO(), operations)
	if err != nil {
		repo.logger.Printf("Error creating/updating items: %s", err)
		return err
	}

	repo.logger.Printf("%d items created/updated", len(items))
	return nil
}

func (repo *Repository) Create(item DetailedItem) error {
	if _, err := repo.database.Collection("item").InsertOne(context.TODO(), item); err != nil {
		repo.logger.Printf("Error creating item: %s", err)
		return err
	}

	repo.logger.Printf("Item created: %s", item)
	return nil
}

func (repo *Repository) GetAll(page int, pageSize int) ([]DetailedItem, error) {
	// Calcular el valor de "skip" para la paginación
	skip := (page - 1) * pageSize

	// Opciones para la consulta
	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(pageSize))

	// Realizar la consulta
	cursor, err := repo.database.Collection("items").Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		repo.logger.Printf("Error fetching items: %s", err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var items []DetailedItem

	// Iterar a través del cursor y decodificar los resultados
	for cursor.Next(context.TODO()) {
		var item DetailedItem
		if err := cursor.Decode(&item); err != nil {
			repo.logger.Printf("Error decoding item: %s", err)
			return nil, err
		}
		items = append(items, item)
	}

	if err := cursor.Err(); err != nil {
		repo.logger.Printf("Cursor error: %s", err)
		return nil, err
	}

	return items, nil
}
