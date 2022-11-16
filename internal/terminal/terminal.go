package terminal

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Read() string {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSuffix(str, "\n")
}

// func ReadInput(prompt, option string) {
// 	str := Read()
// 	c := *config
// 	c.SetOption(option, str)
// }
