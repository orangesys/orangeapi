package main

import (
	_ "os"
	"log"

	"github.com/orangesys/orangeapi/server"
	"github.com/orangesys/orangeapi/k8s"
)

func main() {
		printVersion()
		log.Println("Starting orangeapi...")
		k8s.WaitForKubernetesProxy()
		server.Run()
}
