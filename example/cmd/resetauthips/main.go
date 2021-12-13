package main

import (
	"fmt"
	"strconv"

	"git.tcp.direct/kayos/proxygonanza"
	"git.tcp.direct/kayos/proxygonanza/example"
)

func init() {
	example.ParseArgs()
}

func main() {
	println("clearing all auth IPs...")

	c := proxygonanza.NewAPIClient(example.APIKey)
	deleted, err := c.DeleteAllAuthIPs()
	if err != nil {
		println(err.Error())
		return
	}

	println("deleted " + strconv.Itoa(deleted) + " IPs successfully")


	println("adding current IP to all packages...")
	count := c.AddCurrentIPtoAllPackages()
	if count == 0 {
		println("failed!")
		return
	}

	fmt.Printf("successfully added your external IP to %d packages\n", count)
}
