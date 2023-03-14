package config

type APIConfig struct {
	Port         string
	LoggingLevel string
	ServiceURI   string
}

func GetAPIConfig() *APIConfig {
	return &APIConfig{
		Port:         "",
		LoggingLevel: "",
		ServiceURI:   "",
	}
}
