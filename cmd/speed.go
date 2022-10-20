package cmd

import (
	"fmt"
	"lights/helpers"
	"strconv"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var speedCmd = &cobra.Command{
	Use:   "speed",
	Short: "sets the color changing speed in percent.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ocelet, _ := cmd.Flags().GetString("ocelet")
		host := fmt.Sprintf("192.168.86.%s", ocelet)
		boldWhite := color.New(color.FgWhite, color.Bold)
		speed, _ := strconv.Atoi(args[0])

		if speed > 100 || speed < 0 {
			color.Red("invalid speed value %v", color.YellowString(fmt.Sprintf("%s%%", args[0])))
		} else {
			helpers.RunClient(fmt.Sprintf("%s:%s", host, "38899"), fmt.Sprintf(helpers.Speed, args[0]))
			boldWhite.Printf("[%s] ", host)
			color.Blue("set speed to %v", color.CyanString(fmt.Sprintf("%s%%", args[0])))
		}

	},
}

func init() {
	rootCmd.AddCommand(speedCmd)

	speedCmd.Flags().StringP("ocelet", "o", "", `lightbulb IP ending ocelet ("01")`)
	speedCmd.MarkFlagRequired("ocelet")
}
