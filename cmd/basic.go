package cmd

import (
	"fmt"
	"lights/helpers"
	"strconv"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var basicCmd = &cobra.Command{
	Use:   "basic",
	Short: "warm/cold color range 0-255",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		ocelet, _ := cmd.Flags().GetString("ocelet")
		host := fmt.Sprintf("192.168.86.%s", ocelet)
		boldWhite := color.New(color.FgWhite, color.Bold)
		value, _ := strconv.Atoi(args[1])

		if args[0] == "warm" || args[0] == "cold" {
			if value > 255 || value < 0 {
				color.Red("invalid %s value %v", args[0], color.YellowString(args[1]))
			} else {
				if args[0] == "warm" {
					helpers.RunClient(fmt.Sprintf("%s:%s", host, "38899"), fmt.Sprintf(helpers.Warm, args[1]))
					boldWhite.Printf("[%s] ", host)
					color.Blue("set light to %v", color.YellowString(fmt.Sprintf("%s [warm]", args[1])))
				}
				if args[0] == "cold" {
					helpers.RunClient(fmt.Sprintf("%s:%s", host, "38899"), fmt.Sprintf(helpers.Cold, args[1]))
					boldWhite.Printf("[%s] ", host)
					color.Blue("set light to %v", color.CyanString(fmt.Sprintf("%s [cold]", args[1])))
				}
			}
		} else {
			color.Red("invalid value %v", color.YellowString(args[0]))
		}

	},
}

func init() {
	rootCmd.AddCommand(basicCmd)

	basicCmd.Flags().StringP("ocelet", "o", "", `lightbulb IP ending ocelet ("01")`)
	basicCmd.MarkFlagRequired("ocelet")
}
