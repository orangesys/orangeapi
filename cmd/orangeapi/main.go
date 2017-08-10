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
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	// RootCmd represents the base command when called without any subcommands
	var RootCmd = &cobra.Command{
		Use:   "orangeapi",
		Short: "Orangeapi is an API System",
		Long:  `This is orangesys api system , auto create orangesys.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: RunServer,
	}
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
