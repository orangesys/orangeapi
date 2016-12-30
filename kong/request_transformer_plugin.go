package kong

type RequestTransformerPluginConfig struct {
	RemoveQueryString             string `json:"remove.querystring,omitempty"`
	AddQueryString                string `json:"add.querystring,omitempty"`
}
