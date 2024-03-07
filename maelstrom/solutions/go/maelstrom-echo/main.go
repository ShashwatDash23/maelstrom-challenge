package main

import (
	"encoding/json"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {

	// create a new node
	n := maelstrom.NewNode()

	// create a handler callback function
	n.Handle("echo", func(msg maelstrom.Message) error {

		// value can be anything
		body := map[string]any{}

		// unmarshal the msg into the body map
		err := json.Unmarshal(msg.Body, &body)

		if err != nil {
			return err
		}

		// Update the message type.
		body["type"] = "echo_ok"

		// Echo the original message back with the updated message type.
		return n.Reply(msg, body)
	})

	// Execute the node's message loop
	err := n.Run()
	if err != nil {
		log.Fatal(err)
	}
}
