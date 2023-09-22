package core

import "fmt"

func startProcessing(repository *Repository, filePath, fileType string, apiConfig APIConfig) (string, error) {
	processor := NewFileProcessor(filePath, fileType, apiConfig)
	items, err := processor.Process()
	if err != nil {
		fmt.Printf("Error processing file: %s", err)
		return "", err
	}

	err = repository.CreateMany(items)
	if err != nil {
		fmt.Printf("Error creating items: %s", err)
		return "", err
	}

	return "File processed successfully", nil
}
