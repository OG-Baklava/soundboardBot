package main

import (
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN is not set")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Printf("Received message: %s", update.Message.Text)
			handleUpdate(bot, update)
		}
	}
}

func handleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := strings.ToLower(update.Message.Text)
	log.Printf("Processing message from chat %d: %s", update.Message.Chat.ID, text)

	if strings.Contains(text, "lightweight") {
		log.Println("Trigger: lightweight")
		sendAudio(bot, update.Message.Chat.ID, "lightweight.mp3")
	} else if strings.Contains(text, "pilgrims") {
		log.Println("Trigger: pilgrims")
		sendAudio(bot, update.Message.Chat.ID, "ask-the-pilgrims.mp3")
	} else if strings.Contains(text, "2pm") {
		log.Println("Trigger: 2pm")
		sendAudio(bot, update.Message.Chat.ID, "2pm.mp3")

	} else if strings.Contains(text, "devilz") {
		log.Println("Trigger: devilz")
		sendAudio(bot, update.Message.Chat.ID, "devilz.mp3")
	} else if strings.Contains(text, "beez") {
		log.Println("Trigger: beez")
		sendAudio(bot, update.Message.Chat.ID, "beez.mp3")
	} else if strings.Contains(text, "whadiya") {
		log.Println("Trigger: whadiya")
		sendAudio(bot, update.Message.Chat.ID, "whadiya.mp3")
	} else if strings.Contains(text, "outtahere") {
		log.Println("Trigger: outtahere")
		sendAudio(bot, update.Message.Chat.ID, "outtahere.mp3")
	} else if strings.Contains(text, "kill cunts") {
		log.Println("Trigger: killcunts")
		sendAudio(bot, update.Message.Chat.ID, "killCunts.mp3")
	} else if strings.Contains(text, "fuckye") {
		log.Println("Trigger: fuckye")
		sendAudio(bot, update.Message.Chat.ID, "fuckye.mp3")
	} else if strings.Contains(text, "doctah") {
		log.Println("Trigger: doctah")
		sendAudio(bot, update.Message.Chat.ID, "doctah.mp3")
	} else if strings.Contains(text, "knowledge") {
		log.Println("Trigger: knowledge")
		sendAudio(bot, update.Message.Chat.ID, "dropKnowledge.mp3")
	} else if strings.Contains(text, "sound2") {
		log.Println("Trigger: sound2")
		sendAudio(bot, update.Message.Chat.ID, "so")
	} else if strings.Contains(text, "sound2") {
		log.Println("Trigger: sound2")
		sendAudio(bot, update.Message.Chat.ID, "so")
	} else if strings.Contains(text, "sound2") {
		log.Println("Trigger: sound2")
		sendAudio(bot, update.Message.Chat.ID, "so")
	} else {
		log.Printf("No trigger found in message: %s", text)
	}
}

func sendAudio(bot *tgbotapi.BotAPI, chatID int64, filePath string) {
	audioFile, err := os.Open(filePath)
	if err != nil {
		log.Printf("Failed to open audio file: %v", err)
		return
	}
	defer audioFile.Close()

	fileReader := tgbotapi.FileReader{
		Name:   filePath,
		Reader: audioFile,
	}
	audio := tgbotapi.NewAudio(chatID, fileReader)

	_, err = bot.Send(audio)
	if err != nil {
		log.Printf("Failed to send audio: %v", err)
	} else {
		log.Printf("Sent audio: %s", filePath)
	}
}
