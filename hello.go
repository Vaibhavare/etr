package main

import (
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	fmt.Println("Current Time is: ", time.Now())
	bot, err := tgbotapi.NewBotAPI("6660243646:AAE1yq4g986w5ibk1SSsoIytm3aMVVCcKAQ")
	if err != nil {
		panic(err)
	}
	// Set up a ticker to send messages every hour
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()
	// Channel ID of the Telegram channel where you want to send messages
	channelID := int64(-4069295676)
	// Function to send messages
	sendMessage := func() {
		currentTime := time.Now().Format("15:04:05")
		msg := tgbotapi.NewMessage(channelID, "Current time: "+currentTime)
		// Send the message
		_, err := bot.Send(msg)
		if err != nil {
			fmt.Println("Error sending message:", err)
		}
	}
	// Send an initial message
	sendMessage()
	// Continuously send messages every hour
	for range ticker.C {
		sendMessage()
	}
}
