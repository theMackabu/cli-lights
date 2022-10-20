package cmd

import (
	"fmt"
	"lights/helpers"
	"strconv"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var brightnessCmd = &cobra.Command{
	Use:   "brightness",
	Short: "sets the dimmer of the bulb in percent.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ocelet, _ := cmd.Flags().GetString("ocelet")
		host := fmt.Sprintf("192.168.86.%s", ocelet)
		boldWhite := color.New(color.FgWhite, color.Bold)
		brightness, _ := strconv.Atoi(args[0])

		if brightness > 100 || brightness < 0 {
			color.Red("invalid brightness value %v", color.YellowString(fmt.Sprintf("%s%%", args[0])))
		} else {
			helpers.RunClient(fmt.Sprintf("%s:%s", host, "38899"), fmt.Sprintf(helpers.Dimming, args[0]))
			boldWhite.Printf("[%s] ", host)
			color.Blue("set brightness to %v", color.GreenString(fmt.Sprintf("%s%%", args[0])))
		}

	},
}

func init() {
	rootCmd.AddCommand(brightnessCmd)

	brightnessCmd.Flags().StringP("ocelet", "o", "", `lightbulb IP ending ocelet ("01")`)
	brightnessCmd.MarkFlagRequired("ocelet")
}
