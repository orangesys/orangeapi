package main

import (
	"fmt"
	"log"
	"os"

	"github.com/orangesys/orangeapi/pkg/config"
	"github.com/orangesys/orangeapi/pkg/k8s"
	"github.com/orangesys/orangeapi/pkg/server"
)

func main() {
	firebaseconfig, err := config.LoadFirebaseConfig()
	if err != nil {
		fmt.Println(err)
	}
	if firebaseconfig.FirebaseAuth == "" {
		log.Println("cat not get FirebaseAuth")
		os.Exit(1)
	}
	printVersion()
	log.Println("Starting orangeapi...")
	k8s.WaitForKubernetesProxy()
	server.Run()
}
