package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/hyoinandout/hyoincoin/explorer"
	"github.com/hyoinandout/hyoincoin/rest"
)

func usage() {
	fmt.Printf("Welcome to hyoincoin\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port:		Set the PORT of the server\n")
	fmt.Printf("-mode:		Choose between 'html' and 'rest'\n\n")
	os.Exit(0)
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	case "all":
		go rest.Start(5000)
		explorer.Start(4000)
	default:
		usage()
	}
}
