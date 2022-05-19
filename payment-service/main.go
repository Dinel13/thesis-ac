package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/dinel13/thesis-ac/payment/domain"
	"github.com/dinel13/thesis-ac/payment/repository"
	"github.com/dinel13/thesis-ac/payment/sqs"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"

	mygrpc "github.com/dinel13/thesis-ac/payment/grpc"

	"github.com/dinel13/thesis-ac/payment/proto"
	"github.com/dinel13/thesis-ac/payment/rest"
	"github.com/dinel13/thesis-ac/payment/service"
)

var ipRedis = os.Getenv("IP_REDIS")

// StartRestServer starts the REST server
func StartRestServer() {
	port := ":8082"
	fmt.Printf("Staring REST server on port %s\n", port)

	dbClient, paymentRepo := startRepoRedis()
	defer dbClient.Close()

	cs := rest.NewPaymentRestHandlers(service.NewPaymentService(paymentRepo))

	srv := &http.Server{
		Addr:    port,
		Handler: rest.Routes(cs),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

// StartGRPCServer starts the gRPC server
func StartGRPCServer() {
	port := ":9092"
	fmt.Printf("Staring gRPC server on port %s\n", port)

	dbClient, paymentRepo := startRepoRedis()
	defer dbClient.Close()
	cs := mygrpc.NewGrpcHandler(service.NewPaymentService(paymentRepo))

	// create gRPC server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterPaymentServiceServer(s, cs)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func startRepoRedis() (*redis.Client, domain.PaymentRepository) {
	dbClient := connectRedis()
	crDb := repository.NewPaymentRepoRedisImpl(dbClient)
	return dbClient, crDb
}

func connectRedis() *redis.Client {
	// connect to redis
	log.Println("Connecting to redis...")
	client := redis.NewClient(&redis.Options{
		Addr: ipRedis + ":6379",
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Cannot connect to redis! Dying...")
	}
	log.Println("Connected to redis!")
	return client
}

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	queueURL := "https://sqs.ap-southeast-1.amazonaws.com/511415859101/thesis.fifo"
	waitTime := int64(20)

	// err := sqs.SendMsg(sess, &msgGrup, &queueURL)
	// if err != nil {
	// 	fmt.Println("Got an error sending the message:")
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("Sent message to queue ")
	// println("sleep")
	// time.Sleep(2 * time.Second)

	// var timeout = int64(10)
	//msgResult, err := sqs.GetMessages(sess, &queueURL, &timeout)
	//if err != nil {
	//	fmt.Println("Got an error receiving messages:")
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println("Message ID:     " + *msgResult.Messages[0].MessageId)
	//fmt.Println("Message ID:     " + *msgResult.Messages[0].Body)
	//
	//// snippet-start:[sqs.go.receive_messages.print_handle]
	//fmt.Println("Message Handle: " + *msgResult.Messages[0].ReceiptHandle)

	//err = sqs.DeleteMessage(sess, &queueURL, msgResult.Messages[0].ReceiptHandle)
	//if err != nil {
	//	fmt.Println("Got an error deleting the message:")
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println("Deleted message from queue with URL " + queueURL)

	//if ipRedis == "" {
	//	log.Fatal("IP_REDIS environment variable not set! Dying...")
	//}
	//log.Println("use IP_REDIS: ", ipRedis)
	//go StartRestServer()
	//go StartGRPCServer()
	go sqs.WaitMsgSqs(sess, &queueURL, &waitTime)
	time.Sleep(113880 * time.Hour)

}
