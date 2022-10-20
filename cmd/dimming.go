package cmd

import (
	"fmt"
	"lights/helpers"

	"github.com/spf13/cobra"
)

var dimmingCmd = &cobra.Command{
	Use:   "dimming",
	Short: "changes the lights brightness.",
	Run: func(cmd *cobra.Command, args []string) {
		ocelet, _ := cmd.Flags().GetString("ocelet")
		brightness, _ := cmd.Flags().GetString("brightness")
		host := fmt.Sprintf("192.168.86.%s", ocelet)

		helpers.RunClient(fmt.Sprintf("%s:%s", host, "38899"), fmt.Sprintf(helpers.Dimming, brightness))
	},
}

func init() {
	rootCmd.AddCommand(dimmingCmd)

	dimmingCmd.Flags().StringP("ocelet", "o", "", `lightbulb IP ending ocelet ("01")`)
	dimmingCmd.Flags().StringP("brightness", "b", "", `brightness value ("100")`)
}
