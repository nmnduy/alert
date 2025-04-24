package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	messagePtr := flag.String("message", "", "The message to send via Telegram (required)")

	// Define a custom usage function for the -help flag
	flag.Usage = func() {
		// Use os.Stderr for usage and error messages
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", filepath.Base(os.Args[0]))
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "This program sends a message to a specified Telegram chat using a bot.\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "Required environment variables:\n")
		fmt.Fprintf(os.Stderr, "  ALERT_CHAT_ID    - The target Telegram chat ID.\n")
		fmt.Fprintf(os.Stderr, "  ALERT_BOT_TOKEN  - The Telegram bot token.\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "Flags:\n")
		flag.PrintDefaults() // Print the default help information for flags
	}

	// Get Telegram chat ID from environment variable
	chatID := os.Getenv("ALERT_CHAT_ID")
	if chatID == "" {
		fmt.Fprintln(os.Stderr, "Error: ALERT_CHAT_ID environment variable is not set")
		flag.Usage() // Print usage information on error
		os.Exit(1)
	}

	// Get Telegram bot token from environment variable
	botToken := os.Getenv("ALERT_BOT_TOKEN")
	if botToken == "" {
		fmt.Fprintln(os.Stderr, "Error: ALERT_BOT_TOKEN environment variable is not set")
		flag.Usage() // Print usage information on error
		os.Exit(1)
	}

	flag.Parse() // Parse flags. This will automatically handle -help and call flag.Usage

	if *messagePtr == "" {
		fmt.Fprintln(os.Stderr, "Error: -message flag is required")
		flag.Usage() // Print usage information on error
		os.Exit(1)
	}

	// Send message via Telegram bot
	// (Rest of the code remains unchanged)
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s",
		botToken, chatID, *messagePtr)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error sending Telegram message:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// Consider logging the response body for more details on Telegram API errors
		fmt.Fprintf(os.Stderr, "Error: Telegram API returned non-200 status: %s\n", resp.Status)
		os.Exit(1)
	}

	fmt.Println("Message sent successfully to Telegram")
}
