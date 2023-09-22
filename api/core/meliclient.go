package core

import (
	"net/http"
)

type MeliClient struct {
	BaseURL string
	Token   string
	Client  *http.Client
}

func NewMeliClient(baseURL, token string) *MeliClient {
	return &MeliClient{
		BaseURL: baseURL,
		Token:   token,
		Client:  &http.Client{},
	}
}

func (mc *MeliClient) Get(endpoint string) (*http.Response, error) {
	req, err := http.NewRequest("GET", mc.BaseURL+endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+mc.Token)

	return mc.Client.Do(req)
}
