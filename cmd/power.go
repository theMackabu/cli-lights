package cmd

import (
	"fmt"
	"lights/helpers"

	"github.com/spf13/cobra"
)

// dateCmd represents the date command
var powerCmd = &cobra.Command{
	Use:   "power",
	Short: "Prints the current date.",
	Run: func(cmd *cobra.Command, args []string) {
		ocelet, _ := cmd.Flags().GetString("ocelet")
		power, _ := cmd.Flags().GetString("power")
		host := fmt.Sprintf("192.168.86.%s", ocelet)

		helpers.RunClient(fmt.Sprintf("%s:%s", host, "38899"), fmt.Sprintf(helpers.Power, power))
	},
}

func init() {
	rootCmd.AddCommand(powerCmd)

	powerCmd.Flags().StringP("ocelet", "o", "21", "specify light IP ocelet")
	powerCmd.Flags().StringP("power", "p", "true", "light power state")
}
