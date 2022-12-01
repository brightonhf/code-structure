package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Specification structured configuration variables.
type Specification struct {
	Debug       bool   `envconfig:"DEBUG" default:"false"`
	LogLevel    string `envconfig:"LOG_LEVEL" default:"info"`
	Environment string `envconfig:"ENVIRONMENT" default:"staging"`
	HTTP        struct {
		BaseURL string `envconfig:"HTTP_BASE_URL"`
		Port    int    `envconfig:"HTTP_PORT" default:"8000"`
	}
	APIv2 struct {
		URL     string `envconfig:"API_V2_URL"`
		Timeout int    `envconfig:"API_V2_TIMEOUT" default:"5"`
	}
	RecipeService struct {
		URL string `envconfig:"RECIPE_SERVICE_URL"`
	}
	MenuService struct {
		URL string `envconfig:"MENU_SERVICE_URL"`
	}
	BoxContentService struct {
		URL string `envconfig:"BOX_CONTENT_SERVICE_URL"`
	}
	MenuAddonsService struct {
		URL string `envconfig:"MENU_ADDONS_SERVICE_URL"`
	}
	Client struct {
		TokenURL string `envconfig:"CLIENT_TOKEN_URL"`
		ID       string `envconfig:"CLIENT_ID"`
		Secret   string `envconfig:"CLIENT_SECRET"`
		User     string `envconfig:"CLIENT_USER"`
		Password string `envconfig:"CLIENT_PASSWORD"`
	}
	CircuitBreaker struct {
		MaxRequests      uint32        `envconfig:"CIRCUIT_BREAKER_MAX_REQUESTS" default:"5"`
		Interval         time.Duration `envconfig:"CIRCUIT_BREAKER_INTERVAL" default:"2m"`
		Timeout          time.Duration `envconfig:"CIRCUIT_BREAKER_TIMEOUT" default:"1m"`
		FailureThreshold uint32        `envconfig:"CIRCUIT_BREAKER_FAILURE_THRESHOLD" default:"10"`
	}
}

// JWTConfig holds config for jwt authentication
type JWTConfig struct {
	Alg           string        `envconfig:"JWT_ALG"`
	Key           string        `envconfig:"JWT_KEY"`
	RolesCacheTTL time.Duration `envconfig:"JWT_ROLES_CACHE_TTL" default:"5m"`
	Auth0Endpoint string        `envconfig:"AUTH0_ENDPOINT"`
	Auth0ClientID string        `envconfig:"AUTH0_CLIENT_ID"`
	AzureEndpoint string        `envconfig:"AZURE_ENDPOINT"`
	AzureClientID string        `envconfig:"AZURE_CLIENTID"`
}

// LoadEnv load config variables into Specification.
func LoadEnv() (*Specification, error) {
	var config Specification
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, err
	}

	return &config, err
}
