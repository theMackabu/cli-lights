package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"lights/helpers"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type Response struct {
	Method string `json:"method"`
	ENV    string `json:"env"`
	Result Result `json:"result"`
}

type Result struct {
	MAC     string `json:"mac"`
	Rssi    int64  `json:"rssi"`
	State   bool   `json:"state"`
	SceneID int64  `json:"sceneId"`
	Temp    int64  `json:"temp"`
	Dimming int64  `json:"dimming"`
}

var response Response

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "check the status of your bulb",
	Run: func(cmd *cobra.Command, args []string) {
		ocelet, _ := cmd.Flags().GetString("ocelet")
		debug, _ := cmd.Flags().GetBool("debug")
		host := fmt.Sprintf("192.168.86.%s", ocelet)
		boldWhite := color.New(color.FgWhite, color.Bold)

		res := helpers.RunClient(fmt.Sprintf("%s:%s", host, "38899"), `{"method":"getPilot"}`)
		n := bytes.Index(res[:], []byte{0})
		json.Unmarshal(res[:n], &response)
		boldWhite.Printf("[%s] ", host)

		if debug {
			color.Cyan(string(res))
		} else {
			if response.Result.State {
				color.Cyan("on")
			} else {
				color.Cyan("off")
			}

		}

	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	statusCmd.Flags().StringP("ocelet", "o", "", `lightbulb IP ending ocelet ("01")`)
	statusCmd.Flags().BoolP("debug", "d", false, "enable debug mode")
	statusCmd.MarkFlagRequired("ocelet")
}
