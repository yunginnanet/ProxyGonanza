package main

import (
	"fmt"
	"os"
	"strconv"

	"git.tcp.direct/kayos/proxygonanza"
)

func main() {
	println("clearing all auth IPs...")

	c := proxygonanza.NewApiClient(os.Args[1])
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
