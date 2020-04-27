// Copyright © 2020 NAME HERE <EMAIL ADDRESS>
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

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"sync"
)

var mqttStoredMessages = "0"

var wg = sync.WaitGroup{}
// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("running que monitor..\n")
		clientid := viper.GetString("device.id")
		telPath := viper.GetString("telemetry.path")
		telPort := viper.GetString("telemetry.port")
		fmt.Printf("URL: http://localhost:%s%s",telPort,telPath)
		wg.Add(1)
		go InitMQTTStatsClient(clientid, &mqttStoredMessages)
		http.HandleFunc(telPath, handler)
		log.Fatal(http.ListenAndServe(":"+telPort, nil))

		//go func() {
		//	for {
		//		println(mqttStoredMessages)
		//		time.Sleep(2 * time.Second)
		//	}
		//}()


		wg.Wait()
	},
}

func handler(w http.ResponseWriter, r *http.Request,) {
	fmt.Fprintf(w, mqttStoredMessages)
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
