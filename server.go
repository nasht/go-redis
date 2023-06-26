package main

import (
	"bufio"
	"redis/models"
	"net"
	"log"
	"redis/service"
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

		command := models.NewCommand(conn)
		log.Printf("Command: %v ", command)
		go func ()  {
			writer := bufio.NewWriter(conn)
			response := service.Run(command)
			log.Printf("Response: %v ", response)
			writer.WriteString(response + "\r\n")
			writer.Flush()
		}()

}