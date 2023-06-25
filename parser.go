package main

import (
	"strings"
)
func parse(command string) string {
	switch strings.ToUpper(command) {
	case "PING":
		return "+PONG\r\n"
	default:
		return ""
	}
}