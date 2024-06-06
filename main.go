package main

import (
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// Get environment variables directly
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
		"lightweight":  "audio/lightweight.mp3",
		"pilgrims":     "audio/ask-the-pilgrims.mp3",
		"2pm":          "audio/2pm.mp3",
		"devilz":       "audio/devilz.mp3",
		"beez":         "audio/beez.mp3",
		"whadiya":      "audio/whadiya.mp3",
		"outtahere":    "audio/outtahere.mp3",
		"kill cunts":   "audio/killCunts.mp3",
		"fuck ye":      "audio/fuckye.mp3",
		"fishy":        "audio/fishy.mp3",
		"wut":          "audio/wut.mp3",
		"wer u be":     "audio/werube.mp3",
		"doctah":       "audio/doctah.mp3",
		"knowledge":    "audio/dropKnowledge.mp3",
		"fake news":    "audio/fakeNews.mp3",
		"silly woman":  "audio/sillyWoman.mp3",
		"bomb planted": "audio/csBomb.mp3",
		"nup yep":      "audio/nupyep.mp3",
		"not today":    "audio/notToday.mp3",
		"ye na":        "audio/yena.mp3",
		"gay cunt":     "audio/gaycunt.mp3",
		"blue cunts":   "audio/blueCunts.mp3",
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
