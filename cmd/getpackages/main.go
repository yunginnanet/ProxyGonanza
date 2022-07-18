package main

import (
	"encoding/json"
	"fmt"

	"git.tcp.direct/kayos/proxygonanza"
	"git.tcp.direct/kayos/proxygonanza/internal"
)

func init() {
	internal.ParseArgs()
}

func main() {
	c := proxygonanza.NewAPIClient(internal.APIKey)
	println("getting proxy packages...")

	packs, err := c.GetProxyPackages()
	if err != nil {
		println(err.Error())
		// return
	}
	fmt.Printf("\nfound %d proxy packages\n", len(packs))
	for _, p := range packs {
		pretty, _ := json.MarshalIndent(p, "", "\t")
		fmt.Print(string(pretty))
	}

	println("getting auth IPs...")
	authips, err := c.GetAuthIPs()
	if err != nil {
		println(err.Error())
		// return
	}

	fmt.Printf("\nfound %d auth IPs\n", len(authips))
	for _, i := range authips {
		pretty, _ := json.MarshalIndent(i, "", "\t")
		fmt.Print(string(pretty) + "\n")
	}
}
