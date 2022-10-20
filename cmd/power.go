package cmd

import (
	"fmt"
	"lights/helpers"

	"github.com/spf13/cobra"
)

var powerCmd = &cobra.Command{
	Use:   "power",
	Short: "changes the light state.",
	Run: func(cmd *cobra.Command, args []string) {
		ocelet, _ := cmd.Flags().GetString("ocelet")
		power, _ := cmd.Flags().GetString("power")
		host := fmt.Sprintf("192.168.86.%s", ocelet)

		if power == "on" {
			helpers.RunClient(fmt.Sprintf("%s:%s", host, "38899"), fmt.Sprintf(helpers.Power, "true"))
		}
		if power == "off" {
			helpers.RunClient(fmt.Sprintf("%s:%s", host, "38899"), fmt.Sprintf(helpers.Power, "false"))
		}
	},
}

func init() {
	rootCmd.AddCommand(powerCmd)

	powerCmd.Flags().StringP("ocelet", "o", "", `lightbulb IP ending ocelet ("01")`)
	powerCmd.Flags().StringP("power", "p", "", `power state ("on|off")`)
}
