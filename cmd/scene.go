package cmd

import (
	"fmt"
	"lights/helpers"
	"strconv"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var sceneCmd = &cobra.Command{
	Use:   "scene",
	Short: "calls one of the predefined scenes (int from 1 to 32)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ocelet, _ := cmd.Flags().GetString("ocelet")
		host := fmt.Sprintf("192.168.86.%s", ocelet)
		boldWhite := color.New(color.FgWhite, color.Bold)
		scene, _ := strconv.Atoi(args[0])

		if scene > 32 || scene < 1 {
			color.Red("invalid scene %v", color.YellowString(args[0]))
		} else {
			helpers.RunClient(fmt.Sprintf("%s:%s", host, "38899"), fmt.Sprintf(helpers.Scene, args[0]))
			boldWhite.Printf("[%s] ", host)
			color.Green("light scene changed to %s", helpers.Scenes[scene-1])
		}

	},
}

func init() {
	rootCmd.AddCommand(sceneCmd)

	sceneCmd.Flags().StringP("ocelet", "o", "", `lightbulb IP ending ocelet ("01")`)
	sceneCmd.MarkFlagRequired("ocelet")
}
