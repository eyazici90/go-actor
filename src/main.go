package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/viper"

	"go-projection/application"
	"go-projection/shared/rabbitmq"
)

var config application.AppConfig

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&config)

	if err != nil {
		panic(err)
	}
}

func main() {

	ctx, err := rabbitmq.NewContext(config.AmqpURL).
		WithExchange(config.Consumer.Exchange.Name, config.Consumer.Exchange.Type).
		WithQueue(config.Consumer.Exchange.QueName, config.Consumer.Exchange.BindingKey).
		Connect()

	if err != nil {
		log.Fatalln(err)
		return
	}

	ctx.StartConsumingBy(config.Consumer.Tag, func(msg rabbitmq.Message) {

		log.Println(msg.GetBody())

		j := json.RawMessage(msg.GetBody())

		bytes, _ := j.MarshalJSON()

		handleMsg(msg.GetMessageType(), bytes)
	})

	fmt.Scanln()
}

func handleMsg(msgType string, bytes []byte) {
	switch msgType {
	case "application.StockCreatedEvent":
		event := &application.StockCreatedEvent{}
		_ = json.Unmarshal(bytes, event)

		application.Send(event.ID, *event)

	case "application.ShippedToLocationEvent":
		event := &application.ShippedToLocationEvent{}
		_ = json.Unmarshal(bytes, event)

		application.Send(event.ID, *event)
	}
}

// var stockId = "123"

// func publishRnd(amqpURL string) {
// 	ctx, err := rabbitmq.NewContext(amqpURL).
// 		WithExchange("test", "fanout").
// 		Connect()
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}

// 	stockCreated := application.StockCreatedEvent{
// 		ID:         stockId,
// 		LocationId: "1",
// 		Amount:     10,
// 	}

// 	msg, _ := rabbitmq.NewMessage(stockCreated)

// 	err = ctx.Publish(msg)

// 	for i := 1; i < 10; i++ {
// 		event := application.ShippedToLocationEvent{
// 			ID:         stockId,
// 			LocationId: "1",
// 			Amount:     i,
// 		}
// 		m, _ := rabbitmq.NewMessage(event)

// 		ctx.Publish(m)
// 	}
// }
