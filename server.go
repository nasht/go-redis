package main

import (
	"strings"
	"bufio"
	"net"
	"log"
)


func startServer() {
	log.Println("Starting server...")
	defer	log.Println("Finished")

	listener, err := net.Listen("tcp", ":6379")
	if err != nil {
		log.Fatal("Error:", err, " while opening connection")
	}

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal("Error:", err, " while setting up connection")
		}
		go handleConnection(connection)
	}


}


func handleConnection(conn net.Conn) {

		scanner := bufio.NewScanner(conn)
		var command_set []string

		for scanner.Scan() {
			count := atoi scanner.Text()
			command_set = append(command_set, command)
			log.Println("Got Command", command)
			if strings.HasSuffix(command, "\r\n") {
				break
			}
		}
		response := parse(command_set)
		log.Println("Response", response)
		writer := bufio.NewWriter(conn)
		writer.WriteString(response)
		writer.Flush()


}