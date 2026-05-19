package protocol

import (
	"encoding/json"
	"fmt"
	"io"
)

type Message struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

func EncodeMessage(
	t string,
	v any,
) ([]byte, error) {
	payload, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	msg := Message{
		Type:    t,
		Payload: payload,
	}

	return json.Marshal(msg)
}

func DecodeMessage[T any](
	r io.Reader,
	expectedType string,
) (*T, error) {

	var msg Message

	if err := json.NewDecoder(r).Decode(&msg); err != nil {
		return nil, err
	}

	if msg.Type != expectedType {
		return nil, fmt.Errorf("invalid message type")
	}

	var payload T

	if err := json.Unmarshal(msg.Payload, &payload); err != nil {
		return nil, err
	}

	return &payload, nil
}
