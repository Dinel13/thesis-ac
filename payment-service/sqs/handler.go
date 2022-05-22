package sqs

import (
	"context"
	"fmt"
	"strconv"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/dinel13/thesis-ac/payment/domain"
)

func NewSQSHandler(s domain.PaymentRepository, sess *session.Session, qUrlPay, qUrlKrs *string, waitTime *int64) domain.PaymentSQSHandler {
	msgGrup := "payment"

	return sqsHandler{
		repo:     s,
		sess:     sess,
		qUrlPay:  qUrlPay,
		qUrlKrs:  qUrlKrs,
		waitTime: waitTime,
		msgGrup:  &msgGrup,
	}

}

type sqsHandler struct {
	sess     *session.Session
	qUrlPay  *string
	qUrlKrs  *string
	msgGrup  *string
	waitTime *int64
	repo     domain.PaymentRepository
}

func (h sqsHandler) GetLPMessages() (*sqs.ReceiveMessageOutput, error) {
	svc := sqs.New(h.sess)

	result, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl: h.qUrlPay,
		AttributeNames: aws.StringSlice([]string{
			"SentTimestamp",
		}),
		MaxNumberOfMessages: aws.Int64(1),
		MessageAttributeNames: aws.StringSlice([]string{
			"All",
		}),
		WaitTimeSeconds: h.waitTime,
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (h sqsHandler) SendMsg(msgId *string, msg bool) (*string, error) {
	svc := sqs.New(h.sess)

	// marshal data to json
	// jsonData, err := json.Marshal(data)
	// if err != nil {
	// 	return nil, err
	// }

	// jsonDataStr := string(jsonData)
	msgStr := strconv.FormatBool(msg)

	res, err := svc.SendMessage(&sqs.SendMessageInput{
		MessageGroupId: h.msgGrup,
		DelaySeconds:   aws.Int64(0),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"ID": {
				DataType:    aws.String("String"),
				StringValue: msgId,
			},
		},
		MessageBody: &msgStr,
		QueueUrl:    h.qUrlKrs,
	})
	if err != nil {
		return nil, err
	}

	return res.MessageId, nil
}

func (h sqsHandler) DeleteMessage(messageHandle *string) error {
	fmt.Println("deleting message... ")
	svc := sqs.New(h.sess)

	_, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      h.qUrlPay,
		ReceiptHandle: messageHandle,
	})
	if err != nil {
		return err
	}
	return nil
}

func (h sqsHandler) WaitMsgSqs(wg sync.WaitGroup) {
	for {
		msgResult, err := h.GetLPMessages()
		if err != nil {
			fmt.Println("Got an error receiving messages:")
			fmt.Println(err)
			continue
		}

		// cek if message is empty
		if len(msgResult.Messages) == 0 {
			fmt.Println("No messages to process.")
			continue
		}

		var idReq = msgResult.Messages[0].MessageAttributes["ID"].StringValue
		idMhs, err := strconv.Atoi(*msgResult.Messages[0].Body)
		if err != nil {
			fmt.Println(err)
			continue
		}

		payment, err := h.repo.Verify(context.Background(), idMhs)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Payment: ", payment)

		_, err = h.SendMsg(idReq, payment.IsPay)
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = h.DeleteMessage(msgResult.Messages[0].ReceiptHandle)
		if err != nil {
			fmt.Println("Got an error deleting the message:")
			fmt.Println(err)
			continue
		}

		fmt.Println("finished request " + *idReq)
	}

	wg.Done()
}
