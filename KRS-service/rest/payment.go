package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type payment struct {
	IsPay bool `json:"isPay"`
}

type resDataPayment struct {
	Payment payment `json:"payment"`
}

func VerifyPayment(userId int) (bool, error) {
	var client = &http.Client{}

	ip := os.Getenv("IP")

	request, err := http.NewRequest("GET", fmt.Sprintf("http://%s:8082/verify/%d", ip, userId), nil)
	if err != nil {
		return false, err
	}
	response, err := client.Do(request)
	if err != nil {
		return false, err
	}

	defer response.Body.Close()

	var dataPayment resDataPayment
	err = json.NewDecoder(response.Body).Decode(&dataPayment)
	if err != nil {
		return false, err
	}

	return dataPayment.Payment.IsPay, nil
}

// func Pay(userId int, jummlah float64, metode string) bool {
// 	conn, err := grpc.Dial("localhost:9092", grpc.WithInsecure())
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
