package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Get Telegram chat ID from environment variable
	chatID := os.Getenv("ALERT_CHAT_ID")
	if chatID == "" {
		fmt.Println("Error: ALERT_CHAT_ID environment variable is not set")
		os.Exit(1)
	}

	// Get Telegram bot token from environment variable
	botToken := os.Getenv("ALERT_BOT_TOKEN")
	if botToken == "" {
		fmt.Println("Error: ALERT_BOT_TOKEN environment variable is not set")
		os.Exit(1)
	}

	messagePtr := flag.String("message", "", "The message to send via Telegram")
	flag.Parse()

	if *messagePtr == "" {
		fmt.Println("Error: -message flag is required")
		os.Exit(1)
	}

	// Send message via Telegram bot
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s",
		botToken, chatID, *messagePtr)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending Telegram message:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: Telegram API returned non-200 status:", resp.Status)
		os.Exit(1)
	}

	fmt.Println("Message sent successfully to Telegram")
}
