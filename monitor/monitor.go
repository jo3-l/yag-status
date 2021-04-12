package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/jo3-l/yag-status/botstatus"
)

type webhookPostRequest struct {
	Embeds []statusEmbed `json:"embeds"`
}

type statusEmbed struct {
	Author      statusEmbedAuthor `json:"author"`
	Description string            `json:"description"`
	Color       int               `json:"color"`
	Timestamp   string            `json:"timestamp"`
}

type statusEmbedAuthor struct {
	IconURL string `json:"icon_url"`
	Name    string `json:"name"`
}

type config struct {
	WebhookURL string `json:"webhook_url"`
}

const yagpdbIconURL = "https://yagpdb.xyz/static/img/avatar.png"

var client http.Client = http.Client{Timeout: 5 * time.Second}
var webhookURL string

func init() {
	content, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal("error reading config file: ", err)
	}

	var conf config
	err = json.Unmarshal(content, &conf)
	if err != nil {
		log.Fatal("error unmarshaling config file into struct: ", err)
	}

	webhookURL = conf.WebhookURL
}

func main() {
	fmt.Println("monitoring status...")

	prevOfflineCount := 0

	ticker := time.NewTicker(10 * time.Second)
	for ; true; <-ticker.C {
		status := botstatus.GetStatus()
		offlineCount := len(status.OfflineShards)
		total := status.TotalShards
		online := total - offlineCount

		diff := prevOfflineCount - offlineCount

		var info string
		var color int
		switch {
		case diff > 0:
			info = fmt.Sprintf("`%d` shard(s) went back online (now %d/%d).", diff, online, total)
			color = 0x349721
		case diff == 0:
			info = fmt.Sprintf("No shards changed status since last check (still %d/%d).", online, total)
			color = 0xF0C984
		case diff < 0:
			info = fmt.Sprintf("`%d` shard(s) went offline (now %d/%d).", -diff, online, total)
			color = 0xE07264
		}

		embed := statusEmbed{
			Author:      statusEmbedAuthor{IconURL: yagpdbIconURL, Name: "YAGPDB Status"},
			Color:       color,
			Description: info,
			Timestamp:   time.Now().Format(time.RFC3339),
		}
		sendStatusEmbed(embed)

		prevOfflineCount = offlineCount
	}
}

func sendStatusEmbed(embed statusEmbed) {
	content, err := json.Marshal(webhookPostRequest{[]statusEmbed{embed}})
	if err != nil {
		log.Fatal("error marshaling embed into bytes: ", err)
	}

	_, err = client.Post(webhookURL, "application/json", bytes.NewBuffer(content))
	if err != nil {
		log.Println("error posting to webhook: ", err)
	}
}
