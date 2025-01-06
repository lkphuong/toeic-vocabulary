package models

import "time"

type ChatID struct {
	ChatID    int64     `json:"chat_id" bson:"chat_id"`
	Username  string    `json:"username" bson:"username"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
