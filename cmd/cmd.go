package cmd

import "fmt"

func Usage() {
	fmt.Println("Usage")
	//utils.Tmpl(usageTemplate, commands.AvailableCommands)
}

func Help(args []string) {
	fmt.Println("help")
	// if len(args) == 0 {
	// 	Usage()
	// }
	// if len(args) != 1 {
	// 	utils.PrintErrorAndExit("Too many arguments", ErrorTemplate)
	// }

	// arg := args[0]

	// for _, cmd := range commands.AvailableCommands {
	// 	if cmd.Name() == arg {
	// 		utils.Tmpl(helpTemplate, cmd)
	// 		return
	// 	}
	// }
	// utils.PrintErrorAndExit("Unknown help topic", ErrorTemplate)
}
