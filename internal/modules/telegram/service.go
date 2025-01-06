package telegram

import (
	"context"
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/lkphuong/toeic-vocabulary/internal/models"
	"github.com/lkphuong/toeic-vocabulary/internal/modules/vocabulary"
)

var (
	repository        *Repository
	vocabulary_module *vocabulary.Repository
)

type Service struct{}

func (s *Service) SendVocabularyToAll(ctx context.Context) {

	vocab, _ := vocabulary_module.GetRandomVocabulary(ctx)

	chats, err := repository.GetChatIDs(ctx)
	if err != nil {
		log.Println("Error while getting chat ids", err)
	}

	message := fmt.Sprintf(
		"*Word*: `%s`\n*Type*: `%s`\n*Meaning*: %s\n\n*Examples*:\n%s\n\n*Related Words*:\n%s\n\n*Notes*:\n%s",
		vocab.Word,
		vocab.Type,
		vocab.Meaning,
		formatExamples(vocab.Examples),
		formatRelatedWords(vocab.RelatedWords),
		formatNotes(vocab.Notes),
	)

	for _, chat := range chats {

		msg := tgbotapi.NewMessage(chat.ChatID, message)
		msg.ParseMode = "markdown"
		_, err = bot.Send(msg)
		if err != nil {
			log.Printf("Error sending message to chat ID %d: %v", chat.ChatID, err)
		} else {
			log.Printf("Message sent to chat ID %d", chat.ChatID)
		}
	}
}

func (s *Service) SendVocabularyToUser(ctx context.Context, chatID int64) {

	vocab, _ := vocabulary_module.GetRandomVocabulary(ctx)

	message := fmt.Sprintf(
		"*Word*: `%s`\n*Type*: `%s`\n*Meaning*: %s\n\n*Examples*:\n%s\n\n*Related Words*:\n%s\n\n*Notes*:\n%s",
		vocab.Word,
		vocab.Type,
		vocab.Meaning,
		formatExamples(vocab.Examples),
		formatRelatedWords(vocab.RelatedWords),
		formatNotes(vocab.Notes),
	)
	msg := tgbotapi.NewMessage(chatID, message)
	msg.ParseMode = "markdown"
	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Error sending message to chat ID %d: %v", chatID, err)
	} else {
		log.Printf("Message sent to chat ID %d", chatID)
	}
}

func SaveChatID() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		chatID := update.Message.Chat.ID
		username := update.Message.From.UserName

		log.Printf("[%s] %s", username, update.Message.Text)

		if update.Message.IsCommand() {
			command := update.Message.Command()
			switch command {
			case "more":
				service.SendVocabularyToUser(ctx, chatID)
			}

		}

		isExist := repository.GetByChatID(ctx, chatID)

		if isExist {
			log.Println("Chat id is existed ", chatID)
			continue
		}

		err := repository.Save(ctx, chatID, username)
		if err != nil {
			log.Println("Error while saving chat id", err)
		}

		log.Println("Chat id saved successfully")
	}

}

// Format examples as a string
func formatExamples(examples []models.Example) string {
	var formatted string
	for _, example := range examples {
		formatted += fmt.Sprintf("• `%s`: %s\n", example.English, example.Vietnamese)
	}
	return formatted
}

// Format related words as a string
func formatRelatedWords(relatedWords []models.RelatedWord) string {
	var formatted string
	for _, relatedWord := range relatedWords {
		formatted += fmt.Sprintf("• `%s`: %s (%s)\n", relatedWord.Word, relatedWord.Meaning, relatedWord.Tag)
	}
	return formatted
}

func formatNotes(notes []models.Note) string {
	var formatted string
	for _, note := range notes {
		formatted += fmt.Sprintf("• `%s`: %s\n", note.Word, note.Note)
	}
	return formatted
}
