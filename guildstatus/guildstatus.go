package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/jo3-l/yag-status/botstatus"
)

func main() {
	flag.Parse()
	s := flag.Arg(0)
	id, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("no/invalid guild ID provided")
	}

	status := botstatus.GetStatus()
	shardID := getGuildShard(id, status.TotalShards)

	for _, host := range status.HostStatuses {
		for _, node := range host.Nodes {
			for _, shard := range node.Shards {
				if shard.ShardID == shardID {
					log.Printf("Guild %d gateway status (shard %d): %s", id, shardID, shard.ConnStatus)
					return
				}
			}
		}
	}
}

// https://discord.com/developers/docs/topics/gateway#sharding-sharding-formula
func getGuildShard(guild, totalShards int) int {
	return (guild >> 22) % totalShards
}
