package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:7319")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	err = scanner.Err()
	if err != nil && err != io.EOF {
		log.Println(err)
	}
}
