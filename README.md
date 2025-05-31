# Alert

A simple command-line tool to send messages to a Telegram chat via a bot.

## Features

- Sends a message to a specified Telegram chat using your bot.
- Takes configuration from environment variables.
- Simple CLI interface.

## Prerequisites

- [Go](https://golang.org/dl/) 1.23.4 or higher installed.
- A Telegram bot token. You can create a bot and get a token from [BotFather](https://t.me/botfather) on Telegram.
- The chat ID of the Telegram user or group where messages will be sent.

## Installation

Clone this repository and build the binary:


## Usage

### Set up environment variables

Before running the program, you need to set two environment variables:

- `ALERT_BOT_TOKEN`: Your Telegram bot token.
- `ALERT_CHAT_ID`: The chat ID to which the message will be sent.

Example:


### Running the program

Use the `-message` flag to specify the message to send:


If the message is sent successfully, you will see:


## Notes

- If either environment variable is not set, or the `-message` flag is omitted, the program will print an error and exit.
- The chat ID can be your own user ID or a group/channel ID. For groups, the ID may be negative (e.g., `-1001234567890`).
- You can get your chat ID by messaging your bot or by using services like [userinfobot](https://t.me/userinfobot).

## License

MIT
