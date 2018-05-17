package main

import (
	"log"
	"strings"

	"github.com/atotto/clipboard"
)

func main() {
	clip, err := clipboard.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	rep := strings.TrimSpace(clip)
	rep = strings.Trim(rep, "\"'")

	if err := clipboard.WriteAll(rep); err != nil {
		log.Fatalln(err)
	}
}
