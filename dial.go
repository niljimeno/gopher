package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func Dial(url, route string) []Message {
	conn, err := net.Dial("tcp", url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(conn, route)
	scanner := bufio.NewScanner(conn)

	var messages []Message
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		messages = append(messages, Serialize(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	conn.Close()
	return messages
}
