package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"time"
)

const serverAddr = "0.0.0.0:7319"

func main() {

	var clients = make(map[string]net.Conn)

	listener, err := net.Listen("tcp", serverAddr)
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

	go sendMessageToClients(clients)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("error occuried when accept connection: %v\n", err)
			return
		}

		// добавить соединение в мапу
		clients[connName(conn)] = conn

		go handleConnection(conn, clients)
	}
}

func handleConnection(conn net.Conn, clients map[string]net.Conn) {
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
			// удалить соединение из мапы, так как клиент разорвал соединение
			delete(clients, connName(conn))
			break
		}
	}
}

// читает строки из стандартного потока вводи и отправляет их всем клинетам
func sendMessageToClients(clients map[string]net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		for addr, conn := range clients {
			_, err := io.WriteString(conn, text+"\n")
			if err != nil {
				log.Printf("error send text to client %s: %v", addr, err)
			}
		}
	}
}

func connName(conn net.Conn) string {
	return conn.RemoteAddr().String()
}
