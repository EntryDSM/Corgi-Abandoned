package main

import (
	"log"

	"github.com/EntryDSM/Corgi/initialize"
	"github.com/EntryDSM/Corgi/loop"
)

func main() {
	defer writeFatalLog()

	initialize.Initialize()
	loop.StartLoop()
}

func writeFatalLog() {
	if err := recover(); err != nil {
		log.Fatal(err)
	}
}
