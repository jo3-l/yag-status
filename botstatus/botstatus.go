package botstatus

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/jo3-l/yag-status/structs"
)

const statusEndpoint = "https://yagpdb.xyz/status.json"

var client http.Client = http.Client{Timeout: 5 * time.Second}

// GetStatus returns the bot status.
func GetStatus() structs.BotStatus {
	resp, err := client.Get(statusEndpoint)
	if err != nil {
		log.Fatal("failed fetching data:", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("failed reading response:", err)
	}

	var status structs.BotStatus
	err = json.Unmarshal(body, &status)
	if err != nil {
		log.Fatal("failed reading JSON:", err)
	}

	return status
}
