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

	// Map of trigger words to their corresponding audio file paths
	triggers := map[string]string{
		"lightweight":  "lightweight.mp3",
		"pilgrims":     "ask-the-pilgrims.mp3",
		"2pm":          "2pm.mp3",
		"devilz":       "devilz.mp3",
		"beez":         "beez.mp3",
		"whadiya":      "whadiya.mp3",
		"outtahere":    "outtahere.mp3",
		"kill cunts":   "killCunts.mp3",
		"fuck ye":      "fuckye.mp3",
		"fishy":        "fishy.mp3",
		"wut":          "wut.mp3",
		"wer u be":     "werube.mp3",
		"doctah":       "doctah.mp3",
		"knowledge":    "dropKnowledge.mp3",
		"fake news":    "fakeNews.mp3",
		"silly woman":  "sillyWoman.mp3",
		"bomb planted": "csBomb.mp3",
		"nup yep":      "nupyep.mp3",
		"not today":    "notToday.mp3",
		"ye na":        "yena.mp3",
		"gay cunt":     "gaycunt.mp3",
		"blue cunts":   "blueCunts.mp3",
	}

	for trigger, filePath := range triggers {
		if strings.Contains(text, trigger) {
			log.Printf("Trigger: %s", trigger)
			sendAudio(bot, update.Message.Chat.ID, filePath)
			return
		}
	}

	log.Printf("No trigger found in message: %s", text)
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
