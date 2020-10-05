package mongodb

import (
	"fmt"
)

// Connection repesentation
type Connection struct {
	IsConnected bool
}

// Connect is trying to connect to mongo server. Return error on failer.
func (c Connection) Connect() error {

	if c.IsConnected {
		fmt.Println("Already connected")
	} else {
		fmt.Println("Connecting...")
		fmt.Println("Connected.")
	}

	return nil
}

// Disconnect is trying to disconnect from mongo server. Return error on failer.
func (c Connection) Disconnect() error {

	if c.IsConnected {
		fmt.Println("Disconnected.")
	} else {
		fmt.Println("Already disconnected.")
	}

	return nil
}
