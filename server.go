package main

import (
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

		for scanner.Scan() {
			command := scanner.Text()
			log.Println("Read", command)
			response := parse(command)
			log.Println("Response", response)
			writer := bufio.NewWriter(conn)
			writer.WriteString(response)
			writer.Flush()
		}


}