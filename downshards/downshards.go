package main

import (
	"fmt"

	"github.com/jo3-l/yag-status/botstatus"
)

func main() {
	status := botstatus.GetStatus()
	fmt.Printf("%d/%d shards are down\n", len(status.OfflineShards), status.TotalShards)
}
