package main

import (
	"fmt"
	"os"
//        "reflect"

	_ "github.com/orangesys/orangeapi/helm"
	_ "github.com/orangesys/orangeapi/k8s"
	"github.com/orangesys/orangeapi/kong"
	"github.com/orangesys/orangeapi/config"
)

func create_kong_api_plugin(config *config.KongConfiguration, name, writepassword string) error {
	client := kong.NewClient(nil, config)
	influxdbAPI := &kong.API{
	    Name: name + "-influxdb",
	    UpstreamURL: "http://" + name + "-influxdb.default",
	    RequestHost: name + ".i.orangesys.io",
	}
	_, iresp, err := client.APIService.Create(influxdbAPI)
	if iresp.StatusCode != 201 {
		return fmt.Errorf("%s", "can not create influxdb api")
	}
	if err != nil {
		return err
	}

	grafanaAPI := &kong.API{
	    Name: name + "-grafana",
	    UpstreamURL: "http://" + name + "-grafana.default",
	    RequestHost: name + ".g.orangesys.io",
	}
	_, gresp, err := client.APIService.Create(grafanaAPI)
	if gresp.StatusCode != 201 {
		return fmt.Errorf("%s", "can not create grafana api")
	}
	if err != nil {
		return err
	}

	apiName := name + "-influxdb"
	JWTPlugin := &kong.Plugin{
	    Name: "jwt",
	}
	_, jwtresp, err := client.PluginService.Create(JWTPlugin, apiName)
	if jwtresp.StatusCode != 201 {
		return fmt.Errorf("%s %s", "can not create jwt plugin with api", apiName)
	}
	if err != nil {
		return err
	}

	correlationIDPlugin := &kong.Plugin{
	    Name: "correlation-id",
	    Config: kong.CorrelationIDPluginConfig{
			HeaderName: "Orangesys-Request-ID",
			Generator: "tracker",
	    },
	}
	_, cidresp, err := client.PluginService.Create(correlationIDPlugin, apiName)
	if cidresp.StatusCode != 201 {
		return fmt.Errorf("%s %s", "can not create correlation-id plugin with api", apiName)
	}
	if err != nil {
		return err
	}

	querystring := "u:_write,p:" + writepassword
	requesttransformerPlugin := &kong.Plugin{
	    Name: "request-transformer",
	    Config: kong.RequestTransformerPluginConfig{
			RemoveQueryString: "jwt",
			AddQueryString: querystring,
            },
        }
	_, rfresp, err := client.PluginService.Create(requesttransformerPlugin, apiName)
	if rfresp.StatusCode != 201 {
		return fmt.Errorf("%s %s", "can not create request-transformer plugin with api", apiName)
	}
	if err != nil {
		return err
	}
	return nil
}

func create_kong_consumer_with_jwt(config *config.KongConfiguration, name string) (string, string, error) {
	client := kong.NewClient(nil, config)
	generateConsumer := &kong.Consumer{
	    Username: name,
	}
	_, cresp, err := client.ConsumerService.Create(generateConsumer)
	if cresp.StatusCode != 201 {
		return "", "", fmt.Errorf("%s %s", "can not create consumer", name)
	}
	if err != nil {
		return "", "", err
	}

	_k, _ := kong.UUID()
	_s, _ := kong.UUID()
	generateConfig := &kong.JWTCredential{
		Key: _k,
		Secret: _s,
	}
	_, jwtresp, err := client.JWTService.Create(name, generateConfig)
	if jwtresp.StatusCode != 201 {
		return "", "", fmt.Errorf("%s %s", "can not create jwt plugin with api", name)
	}
	if err != nil {
		return "", "", err
	}
	return _k, _s, nil
}

func main() {
	config, err := config.LoadKongConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
	name := "rlxebz"
	wp := "mypassword"

	err = create_kong_api_plugin(config, name, wp)
        if err != nil {
		fmt.Println(err)
	//	fmt.Errorf("can not create api with kong", err)
		os.Exit(1)
	}
	key, secret, cerr := create_kong_consumer_with_jwt(config, name)
	if cerr != nil {
		fmt.Println(err)
	//	fmt.Errorf("can not create api with kong", err)
		os.Exit(1)
	}
	fmt.Printf("key is %s, secret is %s\n", key, secret)
}
