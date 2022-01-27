package main

import (
	"time"

	"github.com/dinel13/thesis-ac/krs/app"
)

func main() {
	go app.StartRestServer()
	go app.StartGRPCServer()
	time.Sleep(10 * time.Minute)
}
