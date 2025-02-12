package telegram

import (
	"log"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/lkphuong/toeic-vocabulary/internal/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	collection *mongo.Collection
	bot        *tgbotapi.BotAPI
)

func init() {
	client := database.New()
	if client == nil {
		panic("Failed to initialize MongoDB client")
	}
	dbName := os.Getenv("DB_NAME")
	collection = database.MongoClient.Database(dbName).Collection("chat_ids")

	var err error
	// bot, err = tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	// if err != nil {
	// 	log.Fatal("Error initializing bot:", err)
	// }
	retries := 10
	for i := 0; i < retries; i++ {
		bot, err = tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
		if err == nil {
			break
		}
		log.Printf("Retrying (%d/%d): %v", i+1, retries, err)
		time.Sleep(5 * time.Second)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
}
