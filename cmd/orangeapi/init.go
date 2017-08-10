package main

import (
	log "github.com/rs/zerolog/log"

	"github.com/orangesys/orangeapi/pkg/config"
)

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	firebaseconfig, _ := config.LoadFirebaseConfig()

	if firebaseconfig.FirebaseAuth == "" {
		log.Fatal().Msg("cat not get FirebaseAuth")
	}
}
