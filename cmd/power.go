package cmd

import (
	"fmt"
	"lights/helpers"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var powerCmd = &cobra.Command{
	Use:   "power",
	Short: "changes whether it's on or off.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ocelet, _ := cmd.Flags().GetString("ocelet")
		host := fmt.Sprintf("192.168.86.%s", ocelet)
		boldWhite := color.New(color.FgWhite, color.Bold)

		if args[0] == "on" {
			helpers.RunClient(fmt.Sprintf("%s:%s", host, "38899"), fmt.Sprintf(helpers.Power, "true"))
			boldWhite.Printf("[%s] ", host)
			color.Green("light powered on")
		}
		if args[0] == "off" {
			helpers.RunClient(fmt.Sprintf("%s:%s", host, "38899"), fmt.Sprintf(helpers.Power, "false"))
			boldWhite.Printf("[%s] ", host)
			color.Magenta("light powered off")
		}
	},
}

func init() {
	rootCmd.AddCommand(powerCmd)

	powerCmd.Flags().StringP("ocelet", "o", "", `lightbulb IP ending ocelet ("01")`)
	powerCmd.MarkFlagRequired("ocelet")
}
