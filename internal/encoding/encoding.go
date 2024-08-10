package encoding

import (
	"encoding/json"
	"log"
)

type Message struct {
	Text string
}

// Encode `Message` into bytes
func (msg Message) Encode() []byte {
	encoded, err := json.Marshal(msg)
	if err != nil {
		log.Printf("%v", err)
	}
	return encoded
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
