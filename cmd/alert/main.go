package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

func main() {
	// Get recipient email from environment variable
	recipientEmail := os.Getenv("ALERT_EMAIL")
	if recipientEmail == "" {
		fmt.Println("Error: ALERT_EMAIL environment variable is not set")
		os.Exit(1)
	}

	// Parse command line arguments
	messagePtr := flag.String("message", "", "The message to send via email")
	flag.Parse()

	if *messagePtr == "" {
		fmt.Println("Error: -message flag is required")
		os.Exit(1)
	}

	// Create new AWS session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create SES client
	svc := ses.New(sess)

	// Send email
	_, err := svc.SendEmail(&ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{aws.String(recipientEmail)},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Data: aws.String(*messagePtr),
				},
			},
			Subject: &ses.Content{
				Data: aws.String("Alert Notification"),
			},
		},
		Source: aws.String(recipientEmail), // Replace with a verified SES email if different
	})

	if err != nil {
		fmt.Println("Error sending email:", err)
		os.Exit(1)
	}

	fmt.Println("Email sent successfully to:", recipientEmail)
}
