package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":7319")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("start server at %s\n", listener.Addr())

	defer func() {
		err := listener.Close()
		if err != nil {
			log.Printf("server stoped with error %v\n", err)
		} else {
			log.Println("server stoped")
		}
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("error occuried when accept connection: %v\n", err)
		}

		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Printf("error close connection: %v\n", err)
		}
	}()

	for {
		time.Sleep(time.Second)
		_, err := io.WriteString(conn, time.Now().Format("15:04:05\n\r"))
		if err != nil {
			log.Printf("error send time to client: %v", err)
			break
		} else {
			log.Printf("send time to client")
		}
	}
}
