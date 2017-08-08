package controller

import (
	"fmt"

	"github.com/orangesys/orangeapi/pkg/common"
	"github.com/orangesys/orangeapi/pkg/config"
	"github.com/orangesys/orangeapi/pkg/firebase"
	"github.com/orangesys/orangeapi/pkg/k8s"
	"github.com/orangesys/orangeapi/pkg/kong"
	"github.com/orangesys/orangeapi/pkg/wheel"
)

func createKongAPIPlugin(config *config.KongConfiguration, name, writepassword string) error {
	client := kong.NewClient(nil, config)
	influxdbAPI := &kong.API{
		Name:        name + "-influxdb",
		UpstreamURL: "http://" + name + "-i-influxdb.default",
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
		Name:        name + "-grafana",
		UpstreamURL: "http://" + name + "-g-grafana.default",
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
			Generator:  "tracker",
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
			AddQueryString:    querystring,
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

// createKongConsumerWithJwt create kong consumer jwt token
func createKongConsumerWithJwt(config *config.KongConfiguration, name string) (string, string, error) {
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

	_k, _ := common.UUID()
	_s, _ := common.UUID()
	generateConfig := &kong.JWTCredential{
		Key:    _k,
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

// CreateOrangesys create influxdb and grafana
func CreateOrangesys(name, retention, pvcsize string) (string, error) {
	namespace := "default"
	data := "write-password"
	key := name + "-i-influxdb"

	i := wheel.CreateInfluxDB{
		Name:      name,
		Retention: retention,
		PVCSize:   pvcsize,
	}
	if err := i.WheelInfluxdb(); err != nil {
		return "", fmt.Errorf("%s %s", "can not deploy influxdb with", name)
	}

	g := wheel.CreateGrafana{
		Name: name,
	}

	if err := g.WheelGrafana(); err != nil {
		return "", fmt.Errorf("%s %s", "can not deploy grafana with", name)
	}
	writepassword, err := k8s.GetSecret(namespace, key, data)
	if err != nil {
		return "", err
	}
	return writepassword, nil
}

// CreateConsumer create consumer with kong
func CreateConsumer(name, retention, pvcsize, uuid string) error {
	//	name := "rlxebz"
	//	wp := "mypassword"
	//	uuid := "iGzNX6QzfudVlwKtR8CQCj0itIU2"

	wp, err := CreateOrangesys(name, retention, pvcsize)
	if err != nil {
		fmt.Println("can not get grafana write password")
		return err
	}

	kongconfig, err := config.LoadKongConfig()
	if err != nil {
		return err
	}

	if err := createKongAPIPlugin(kongconfig, name, wp); err != nil {
		return err
	}

	key, secret, err := createKongConsumerWithJwt(kongconfig, name)
	if err != nil {
		return err
	}
	consumer := common.Consumer{
		Iss:    key,
		Secret: secret,
	}
	consumerJwtToken, err := consumer.CreateToken()
	if err != nil {
		return err
	}

	firebaseconfig, err := config.LoadFirebaseConfig()

	if err != nil {
		return err
	}
	user := firebase.FirebaseConfiguration{
		Config:     firebaseconfig,
		UUID:       uuid,
		ConsumerID: name,
		Token:      consumerJwtToken,
	}
	if err := user.SaveToken(); err != nil {
		return err
	}
	return nil
}
