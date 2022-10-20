package main

import "lights/cmd"

func main() {
	cmd.Execute()
}

// func main() {
// 	flag.Parse()

// 	if len(flag.Args()) < 1 {
// 		fmt.Println("lightbulb ending IP ocelet required")
// 		return
// 	}

// 	serverHost := fmt.Sprintf("192.168.86.%s", flag.Arg(0))
// 	fmt.Printf("%s:%s\n", serverHost, "38899")

// 	startClient(fmt.Sprintf("%s:%s", serverHost, "38899"), fmt.Sprintf(power, "false"))
// }
