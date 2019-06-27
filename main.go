package main

import (
	"flag"
	"fmt"

	//"./cmd"
	"github.com/freewu/blogo/cmd"
)

func main() {

	flag.Usage = cmd.Usage
	flag.Parse()
	args := flag.Args()

	fmt.Print(len(args))
	if len(args) < 1 {
		cmd.Usage()
		//os.Exit(2)
		return
	}

	if args[0] == "help" {
		cmd.Help(args[1:])
		return
	}

	fmt.Println("other")
	return
}
