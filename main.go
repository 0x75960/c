package main

import (
	"log"
	"strings"

	"github.com/atotto/clipboard"
)

func cleanup(input string) string {
	rep := strings.TrimSpace(input)
	rep = strings.Trim(rep, "\"'")
	return rep
}

func main() {
	clip, err := clipboard.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	rep := cleanup(clip)

	if err := clipboard.WriteAll(rep); err != nil {
		log.Fatalln(err)
	}
}
