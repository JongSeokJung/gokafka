package main

import (
	"fmt"

	"github.com/JongSeokJung/gokafka/internal/log"
)

func main() {
	myMessage := log.Message{
		Offset: 123,
		Timestamp: 456,
		Payload: []byte("Hello, World!"),
	}

	encoded := myMessage.Encode()
	fmt.Printf("Encoded Message: %x\n", encoded)

	decoded := log.Decode(encoded)
	fmt.Printf("Decoded Message: Offset=%d, Timestamp=%d, Payload=%s\n", decoded.Offset, decoded.Timestamp, string(decoded.Payload))

}
