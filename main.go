package main

import (
	log "github.com/rs/zerolog/log"

	"github.com/orangesys/orangeapi/pkg/config"
	"github.com/orangesys/orangeapi/pkg/k8s"
	"github.com/orangesys/orangeapi/pkg/server"
)

func main() {
	firebaseconfig, _ := config.LoadFirebaseConfig()

	if firebaseconfig.FirebaseAuth == "" {
		log.Fatal().Msg("cat not get FirebaseAuth")
	}
	printVersion()
	log.Info().Msg("Starting orangeapi")
	k8s.WaitForKubernetesProxy()
	server.Run()
}
