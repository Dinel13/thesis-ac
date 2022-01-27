package grpc

import (
	"context"
	"log"

	"github.com/dinel13/thesis-ac/payment/proto"
	"google.golang.org/grpc"
)

func Create() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := proto.NewPaymentServiceClient(conn)

	r := &proto.CreatePaymentRequest{
		IdMahasiswa: 1,
		Jumlah:      100,
		Metode:      "Transfer",
	}

	res, err := c.Create(context.Background(), r)

	if err != nil {
		log.Fatalf("Error while calling Create: %v", err)
	}

	if res.IsPay {
		log.Println("Payment success")
	} else {
		log.Println("Payment failed")
	}
}

func Read() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := proto.NewPaymentServiceClient(conn)

	r := &proto.VerifyPaymentRequest{IdMahasiswa: 1}

	res, err := c.Verify(context.Background(), r)

	if err != nil {
		log.Println(err)
	}

	if res.IsPay {
		log.Println("Payment success")
	} else {
		log.Println("Payment failed")
	}

}
