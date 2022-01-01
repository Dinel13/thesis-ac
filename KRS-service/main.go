package main

import "github.com/dinel13/thesis-ac/krs/app"

func main() {
	app.StartRestServer()
	// go app.StartGRPCServer()
	// Calling Sleep method
	// time.Sleep(8 * time.Second)

	// Printed after sleep is over
	// fmt.Println("Sleep Over.....")
	// grpc.Create()
	// grpc.Read()
	// grpc.Update()
	// grpc.Delete()
	// app.StartGRPCServer()
}
