// Copyright © 2017 gavin zhou <gavin@orangesys.io>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	log "github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/orangesys/orangeapi/pkg/k8s"
	"github.com/orangesys/orangeapi/pkg/server"
)

// RunServer is the run command to start orangeapi
func RunServer(cmd *cobra.Command, arg []string) {
	initConfig()
	log.Info().Msg("Starting orangeapi")
	k8s.WaitForKubernetesProxy()
	server.Run()
}
