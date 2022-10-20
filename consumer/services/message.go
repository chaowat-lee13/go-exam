package services

import (
	"consumer/repositories"
	"encoding/json"
	"events"
	"log"
	"reflect"
)

type EventHandler interface {
	Handle(topic string, eventBytes []byte)
}

type messageEventHandler struct {
	messageRepo repositories.MessageRepository
}

func NewMessageEventHandler(messageRepo repositories.MessageRepository) EventHandler {
	return messageEventHandler{messageRepo}
}

func (obj messageEventHandler) Handle(topic string, eventBytes []byte) {
	switch topic {
	case reflect.TypeOf(events.ExamMessage{}).Name():
		event := &events.ExamMessage{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			log.Println(err)
			return
		}
		message := repositories.Message{
			Msg_id:        event.Msg_id,
			Sender:        event.Sender,
			Msg:           event.Msg,
			Code:          event.Code,
			Received_time: event.Received_time,
		}
		err = obj.messageRepo.Save(message)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
