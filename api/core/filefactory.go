package core

type fileItem struct {
	Site string
	ID   string
}

type FileProcessor interface {
	Process() ([]DetailedItem, error)
}

func NewFileProcessor(path, ext string, apiconfig APIConfig) FileProcessor {
	switch ext {
	case ".csv":
		return &CSVProcessor{path: path, apiConfig: apiconfig}
	case ".jsonl":
		return &JSONProcessor{path: path, apiConfig: apiconfig}
	case ".txt":
		return &TextProcessor{path: path, apiConfig: apiconfig}
	default:
		return nil
	}
}
