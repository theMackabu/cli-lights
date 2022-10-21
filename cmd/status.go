package cmd

import (
	"fmt"
	"lights/helpers"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "check the status of your bulb",
	Run: func(cmd *cobra.Command, args []string) {
		ocelet, _ := cmd.Flags().GetString("ocelet")
		debug, _ := cmd.Flags().GetString("ocelet")
		host := fmt.Sprintf("192.168.86.%s", ocelet)
		boldWhite := color.New(color.FgWhite, color.Bold)

		res := helpers.RunClient(fmt.Sprintf("%s:%s", host, "38899"), `{"method":"getPilot"}`)
		boldWhite.Printf("[%s] ", host)

		if ()

		color.Blue(res)
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	statusCmd.Flags().StringP("ocelet", "o", "", `lightbulb IP ending ocelet ("01")`)
	statusCmd.Flags().BoolVarP("debug", "d", "", `enable debug mode`)
	statusCmd.MarkFlagRequired("ocelet")
}
