package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"encoding/json"
	"store"
)

var subscribers map[string][]string;

func contains(haystack []string, needle string) bool {
	for _, val := range haystack {
		if val == needle {
			return true;
		}
	}
	return false;
}

func push(subscriber, mess string) {
	fmt.Println(subscriber, mess)
}

func subscribe(id, subscriber string) {
	if subscribers[id] == nil {
		subscribers[id] = make([]string, 0)
	}

	if !contains(subscribers[id], subscriber) {
		subscribers[id] = append(subscribers[id], subscriber)
	}
}

func respond(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {

		content, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		content = content[:len(content) - 2]
		parts := strings.Split(content, " ")
		cmd := parts[0]
		switch cmd {

		case "sub":
			subscribe(parts[2], parts[1])
			b, _ := json.Marshal(subscribers)
			conn.Write(b)

		case "msub":
			for _, val := range parts[2:] {
				subscribe(val, parts[1])
			}
			b, _ := json.Marshal(subscribers)
			conn.Write(b)

		case "push":
			for _, val := range subscribers[parts[1]] {
				go push(val, parts[2])
			}

		}
		conn.Write([]uint8("\n"))
		conn.Close()
	}
}

func main() {

	if err != nil {
		panic(err)
	}

	subscribers = make(map[string][]string)

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			return
		}

		fmt.Println("Got connection.")
		go respond(conn)
	}
}
