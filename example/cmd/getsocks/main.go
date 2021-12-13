package main

import (
	"os"

	"git.tcp.direct/kayos/proxygonanza"
	"git.tcp.direct/kayos/proxygonanza/example"
)

func init() {
	example.ParseArgs()
}

func main() {
	c := proxygonanza.NewAPIClient(example.APIKey)
	if example.Debug {
		c.Debug = true
		println("debug enabled")
	}
	socks, err := c.GetAllSOCKSIPsAndPorts()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	for _, line := range socks {
		println(line)
	}
}
