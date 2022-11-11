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
	Warm    int64  `json:"w"`
	Cold    int64  `json:"c"`
	Red     int64  `json:"r"`
	Green   int64  `json:"g"`
	Blue    int64  `json:"b"`
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
		boldYellow := color.New(color.FgYellow, color.Bold)

		res := helpers.RunClient(fmt.Sprintf("%s:%s", host, "38899"), `{"method":"getPilot"}`)
		n := bytes.Index(res[:], []byte{0})
		json.Unmarshal(res[:n], &response)

		if debug {
			boldWhite.Printf("[%s] ", host)
			color.Cyan(string(res))
		} else {
			boldWhite.Printf("Information for %s:\n", host)
			boldYellow.Print("  - MAC Address: ")
			color.Blue(response.Result.MAC)
			boldYellow.Print("  - RSSi: ")
			color.Blue(fmt.Sprintf("%d", response.Result.Rssi))
			boldWhite.Printf("\nStatus for %s:\n", host)
			if response.Result.State {
				boldYellow.Print("  - Power: ")
				color.Green("on")
			} else {
				boldYellow.Print("  - Power: ")
				color.Red("off")
			}
			boldYellow.Print("  - Brightness: ")
			color.Cyan(fmt.Sprintf("%d%%", response.Result.Dimming))
			if response.Result.SceneID > 0 {
				boldYellow.Print("  - Scene: ")
				color.Magenta(helpers.Scenes[response.Result.SceneID-1])
			}
			if response.Result.Temp > 0 {
				boldYellow.Print("  - Color Temp [kelvin]: ")
				color.Cyan(fmt.Sprintf("%d", response.Result.Temp))
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
