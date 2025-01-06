package telegram

import (
	"context"
	"time"

	"github.com/robfig/cron"
)

var (
	service *Service
)

func SendVocabulary() error {
	job := cron.New()
	job.AddFunc("0 0 * * * *", func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // Tăng từ 5s lên 10s
		defer cancel()
		service.SendVocabularyToAll(ctx)
	})

	job.Start()

	return nil
}
