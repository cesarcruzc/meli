package core

import (
	"encoding/csv"
	"fmt"
	"os"
	"sync"
)

type CSVProcessor struct {
	path      string
	apiConfig APIConfig
}

func (p *CSVProcessor) Process() ([]DetailedItem, error) {
	file, err := os.Open(p.path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Leer la línea de encabezado y descartarla
	if _, err := reader.Read(); err != nil {
		return nil, fmt.Errorf("error reading header line: %w", err)
	}

	var meliItems []DetailedItem

	chunkSize := 20

	var wg sync.WaitGroup
	chunkChannel := make(chan []DetailedItem)

	for {
		// Leer un chunk de líneas
		chunk, err := p.readChunk(reader, chunkSize)
		if err != nil {
			return nil, err
		}

		// Si el chunk está vacío, hemos terminado de leer el archivo
		if len(chunk) == 0 {
			break
		}

		wg.Add(1)
		go func([][]string) {
			defer wg.Done()

			var items []fileItem
			// Procesar las líneas del chunk y convertirlas en []Item
			for _, line := range chunk {

				item := fileItem{
					Site: line[0],
					ID:   line[1],
				}
				items = append(items, item)
			}

			meliApi := NewWebService()
			meliClient := NewMeliClient(p.apiConfig.URL, p.apiConfig.Token)

			details, err := meliApi.fetchDetailsFromMercadoLibre(meliClient, items)
			if err != nil {
				return
			}

			chunkChannel <- details

		}(chunk)
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

// readChunk lee un número específico de líneas del CSV
func (p *CSVProcessor) readChunk(reader *csv.Reader, numLines int) ([][]string, error) {
	var chunk [][]string
	for i := 0; i < numLines; i++ {
		line, err := reader.Read()
		// Si err es io.EOF hemos llegado al final del archivo
		if err != nil {
			break
		}
		chunk = append(chunk, line)
	}
	return chunk, nil
}
