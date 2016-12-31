package main

import (
	"fmt"
	"os"
	"log"

	"github.com/orangesys/orangeapi/server"
	"github.com/orangesys/orangeapi/k8s"
	"github.com/orangesys/orangeapi/config"
)

func main() {
		if firebaseconfig, err := config.LoadFirebaseConfig()
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
