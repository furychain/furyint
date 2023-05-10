package registry

import (
	"github.com/furychain/furyint/settlement"
	"github.com/furychain/furyint/settlement/furya"
	"github.com/furychain/furyint/settlement/mock"
)

// Client represents a settlement layer client
type Client string

const (
	// Mock is a mock client for the settlement layer
	Mock Client = "mock"
	// Furya is a client for interacting with furya settlement layer
	Furya Client = "furya"
)

// A central registry for all Settlement Layer Clients
var clients = map[Client]func() settlement.LayerClient{
	Mock:      func() settlement.LayerClient { return &mock.SettlementLayerClient{} },
	Furya: func() settlement.LayerClient { return &furya.LayerClient{} },
}

// GetClient returns client identified by name.
func GetClient(client Client) settlement.LayerClient {
	f, ok := clients[client]
	if !ok {
		return nil
	}
	return f()
}

// RegisteredClients returns names of all settlement clients in registry.
func RegisteredClients() []Client {
	registered := make([]Client, 0, len(clients))
	for client := range clients {
		registered = append(registered, client)
	}
	return registered
}
