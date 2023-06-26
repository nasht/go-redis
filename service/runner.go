package service

import (
	"log"
	"redis/models"
)

func Run(command models.Command) string {
	switch command.Name {
	case "PING":
		 return PingRunner{}.Run(command)
	case "SET":
		return SetRunner{}.Run(command)
		case "GET":
		 return GetRunner{}.Run(command)

	default:
		return DefaultRunner{}.Run(command)
	}
}

var dataStore = make(map[string]string)

type Runner interface {
	Run(models.Command) string
}


type PingRunner struct {}
func (r PingRunner) Run(models.Command) string {
	return "+PONG"
}

type SetRunner struct {}
func (r SetRunner) Run(c models.Command) string {
	log.Printf("Got: %v ", c)
	if len(c.Params) < 2 {
		return "-ERR wrong number of arguments for 'set' command"
	}
	dataStore[c.Params[0]] = c.Params[1]
	log.Printf("DataStore: %v ", dataStore)
	return "+OK"
}


type GetRunner struct {}
func (r GetRunner) Run(c models.Command) string {
	if len(c.Params) < 1 {
			return "-ERR wrong number of arguments for 'set' command"
	}
	if dataStore[c.Params[0]] == "" {
		return "-ERR no value for " + c.Params[0]
	}
	return "+" + string(dataStore[c.Params[0]])
}

type DefaultRunner struct {}
func (r DefaultRunner) Run(c models.Command) string {
	return "-Unknown Command " + c.Name
}

