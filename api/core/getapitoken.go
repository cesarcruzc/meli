package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetApiToken() (string, error) {
	apiURL := os.Getenv("API_TOKEN_URL")

	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		return "", err
	}

	xApiKey := os.Getenv("X_API_KEY")

	fmt.Printf("X_API_KEY: %v\n", xApiKey)

	req.Header.Add(`Meli-X-Api-Key`, xApiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var responseMap map[string]interface{}
	json.Unmarshal(body, &responseMap)

	token, exists := responseMap["access_token"].(string)

	if !exists {
		return "", fmt.Errorf("Token not found in API response")
	}

	return token, nil
}
