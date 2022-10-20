package cmd

import (
	"fmt"
	"lights/helpers"
	"strconv"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var colorCmd = &cobra.Command{
	Use:   "color",
	Short: "RGB color range 0-255",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		ocelet, _ := cmd.Flags().GetString("ocelet")
		host := fmt.Sprintf("192.168.86.%s", ocelet)
		boldWhite := color.New(color.FgWhite, color.Bold)
		red, _ := strconv.Atoi(args[0])
		green, _ := strconv.Atoi(args[1])
		blue, _ := strconv.Atoi(args[2])

		if red > 255 || red < 0 || green > 255 || green < 0 || blue > 255 || blue < 0 {
			color.Red("invalid color range")
		} else {
			helpers.RunClient(fmt.Sprintf("%s:%s", host, "38899"), fmt.Sprintf(helpers.Warm, args[1]))
			boldWhite.Printf("[%s] ", host)
			color.Yellow("set color to \033[1;37mr:%v \033[1;37mg:%v \033[1;37mb:%v", color.RedString(args[0]), color.GreenString(args[1]), color.BlueString(args[2]))
		}
	},
}

func init() {
	rootCmd.AddCommand(colorCmd)

	colorCmd.Flags().StringP("ocelet", "o", "", `lightbulb IP ending ocelet ("01")`)
	colorCmd.MarkFlagRequired("ocelet")
}
