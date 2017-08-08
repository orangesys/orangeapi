package kong

type CorrelationIDPluginConfig struct {
	HeaderName      string `form:"header_name" json:"header_name,omitempty"`
	Generator       string `form:"generator" json:"generator,omitempty"`
	EchoDownstreams bool   `form:"echo_downstream" json:"echo_downstream,omitempty"`
}
