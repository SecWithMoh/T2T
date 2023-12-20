# Telegram Twitter Bot 🤖

This is a Go-based Telegram bot that allows you to post tweets to Twitter. 🐦

## Prerequisites 📋

Before you begin, ensure you have the following:

- Go programming language installed.
- Twitter API credentials (consumer key, consumer secret, access token, access secret).
- Telegram bot token.
- Authorized Telegram user ID. 👤

## Configuration ⚙️

1. Open the `main.go` file.
2. Replace the following constants with your API credentials and settings:

```go
const (
    consumerKey       = "YOUR_TWITTER_CONSUMER_KEY"
    consumerSecret    = "YOUR_TWITTER_CONSUMER_SECRET"
    accessToken       = "YOUR_TWITTER_ACCESS_TOKEN"
    accessSecret      = "YOUR_TWITTER_ACCESS_SECRET"
    telegramBotToken  = "YOUR_TELEGRAM_BOT_TOKEN"
    authorizedUserID = "YOUR_TELEGRAM_USER_ID" // Replace with your Telegram User ID
)

```


## Usage 🚀

To use this Telegram bot, follow these commands:

- `/tweet <tweet_text>`: Post a tweet.
- `/commands`: Show available commands.

## How It Works 🛠️

The bot uses the Twitter API to post tweets based on your commands.

## License 📜

This project is licensed under the [GNU General Public License, version 3 (GPL-3.0)](LICENSE) - see the [LICENSE](LICENSE) file for details.

## Acknowledgments 🙌

- [GitHub](https://github.com) for hosting this repository.
- [Telegram Bot API](https://github.com/go-telegram-bot-api/telegram-bot-api) for the Telegram bot library.

## Author ✍️

SecWithMoh




