package main

import (
	"go.uber.org/zap"

	"github.com/orangesys/orangeapi/pkg/config"
	"github.com/orangesys/orangeapi/pkg/k8s"
	"github.com/orangesys/orangeapi/pkg/server"
)

func main() {
	log, _ := zap.NewProduction()
	firebaseconfig, _ := config.LoadFirebaseConfig()

	if firebaseconfig.FirebaseAuth == "" {
		log.Fatal("cat not get FirebaseAuth")
	}
	printVersion()
	log.Info("Starting orangeapi")
	k8s.WaitForKubernetesProxy()
	server.Run()
}
