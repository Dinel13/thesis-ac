package grpc

import (
	"context"
	"log"

	"github.com/dinel13/thesis-ac/krs/proto"
	"google.golang.org/grpc"
)

func Create() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := proto.NewKrsServiceClient(conn)

	r, err := c.Create(context.Background(), &proto.Krs{
		Id:          int32(1),
		Name:        "Go",
		Description: "Go is a programming language",
	})
	if err != nil {
		log.Fatalf("could not get krs: %v", err)
	}

	log.Printf("Krs: %s", r.Krss.Name)
	log.Printf("Krs: %s", r.Krss.Description)
}

func Read() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := proto.NewKrsServiceClient(conn)

	r, err := c.Read(context.Background(), &proto.KrsRequest{Id: int32(1)})
	if err != nil {
		log.Fatalf("could not get krs: %v", err)
	}

	log.Printf("Krs: %s", r.Krss.Name)
	log.Printf("Krs: %s", r.Krss.Description)
}
