package main

import (
	"fmt"
	"os"

	"github.com/orangesys/orangeapi/kong"
	"github.com/orangesys/orangeapi/helm"
	"github.com/orangesys/orangeapi/k8s"
	"github.com/orangesys/orangeapi/config"
)


func create_kong_api(config *config.KongConfiguration) {
	client := kong.NewClient(nil, config)
        apis, _, _ := client.APIService.List()
        fmt.Printf("apis is %+v \n", apis)
	body := &kong.API{
	    Name: "test02",
	    UpstreamURL: "http://test02.inside",
	    RequestHost: "test02.outside",
	}
	_, resp, _ := client.APIService.Create(body)
	fmt.Printf("api resp is %+v \n\n", resp)
}

func create_kong_consumer(config *config.KongConfiguration) {
	client := kong.NewClient(nil, config)
	body := &kong.Consumer{
	    Username: "test02",
	}
	_, resp, _ := client.ConsumerService.Create(body)
	fmt.Printf("consumer resp is %+v \n\n", resp)
}

func create_kong_plugin_jwt(config *config.KongConfiguration) {
	client := kong.NewClient(nil, config)
	apiName := "test02"
	generatePlugin := &kong.Plugin{
	    Name: "jwt",
	}
	plugin, resp, _ := client.PluginService.Create(generatePlugin, apiName)
	fmt.Printf("plugin is %+v \n\n resp is %+v \n\n\n ", plugin, resp)
}

func config_kong_plugin_jwt(config *config.KongConfiguration) {
	client := kong.NewClient(nil, config)
	consumerName := "test02"
	_k, _ := kong.UUID()
	_s, _ := kong.UUID()
	generateConfig := &kong.JWTCredential{
		Key: _k,
		Secret: _s,
	}
	credential, resp, err := client.JWTService.Create(consumerName, generateConfig)
	fmt.Printf("Create JWT credential :\n%v\n%v\n%v\n%v\n", generateConfig, credential, resp, err)
}

func config_kong_plugin_correlationid(config *config.KongConfiguration) {
	client := kong.NewClient(nil, config)
	apiName := "test02"
	generatePlugin := &kong.Plugin{
	    Name: "correlation-id",
	    Config: kong.CorrelationIDPluginConfig{
			HeaderName: "Orangesys-Request-ID",
			Generator: "tracker",
	    },
	}
	plugin, resp, err := client.PluginService.Create(generatePlugin, apiName)
	fmt.Printf("Create C-ID plugin :\n%v\n%v\n%v\n%v\n", generatePlugin, plugin, resp, err)
}

func config_kong_plugin_request_transformer(config *config.KongConfiguration) {
	client := kong.NewClient(nil, config)
	apiName := "test02"
	writepassword := "passwordwrite"
	querystring := "u:_write,p:" + writepassword
	generatePlugin := &kong.Plugin{
	    Name: "request-transformer",
	    Config: kong.RequestTransformerPluginConfig{
			RemoveQueryString: "jwt",
			AddQueryString: querystring,
            },
        }
	plugin, resp, err := client.PluginService.Create(generatePlugin, apiName)
	fmt.Printf("Create Request-Transformer plugin :\n%v\n%v\n%v\n%v\n", generatePlugin, plugin, resp, err)
}

func main() {
	config, err := config.LoadKongConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
	create_kong_api(config)
	create_kong_consumer(config)
	create_kong_plugin_jwt(config)
	config_kong_plugin_jwt(config)
	config_kong_plugin_correlationid(config)
	config_kong_plugin_request_transformer(config)

//	client := kong.NewClient(nil, config)
//
//	api, _, _ := client.APIService.Get("mockbin")
//	fmt.Printf("API:\n%v\n", api)
//
//	plugins, _, _ := client.PluginService.List("mockbin")
//	fmt.Printf("Plugins:\n%v\n", plugins)
//
//	plugin, _, _ := client.PluginService.Get("8e7459c9-0e4e-4307-828b-f27cf7574c77", "mockbin")
//	fmt.Printf("Plugin:\n%v\n", plugin)

//	enablePlugins, resp, err := client.PluginService.GetEnabledPlugins()
//	fmt.Printf("Enable Plugins :\n%v\n%v\n%v\n", enablePlugins, resp, err)

//	generatePlugin := &kong.Plugin{
//		Name: "oauth2",
//		Config: kong.OAuth2PluginConfig{
//			EnableClientCredentials: true,
//		},
//	}
//	plugin, resp, err := client.PluginService.CreateOAuth(generatePlugin, "mockbin")
//	fmt.Printf("Create Plugin :\n%v\n%v\n%v\n%v\n", generatePlugin, plugin, resp, err)


}
