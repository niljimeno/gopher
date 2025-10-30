package tcp

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func Dial(url, route string) []Message {
	return labrat()

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
		return nil
	}

	conn.Close()
	return messages
}

func labrat() []Message {
	time.Sleep(time.Second * 0)

	rat, err := os.OpenFile("example", os.O_RDONLY, 0440)
	if err != nil {
		return nil
	}

	scanner := bufio.NewScanner(rat)
	var messages []Message
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		messages = append(messages, Serialize(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		return nil
	}

	return messages
}
