package controlers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func CreateQueue(name string) error {

	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		return err
	}

	svc := sqs.NewFromConfig(cfg)

	input := &sqs.CreateQueueInput{
		QueueName: aws.String(name),
	}

	_, err = svc.CreateQueue(context.TODO(), input)

	if err != nil {
		return err
	}

	return nil
}

func ListQueues() ([]string, error) {

	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		return nil, err
	}

	svc := sqs.NewFromConfig(cfg)
	input := &sqs.ListQueuesInput{}

	rs, err := svc.ListQueues(context.TODO(), input)

	if err != nil {
		return nil, err
	}
	var queueURLs []string

	for _, i := range rs.QueueUrls {
		queueURLs = append(queueURLs, i)
	}

	return queueURLs, nil
}

func GetQueueURL(name string) (string, error) {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return "", err
	}

	svc := sqs.NewFromConfig(cfg)

	input := &sqs.GetQueueUrlInput{
		QueueName: aws.String(name),
	}

	res, err := svc.GetQueueUrl(context.TODO(), input)

	if err != nil {
		return "", err
	}

	return *res.QueueUrl, nil
}

func DeleteQueue(queueURL string) error {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return err
	}

	svc := sqs.NewFromConfig(cfg)

	input := &sqs.DeleteQueueInput{
		QueueUrl: aws.String(queueURL),
	}

	_, err = svc.DeleteQueue(context.TODO(), input)

	if err != nil {
		return err
	}

	return nil
}

func SendMessage(queueURL, messageBody string) error {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return err
	}

	svc := sqs.NewFromConfig(cfg)

	input := &sqs.SendMessageInput{
		QueueUrl:    aws.String(queueURL),
		MessageBody: aws.String(messageBody),
	}

	_, err = svc.SendMessage(context.TODO(), input)

	if err != nil {
		return err
	}

	return nil
}

func ReceiveMessages(queueURL string, maxMessages int) (*sqs.ReceiveMessageOutput, error) {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	svc := sqs.NewFromConfig(cfg)

	input := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queueURL),
		MaxNumberOfMessages: *aws.Int32(int32(maxMessages)),
	}

	rs, err := svc.ReceiveMessage(context.TODO(), input)

	if err != nil {
		return nil, err
	}

	return rs, nil
}

func DeleteMessage(queueURL, receiptHandle string) error {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return err
	}

	svc := sqs.NewFromConfig(cfg)

	input := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(queueURL),
		ReceiptHandle: aws.String(receiptHandle),
	}

	_, err = svc.DeleteMessage(context.TODO(), input)

	if err != nil {
		return err
	}

	return nil

}
