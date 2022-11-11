package core

import (
	"bufio"
	"log"
	"os"
)

func ReadTerminal() string {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return str
}
