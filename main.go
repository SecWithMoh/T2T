package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "strconv"

    "github.com/dghubble/oauth1"
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TwitterResponse struct {
    Data struct {
        ID string `json:"id"`
    } `json:"data"`
}

const (
    consumerKey    = "YOUR_TWITTER_CONSUMER_KEY"
    consumerSecret = "YOUR_TWITTER_CONSUMER_SECRET"
    accessToken    = "YOUR_TWITTER_ACCESS_TOKEN"
    accessSecret   = "YOUR_TWITTER_ACCESS_SECRET"
    telegramBotToken = "YOUR_TELEGRAM_BOT_TOKEN"
    authorizedUserID = "YOUR_TELEGRAM_USER_ID" // Replace with your Telegram User ID
)

func main() {
    bot, err := tgbotapi.NewBotAPI(telegramBotToken)
    if err != nil {
        log.Fatalf("Error creating Telegram bot: %v", err)
    }

    log.Printf("Authorized on account %s", bot.Self.UserName)

    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60
    updates, err := bot.GetUpdatesChan(u)

    for update := range updates {
        if update.Message == nil {
            continue
        }

        if strconv.Itoa(update.Message.From.ID) != authorizedUserID {
            log.Printf("Unauthorized user: %d", update.Message.From.ID)
            continue
        }

        if update.Message.IsCommand() {
            switch update.Message.Command() {
            case "tweet":
                tweetText := update.Message.CommandArguments()
                tweetURL, err := sendTweet(tweetText)
                if err != nil {
                    log.Printf("Error sending tweet: %v", err)
                    bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Error: "+err.Error()))
                    continue
                }
                bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Tweet sent successfully: "+tweetURL))

            case "commands":
                commands := "Commands:\n/tweet - Post a tweet\n/commands - Show this message"
                bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, commands))
            }
        }
    }
}

func sendTweet(tweetText string) (string, error) {
    config := oauth1.NewConfig(consumerKey, consumerSecret)
    token := oauth1.NewToken(accessToken, accessSecret)
    httpClient := config.Client(oauth1.NoContext, token)

    requestBody := map[string]interface{}{
        "text": tweetText,
    }
    requestBodyBytes, err := json.Marshal(requestBody)
    if err != nil {
        return "", fmt.Errorf("error marshaling request body: %v", err)
    }

    request, err := http.NewRequest("POST", "https://api.twitter.com/2/tweets", bytes.NewBuffer(requestBodyBytes))
    if err != nil {
        return "", fmt.Errorf("error creating request: %v", err)
    }
    request.Header.Set("Content-Type", "application/json")

    response, err := httpClient.Do(request)
    if err != nil {
        return "", fmt.Errorf("error sending request to Twitter API: %v", err)
    }
    defer response.Body.Close()

    responseBody, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return "", fmt.Errorf("error reading response body: %v", err)
    }

    log.Printf("Twitter API response: %s", string(responseBody))

    if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated {
        return "", fmt.Errorf("failed to post tweet, status code: %d, response: %s", response.StatusCode, string(responseBody))
    }

    var twitterResponse TwitterResponse
    err = json.Unmarshal(responseBody, &twitterResponse)
    if err != nil {
        return "", fmt.Errorf("error unmarshaling Twitter response: %v", err)
    }

    tweetID := twitterResponse.Data.ID
    if tweetID == "" {
        return "", fmt.Errorf("failed to extract tweet ID, response: %s", string(responseBody))
    }

    tweetURL := fmt.Sprintf("https://twitter.com/user/status/%s", tweetID)
    return tweetURL, nil
}
