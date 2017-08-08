package kong

import (
	_ "fmt"
	"net/http"

	"github.com/orangesys/orangeapi/pkg/config"
)

type Client struct {
	APIService      *APIService
	ConsumerService *ConsumerService
	PluginService   *PluginService
	JWTService      *JWTService
}

// NewClient returns a new Client
func NewClient(httpClient *http.Client, config *config.KongConfiguration) *Client {
	return &Client{
		APIService:      NewAPIService(httpClient, config),
		ConsumerService: NewConsumerService(httpClient, config),
		PluginService:   NewPluginService(httpClient, config),
		JWTService:      NewJWTService(httpClient, config),
	}
}
