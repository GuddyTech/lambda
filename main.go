package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

type MyEvent struct {
	Message string `json:"message"`
}

func handler(ctx context.Context, event MyEvent) error {
	// Create a new AWS session
	sess := session.Must(session.NewSession())

	// Create a new SNS client
	svc := sns.New(sess)

	// Define the SNS topic ARN
	topicArn := "arn:aws:sns:us-east-1:509344471761:lamda"

	// Prepare the message to be sent
	message := event.Message

	// Publish the message to SNS
	_, err := svc.Publish(&sns.PublishInput{
		Message:  &message,
		TopicArn: &topicArn,
	})
	if err != nil {
		return fmt.Errorf("failed to publish SNS message: %v", err)
	}

	return nil
}

func main() {
	lambda.Start(handler)
}