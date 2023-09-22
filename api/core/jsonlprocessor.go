package core

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type JSONProcessor struct {
	path      string
	apiConfig APIConfig
}

func (p *JSONProcessor) Process() ([]DetailedItem, error) {
	file, err := os.Open(p.path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var meliItems []DetailedItem

	var wg sync.WaitGroup
	chunkChannel := make(chan []DetailedItem)

	scanner := bufio.NewScanner(file)

	chunkSize := 20
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)

		// Si se ha alcanzado el tamaño del chunk, procesar las líneas
		if len(lines) >= chunkSize {
			wg.Add(1)
			go func(lines []string) {
				defer wg.Done()

				var items []fileItem
				for _, line := range lines {
					var item fileItem
					if err := json.Unmarshal([]byte(line), &item); err != nil {
						fmt.Printf("Error decoding JSON line: %v\n", err)
						continue
					}
					items = append(items, item)
				}

				meliApi := NewWebService()
				meliClient := NewMeliClient(p.apiConfig.URL, p.apiConfig.Token)

				fmt.Printf("Fetching details for %d items\n", len(items))
				details, err := meliApi.fetchDetailsFromMercadoLibre(meliClient, items)
				if err != nil {
					return
				}

				chunkChannel <- details

			}(lines)

			// Limpiar el slice de líneas procesadas
			lines = nil
		}
	}

	// Procesar las líneas restantes si no se ha alcanzado el tamaño del chunk
	if len(lines) > 0 {
		wg.Add(1)
		go func(lines []string) {
			defer wg.Done()

			var items []fileItem
			for _, line := range lines {
				var item fileItem
				if err := json.Unmarshal([]byte(line), &item); err != nil {
					fmt.Printf("Error decoding JSON line: %v\n", err)
					continue
				}
				items = append(items, item)
			}

			meliApi := NewWebService()
			meliClient := NewMeliClient(p.apiConfig.URL, p.apiConfig.Token)

			fmt.Printf("Fetching details for %d items\n", len(items))
			details, err := meliApi.fetchDetailsFromMercadoLibre(meliClient, items)
			if err != nil {
				return
			}

			chunkChannel <- details

		}(lines)
	}

	go func() {
		wg.Wait()
		close(chunkChannel)
	}()

	// Recopilar resultados de las goroutines
	for meliItemsDetails := range chunkChannel {
		meliItems = append(meliItems, meliItemsDetails...)
	}

	return meliItems, nil
}
