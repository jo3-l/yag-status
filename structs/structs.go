package structs

import "time"

// BotStatus is the overall status of the bot.
type BotStatus struct {
	HostStatuses []*HostStatus `json:"host_statuses"`
	NumNodes     int           `json:"num_nodes"`
	TotalShards  int           `json:"total_shards"`

	UnavailableGuilds int   `json:"unavailable_guilds"`
	OfflineShards     []int `json:"offline_shards"`

	EventsPerSecondAverage float64 `json:"events_per_second_average"`
	EventsPerSecondMin     float64 `json:"events_per_second_min"`
	EventsPerSecondMax     float64 `json:"events_per_second_max"`

	UptimeMax time.Duration `json:"uptime_max"`
	UptimeMin time.Duration `json:"uptime_min"`
}

// HostStatus is the status of a host.
type HostStatus struct {
	Name string

	EventsPerSecond float64
	TotalEvents     int64

	Nodes []*NodeStatus
}

// NodeStatus is the status of a node.
type NodeStatus struct {
	ID     string         `json:"id"`
	Shards []*ShardStatus `json:"shards"`
	Host   string         `json:"host"`
	Uptime time.Duration  `json:"uptime"`
}

// ShardStatus is the status of a shard.
type ShardStatus struct {
	ShardID         int     `json:"shard_id"`
	TotalEvents     int64   `json:"total_events"`
	EventsPerSecond float64 `json:"events_per_second"`

	ConnStatus GatewayStatus `json:"conn_status"`

	LastHeartbeatSend time.Time `json:"last_heartbeat_send"`
	LastHeartbeatAck  time.Time `json:"last_heartbeat_ack"`

	NumGuilds         int
	UnavailableGuilds int
}

// GatewayStatus is the status of the Discord connection.
type GatewayStatus int

// Possible gateway statuses.
const (
	GatewayStatusDisconnected GatewayStatus = iota
	GatewayStatusConnecting
	GatewayStatusIdentifying
	GatewayStatusResuming
	GatewayStatusReady
)

func (gs GatewayStatus) String() string {
	switch gs {
	case GatewayStatusDisconnected:
		return "Disconnected"
	case GatewayStatusConnecting:
		return "Connecting"
	case GatewayStatusIdentifying:
		return "Identifying"
	case GatewayStatusResuming:
		return "Resuming"
	case GatewayStatusReady:
		return "Ready"
	}

	return "??"
}
