package core

type APIConfig struct {
	URL   string
	Token string
}

func NewAPIConfig(url, token string) APIConfig {
	return APIConfig{
		URL:   url,
		Token: token,
	}
}
