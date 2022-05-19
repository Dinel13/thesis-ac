package sqs

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func GetLPMessages(sess *session.Session, queueURL *string, waitTime *int64) (*sqs.ReceiveMessageOutput, error) {
	// var msgs []*sqs.Message
	svc := sqs.New(sess)

	// snippet-start:[sqs.go.receive_lp_message.call]
	result, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl: queueURL,
		AttributeNames: aws.StringSlice([]string{
			"SentTimestamp",
		}),
		MaxNumberOfMessages: aws.Int64(1),
		MessageAttributeNames: aws.StringSlice([]string{
			"All",
		}),
		WaitTimeSeconds: waitTime,
	})
	// snippet-end:[sqs.go.receive_lp_message.call]
	if err != nil {
		return nil, err
	}

	return result, nil
}

func WaitMsgSqs(sess *session.Session, queueURL *string, waitTime *int64) {
	for {
		msgResult, err := GetLPMessages(sess, queueURL, waitTime)
		if err != nil {
			fmt.Println("Got an error receiving messages:")
			fmt.Println(err)
			return
		}

		// cek if message is empty
		if len(msgResult.Messages) == 0 {
			continue
		}

		fmt.Println("Message ID:     " + *msgResult.Messages[0].MessageId)
		fmt.Println("Message ID:     " + *msgResult.Messages[0].Body)

		// snippet-start:[sqs.go.receive_messages.print_handle]
		fmt.Println("Message Handle: " + *msgResult.Messages[0].ReceiptHandle)

		err = DeleteMessage(sess, queueURL, msgResult.Messages[0].ReceiptHandle)
		if err != nil {
			fmt.Println("Got an error deleting the message:")
			fmt.Println(err)
			return
		}

		fmt.Println("Deleted message from queue with URL " + *queueURL)
	}
}
