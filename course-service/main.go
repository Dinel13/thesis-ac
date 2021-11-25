package main

import (
	"fmt"
	"time"

	"github.com/dinel13/thesis-ac/course/app"
	"github.com/dinel13/thesis-ac/course/grpc"
)

func main() {
	go app.StartRestServer()
	go app.StartGRPCServer()
	// Calling Sleep method
	time.Sleep(8 * time.Second)

	// Printed after sleep is over
	fmt.Println("Sleep Over.....")
	grpc.Main()
}
