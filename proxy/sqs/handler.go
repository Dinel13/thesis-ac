package sqs

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/dinel13/thesis-ac/test/domain"
)

func NewSQSHandler(sess *session.Session, in, out *string, waitTime *int64) domain.PaymentSQSHandler {
	msgGrup := "payment"

	return sqsHandler{
		sess:     sess,
		in:       in,
		out:      out,
		waitTime: waitTime,
		msgGrup:  &msgGrup,
	}

}

type sqsHandler struct {
	sess     *session.Session
	in       *string
	out      *string
	msgGrup  *string
	waitTime *int64
}

func (h sqsHandler) GetLPMessages() (*sqs.ReceiveMessageOutput, error) {
	svc := sqs.New(h.sess)

	result, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl: h.in,
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

func (h sqsHandler) SendMsg(msgId, body *string) (*string, error) {
	svc := sqs.New(h.sess)

	res, err := svc.SendMessage(&sqs.SendMessageInput{
		MessageGroupId: h.msgGrup,
		DelaySeconds:   aws.Int64(0),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"ID": {
				DataType:    aws.String("String"),
				StringValue: msgId,
			},
		},
		MessageBody: body,
		QueueUrl:    h.out,
	})
	if err != nil {
		return nil, err
	}

	fmt.Printf("send id %s with body %s", *msgId, *body)
	return res.MessageId, nil
}

func (h sqsHandler) DeleteMessage(messageHandle *string) error {
	svc := sqs.New(h.sess)

	_, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      h.in,
		ReceiptHandle: messageHandle,
	})
	if err != nil {
		return err
	}
	fmt.Println("Deleted message from queue with URL " + *h.in)

	return nil
}

func (h sqsHandler) WaitMsgSqs(responChan chan domain.ResponMsgWaitChan, id *string) {
	fmt.Printf("wait for %s", *id)
	for {
		msgResult, err := h.GetLPMessages()
		if err != nil {
			fmt.Println("Got an error receiving messages:")
			responChan <- domain.ResponMsgWaitChan{
				IsPay: false,
				Err:   err,
			}
			return
		}

		// cek if message is empty
		if len(msgResult.Messages) == 0 {
			continue
		}

		var idReq = msgResult.Messages[0].MessageAttributes["ID"].StringValue
		respon := *msgResult.Messages[0].Body

		fmt.Println(*idReq, respon, *id)
		if *id != *idReq {
			fmt.Println("tidak sama")
			continue
		}

		err = h.DeleteMessage(msgResult.Messages[0].ReceiptHandle)
		if err != nil {
			fmt.Println("Got an error deleting the message:")
			responChan <- domain.ResponMsgWaitChan{
				IsPay: false,
				Err:   err,
			}
			return
		}

		fmt.Printf("wait for %s is FInished", *id)
		responChan <- domain.ResponMsgWaitChan{
			IsPay: respon == "true",
			Err:   nil,
		}
		return
	}
}
