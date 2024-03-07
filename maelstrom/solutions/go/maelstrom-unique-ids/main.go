package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func generateCustomId() string {

	// get current timestamp for random seed
	timestamp := time.Now().UnixNano()
	// get a random seed
	source := rand.NewSource(timestamp)
	randSeed := rand.New(source)
	randNo := randSeed.Intn(10000)

	// unique id is a combination of timestamp and random number generated
	uniqueId := fmt.Sprintf("%d-%d", timestamp, randNo)

	return uniqueId
}

func generateUUID() string {
	uniqueId := uuid.NewString()

	return uniqueId
}

func main() {

	// create a new node
	n := maelstrom.NewNode()

	// create a handler callback function
	n.Handle("generate", func(msg maelstrom.Message) error {

		// value can be anything
		body := map[string]any{}

		// unmarshal the msg into the body map
		err := json.Unmarshal(msg.Body, &body)

		if err != nil {
			return err
		}

		// update the body's type and generate a new unique id
		body["type"] = "generate_ok"
		body["id"] = generateCustomId()

		// return the original msg back with updated body
		return n.Reply(msg, body)
	})

	err := n.Run()
	if err != nil {
		log.Fatal(err)
	}
}
