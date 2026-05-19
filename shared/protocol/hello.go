package protocol

type Hello struct {
	NodeID string `json:"node_id"`
	Name   string `json:"name"`
}
