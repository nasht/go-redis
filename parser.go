package main

import (
	"log"
	"strings"
	"redis/service"
)



func parse(commands []string) string {
	store := service.NewStorageService()
	log.Printf("GOT %v \n" , commands)
	command, params := shift(commands)

	switch strings.ToUpper(command) {
	case "PING":
		return "+PONG\r\n"
	case "GET":

		for _, element := range params {
			if !(strings.HasPrefix(element, "*") || strings.HasPrefix(element, "$")) {
				return store.Get(element)
			}
		}
		return "-Incomplete GET command. Supply the key"

	default:
		return "-Unknown command " + command + "\r\n"
	}
}


func pop[T comparable](a []T) (T, []T) {
	return a[len(a)-1], a[:len(a)-1]
}

func shift[T comparable](a []T) (T, []T) {
	return a[0], a[1:]
}