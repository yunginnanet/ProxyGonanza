package example

import (
	"fmt"
	"os"
)

var (
	Debug = false
	APIKey string
)

func ParseArgs() {
	if len(os.Args) < 2 {
		fmt.Printf("\t  ~*~ ProxyGonanza ~*~ \nhttps://git.tcp.direct/kayos/proxygonanza\n\nFatal: missing API Key \n\nUsage: %s [--verbose|-v] '<apikey>'\n\n", os.Args[0])
		os.Exit(1)
	}
	for _, arg := range os.Args {
		switch arg {
		case "-d", "--debug", "-v", "--verbose":
			Debug = true
			break
		default:
			APIKey = arg
		}

	}
}
