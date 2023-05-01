package main

import (
	"github.com/javier-tw/learning-go/cmd/challenge-4/server"
	"log"
)

func main() {
	if err := server.SetupEngine().Run(); err != nil {
		log.Fatal(err)
	}
}
