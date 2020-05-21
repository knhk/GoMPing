package main

import (
	"flag"
	"fmt"
	"github.com/knhk/GoMPing/internal/app/gomping"
)

var fHelp bool

func main() {
	fConfig := flag.String("conf", "gomping.yml", "config file path")
	flag.BoolVar(&fHelp, "help", false, "help command")
	flag.Parse()

	if fHelp {
		fmt.Print(helpMessage)
		return
	}

	gomping.Run(*fConfig)
}

const helpMessage = `
GoMPing cli options

 -help : show this help messages.
 -conf [FILE]: default config is "./gomping.yml". 

`