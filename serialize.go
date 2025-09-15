package main

import (
	"fmt"
	messageTypes "github.com/niljimeno/gopher/types"
	"strconv"
	"strings"
)

type Message struct {
	Content      string
	Host         string
	Port         int
	Subdirectory string
	Type         byte
}

var malformedMessage Message = Message{
	Type:    messageTypes.Error,
	Content: "ERR: malformed message",
}

func Serialize(m string) Message {
	messageParts := strings.Split(tail(m), "\t")
	if len(messageParts) < 4 {
		return malformedMessage
	}

	port, err := strconv.Atoi(messageParts[3])
	if err != nil {
		return malformedMessage
	}

	return Message{
		Type:         m[0],
		Content:      messageParts[0],
		Subdirectory: messageParts[1],
		Host:         messageParts[2],
		Port:         port,
	}
}

func tail(s string) string {
	return s[1:]
}

func (m *Message) Print() {
	fmt.Printf("%c :: %s :: %s :: %s :: %d\n",
		m.Type,
		m.Content,
		m.Subdirectory,
		m.Host,
		m.Port)
}
