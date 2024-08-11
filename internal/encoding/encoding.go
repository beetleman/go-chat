package encoding

import (
	"encoding/json"
	"fmt"
	"log"
)

type Message struct {
	Text string
	User string
}

func (msg Message) String() string {
	return fmt.Sprintf("%s: %s", msg.User, msg.Text)
}

// Encode `Message` into bytes
func (msg Message) Encode() []byte {
	encoded, err := json.Marshal(msg)
	if err != nil {
		log.Printf("%v", err)
	}
	return append(encoded, []byte("\n")...)
}

// Decode `Message` from bytes
func Decode(data []byte) Message {
	msg := Message{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		log.Printf("%v", err)
	}
	return msg
}
