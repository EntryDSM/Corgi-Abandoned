package main

import (
	"github.com/EntryDSM/Corgi/initialize"
	"github.com/EntryDSM/Corgi/loop"
)

func main() {
	initialize.Initialize()
	loop.StartLoop()
}
