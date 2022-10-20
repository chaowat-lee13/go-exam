package services

import (
	"events"
	"log"
	"producer/commands"
)

type MessageService interface {
	OpenMessage(command commands.ExamMessageCommand) (id string, err error)
}

type messageService struct {
	eventProducer EventProducer
}

func NewMessageService(eventProducer EventProducer) MessageService {
	return messageService{eventProducer}
}

func (obj messageService) OpenMessage(command commands.ExamMessageCommand) (id string, err error) {

	event := events.ExamMessage{
		Msg_id:        command.Msg_id,
		Sender:        command.Sender,
		Msg:           command.Msg,
		Code:          command.Code,
		Received_time: command.Received_time,
	}

	log.Printf("%#v", event)
	return string(event.Msg_id), obj.eventProducer.Produce(event)
}
