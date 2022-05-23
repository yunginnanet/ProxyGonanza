package main

import (
	"fmt"
	"git.tcp.direct/kayos/proxygonanza"
	"git.tcp.direct/kayos/proxygonanza/example"
	"strconv"
)

func init() {
	example.ParseArgs()
}

func purge(c *proxygonanza.APIClient) {
	println("clearing all auth IPs...")

	deleted, err := c.DeleteAllAuthIPs()
	if err != nil {
		println(err.Error())
		return
	}

	println("deleted " + strconv.Itoa(deleted) + " IPs successfully")
}

func main() {
	c := proxygonanza.NewAPIClient(example.APIKey)

	if example.Purge {
		purge(c)
	}

	println("adding current IP to all packages...")
	count := c.AddCurrentIPtoAllPackages()
	if count == 0 {
		println("failed!")
		return
	}

	fmt.Printf("successfully added your external IP to %d packages\n", count)
}
