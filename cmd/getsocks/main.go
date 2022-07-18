package main

import (
	"fmt"
	"os"

	"git.tcp.direct/kayos/proxygonanza"
	"git.tcp.direct/kayos/proxygonanza/internal"
)

func init() {
	internal.ParseArgs()
}

func main() {
	c := proxygonanza.NewAPIClient(internal.APIKey)
	if internal.Debug {
		c.Debug = true
		println("debug enabled")
	}
	socks, err := c.GetAllSOCKSIPsAndPorts()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	for _, line := range socks {
		fmt.Fprint(os.Stdout, line+"\n")
	}
}
