package grpc

import (
	"context"
	"fmt"
	"log"

	"github.com/dinel13/thesis-ac/krs/proto"
	"google.golang.org/grpc"
)

func VerifyPayment(userId int) (bool, error) {
	conn, err := grpc.Dial("localhost:9092", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	defer conn.Close()
	c := proto.NewPaymentServiceClient(conn)

	r, err := c.Verify(context.Background(), &proto.VerifyPaymentRequest{
		IdMahasiswa: int32(userId),
	})

	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return r.IsPay, nil
}

func Pay(userId int, jummlah float64, metode string) bool {
	conn, err := grpc.Dial("localhost:9092", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := proto.NewPaymentServiceClient(conn)

	r, err := c.Create(context.Background(), &proto.CreatePaymentRequest{
		IdMahasiswa: int32(1),
		Jumlah:      float32(jummlah),
		Metode:      metode,
	},
	)

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	return r.IsPay

}
