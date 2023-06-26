package models

import (
	"log"
	"io"
	"bufio"
	"strconv"
	"strings"
	"redis/utils/stringutils"
	"redis/utils/array"
)

type Command struct {
	Name string
	Params []string
}


func NewCommand(reader io.Reader ) Command {
		scanner := bufio.NewScanner(reader)
		scanner.Scan()
		text := scanner.Text()
		log.Println("Got input: ",  text)
		_, countText := stringutils.Shift(text)
		count, err := strconv.Atoi(countText)
		if err != nil {
			log.Fatal("Can't parse parameters. Input:", text)
		}

		log.Println("Expected Params: ", count)
		params := []string{}
		for i := 0; i < count; i++ {
			scanner.Scan()
			numChars := scanner.Text() // This should just show how many characters to expect
			log.Println("NumChars:", numChars)
			scanner.Scan()
			inputVal := scanner.Text()
			log.Println("Input:", inputVal)
			params = append(params,inputVal)
			log.Println("Added: ",  inputVal)
		}

		commandName, commandValues := array.Shift(params)
		return Command { Name: strings.ToUpper(commandName), Params: commandValues }
}
