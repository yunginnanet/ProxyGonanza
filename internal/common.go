package internal

import (
	"fmt"
	"net/http"
	"os"
)

var (
	Debug  = false
	APIKey = os.Getenv("PROXY_BONANZA")
	Purge  = false
)

func ParseArgs() {
	if len(os.Args) < 2 && len(APIKey) < 2 {
		fmt.Printf("\t  ~*~ ProxyGonanza ~*~ \nhttps://git.tcp.direct/kayos/proxygonanza\n\nFatal: missing API Key \n\nUsage: %s [--verbose|-v] '<apikey>'\n\n", os.Args[0])
		os.Exit(1)
	}
	for i, arg := range os.Args {
		if i == 0 {
			continue
		}
		switch arg {
		case "-d", "--debug", "-v", "--verbose":
			Debug = true
		case "-p", "--purge":
			Purge = true
		default:
			APIKey = arg
		}
	}

}

// CloseBody is crude error handling for any potential errors closing the response body.
func CloseBody(res *http.Response) {
	if res == nil || res.Body == nil {
		return
	}
	err := res.Body.Close()
	if err != nil {
		println("WARN: ProxyGonanza failed to close body for request to ",
			res.Request.RequestURI+": "+err.Error())
	}
}
