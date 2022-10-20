package cmd

import (
	"fmt"
	"lights/helpers"
	"strconv"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var tempCmd = &cobra.Command{
	Use:   "temp",
	Short: "sets the color temperature in kelvins",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ocelet, _ := cmd.Flags().GetString("ocelet")
		host := fmt.Sprintf("192.168.86.%s", ocelet)
		boldWhite := color.New(color.FgWhite, color.Bold)
		value, _ := strconv.Atoi(args[0])

		if value > 6500 || value < 2200 {
			color.Red("invalid kelvin value %v", color.YellowString(args[0]))
		} else {
			helpers.RunClient(fmt.Sprintf("%s:%s", host, "38899"), fmt.Sprintf(helpers.Temp, args[0]))
			boldWhite.Printf("[%s] ", host)
			color.Blue("set light to %v", color.CyanString(fmt.Sprintf("%s [kelvin]", args[0])))
		}

	},
}

func init() {
	rootCmd.AddCommand(tempCmd)

	tempCmd.Flags().StringP("ocelet", "o", "", `lightbulb IP ending ocelet ("01")`)
	tempCmd.MarkFlagRequired("ocelet")
}
