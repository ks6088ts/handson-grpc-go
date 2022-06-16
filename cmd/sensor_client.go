/*
Copyright © 2022 ks6088ts

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"log"

	"github.com/ks6088ts/handson-grpc-go/services/sensor/client"

	"github.com/spf13/cobra"
)

// sensorClientCmd represents the client command
var sensorClientCmd = &cobra.Command{
	Use:   "client",
	Short: "run sensor client",
	Long:  `run sensor client`,
	Run: func(cmd *cobra.Command, args []string) {
		addr, err := cmd.Flags().GetString("addr")
		if err != nil {
			log.Fatalf("failed to parse `addr`: %v", err)
		}

		client, err := client.NewSensorClient(addr)
		if err != nil {
			log.Fatalf("fail to create client: %v", err)
		}

		// GetSensorState
		state, err := client.GetSensorState()
		if err != nil {
			log.Fatalf("client.GetSensorState failed: %v", err)
		}
		log.Printf("state: %v", state)

		// GetSensorStates
		states, err := client.GetSensorStates()
		if err != nil {
			log.Fatalf("client.GetSensorState failed: %v", err)
		}
		for idx, state := range states {
			log.Printf("state[%v] = %v", idx, state)
		}
	},
}

func init() {
	sensorCmd.AddCommand(sensorClientCmd)

	sensorClientCmd.Flags().StringP("addr", "a", "localhost:50051", "the address to connect to")
}
