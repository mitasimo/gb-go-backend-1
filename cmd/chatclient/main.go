package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("usage: chatclient username")
		return
	}

	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	n, err := io.WriteString(conn, os.Args[1]+"\n")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("send %d bytes\n", n)

	go func() {
		io.Copy(os.Stdout, conn)
	}()
	io.Copy(conn, os.Stdin) // until you send ^Z
	fmt.Printf("%s: exit", conn.LocalAddr())
}
