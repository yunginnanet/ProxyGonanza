package main

import (
	"errors"
	"fmt"
	"os"

	"git.tcp.direct/kayos/proxygonanza"
	"git.tcp.direct/kayos/proxygonanza/internal"
)

func init() {
	internal.ParseArgs()
}

var ErrNoNonMatchingIPs = errors.New("no non-matching IPs found or deleted")

func purge(c *proxygonanza.APIClient) error {
	println("clearing all auth IPs that do not match current IP...")
	deleted, err := c.DeleteOtherAuthIPs()
	if err != nil {
		return fmt.Errorf("failed to delete other auth IPs: %w", err)
	}
	if len(deleted) == 0 {
		return ErrNoNonMatchingIPs
	}
	println("deleted other auth IP IDs:")
	for _, del := range deleted {
		fmt.Println(del)
	}
	return nil
}

func main() {
	c := proxygonanza.NewAPIClient(internal.APIKey)
	if internal.Purge {
		if err := purge(c); err != nil && !errors.Is(err, ErrNoNonMatchingIPs) {
			println(err.Error())
			os.Exit(1)
		}
	}
	println("adding current IP to all packages...")
	count := c.AddCurrentIPtoAllPackages()
	if count == 0 {
		println("all authentication IPs are already set")
		os.Exit(0)
	}
	fmt.Printf("successfully added your external IP to %d packages\n", count)
}
