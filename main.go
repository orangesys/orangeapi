package main

import (
	"fmt"
	"os"

	"github.com/orangesys/orangeapi/kong"
	"github.com/orangesys/orangeapi/config"
)


func api(config *config.KongConfiguration) {
	client := kong.NewClient(nil, config)
	apis, _, _ := client.APIService.List()
	fmt.Printf("apis is %+v \n", apis)
	body := &kong.API{
	    Name: "test02",
	    UpstreamURL: "http://test02.inside",
	    RequestHost: "test02.outside",
	}
	_, resp, _ := client.APIService.Create(body)
	fmt.Printf("api resp is %+v \n", resp)
}

func consumer(config *config.KongConfiguration) {
	client := kong.NewClient(nil, config)
	body := &kong.Consumer{
	    Username: "test02",
	}
	_, resp, _ := client.ConsumerService.Create(body)
	fmt.Printf("consumer resp is %+v \n", resp)
}

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}

	consumer(config)
	api(config)
}
