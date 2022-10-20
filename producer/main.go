package main

import (
	"producer/controllers"
	"producer/services"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func main() {
	producer, err := sarama.NewSyncProducer(viper.GetStringSlice("kafka.servers"), nil)
	if err != nil {
		panic(err)
	}

	defer producer.Close()

	eventProducer := services.NewEventProducer(producer)
	messageService := services.NewMessageService(eventProducer)
	messageController := controllers.NewMessageController(messageService)

	app := fiber.New()

	app.Use(cors.New())

	app.Post("/v1/message", messageController.OpenMessage)

	app.Listen(":8000")
}

// func main() {
// 	servers := []string{"localhost:9092"}

// 	producer, err := sarama.NewSyncProducer(servers, nil)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer producer.Close()

// 	msg := sarama.ProducerMessage{
// 		Topic: "exam",
// 		Value: sarama.StringEncoder("Hello World"),
// 	}

// 	p, o, err := producer.SendMessage(&msg)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("partition=%v, offset=%v", p, o)
// }
