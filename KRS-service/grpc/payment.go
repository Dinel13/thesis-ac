package grpc

import (
	"context"
	"fmt"

	"github.com/dinel13/thesis-ac/krs/proto"
)

// var urlPay = os.Getenv("URL_PAYMENT")

// type IsPay struct {
// 	IsPay bool
// 	Err   error
// }

// type IsAuth struct {
// 	IsAuth bool
// 	Err    error
// }

// func verifyPayment(userId int) IsPay {
// 	conn, err := grpc.Dial(fmt.Sprintf("%s:9092", urlPay), grpc.WithInsecure())
// 	if err != nil {
// 		fmt.Println(err)
// 		return IsPay{
// 			IsPay: false,
// 			Err:   err,
// 		}
// 	}

// 	defer conn.Close()
// 	c := proto.NewPaymentServiceClient(conn)

// 	r, err := c.Verify(context.Background(), &proto.VerifyPaymentRequest{
// 		IdMahasiswa: int32(userId),
// 	})

// 	if err != nil {
// 		fmt.Println(err)
// 		return IsPay{
// 			IsPay: false,
// 			Err:   err,
// 		}
// 	}
// 	return IsPay{
// 		IsPay: r.IsPay,
// 		Err:   nil,
// 	}
// }

func VerifyPayment(ctx context.Context, c proto.PaymentServiceClient, userId int) (bool, error) {
	r, err := c.Verify(ctx, &proto.VerifyPaymentRequest{
		IdMahasiswa: int32(userId),
	})

	if err != nil {
		fmt.Println(err)
		return false, err

	}
	return r.IsPay, nil

}

// func pay(userId int, jummlah float64, metode string) bool {
// 	conn, err := grpc.Dial(fmt.Sprintf("%s:9092", urlPay), grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("did not connect: %v", err)
// 	}

// 	defer conn.Close()
// 	c := proto.NewPaymentServiceClient(conn)

// 	r, err := c.Create(context.Background(), &proto.CreatePaymentRequest{
// 		IdMahasiswa: int32(1),
// 		Jumlah:      float32(jummlah),
// 		Metode:      metode,
// 	},
// 	)

// 	if err != nil {
// 		log.Fatalf("could not greet: %v", err)
// 	}

// 	return r.IsPay

// }
