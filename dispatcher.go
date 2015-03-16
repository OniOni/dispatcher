package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"encoding/json"
	"github.com/OniOni/dispatcher/store"
)

var subscribers *store.Store;

func push(subscriber, mess string) {
	fmt.Println(subscriber, mess)
}

func subscribe(id, subscriber string) {
	err := subscribers.AddSubsriber(id, subscriber)

	if err != nil {
		fmt.Println("Error: ", err)
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
			subs, _ := subscribers.GetSubscribers(parts[2])
			b, _ := json.Marshal(subs)
			conn.Write(b)

		case "msub":
			for _, val := range parts[2:] {
				subscribe(val, parts[1])
			}
			subs, _ := subscribers.GetSubscribers(parts[2])
			b, _ := json.Marshal(subs)
			conn.Write(b)

		case "push":
			subs, _ := subscribers.GetSubscribers(parts[1])
			for _, val := range subs  {
				go push(val, parts[2])
			}

		}
		conn.Write([]uint8("\n"))
		conn.Close()
	}
}

func main() {
	var err error;
	subscribers, err = store.NewStore()

	if err != nil {
		panic(err)
	}

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
