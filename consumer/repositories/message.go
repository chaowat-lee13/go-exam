package repositories

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	Msg_id        int
	Sender        string
	Msg           string
	Code          string
	Received_time time.Time
}

type MessageRepository interface {
	Save(message Message) error
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	db.AutoMigrate(&Message{})
	return messageRepository{db}
}

func (obj messageRepository) Save(message Message) error {
	return obj.db.Save(message).Error
}
