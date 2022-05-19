package sqs

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func DeleteMessage(sess *session.Session, queueURL *string, messageHandle *string) error {
	// snippet-start:[sqs.go.delete_message.call]
	svc := sqs.New(sess)

	_, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      queueURL,
		ReceiptHandle: messageHandle,
	})
	// snippet-end:[sqs.go.delete_message.call]
	if err != nil {
		return err
	}

	return nil
}
