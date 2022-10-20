package controllers

import (
	"producer/commands"
	"producer/services"
	"time"

	"github.com/gofiber/fiber/v2"
)

type MessageController interface {
	OpenMessage(c *fiber.Ctx) error
}

type messageController struct {
	messageService services.MessageService
}

func NewMessageController(messageService services.MessageService) MessageController {
	return messageController{messageService}
}

func (obj messageController) OpenMessage(c *fiber.Ctx) error {
	command := commands.ExamMessageCommand{}

	err := c.BodyParser(&command)
	if err != nil {
		return err
	}

	_, err = obj.messageService.OpenMessage(command)
	if err != nil {
		return err
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{
		"Code":          "OK",
		"Received_time": time.Now(),
	})
}
