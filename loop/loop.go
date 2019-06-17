package loop

import (
	"fmt"
	"time"
)

// import (
// 	"log"

// 	"github.com/EntryDSM/Corgi/config"
// 	"github.com/EntryDSM/Corgi/fetch"
// )

func StartLoop() {
	mainLoop()
}

func mainLoop() {
	for {
		fmt.Println("Waiting...")
		time.Sleep(time.Second)
	}
}
