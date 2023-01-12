package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	database "ql/first-steps/internal/order/infra/database/order_repository"
	"ql/first-steps/internal/order/usecase/calculate_price"
	"ql/first-steps/pkg/rabbitmq"

	_ "github.com/mattn/go-sqlite3"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	db, err := sql.Open("sqlite3", "./orders.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repository := database.NewOrderRepository(db)
	usecase := calculate_price.NewCalculateFinalPriceUseCase(repository)
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	out := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, out)

	for msg := range out {
		var inputDTO calculate_price.OrderInputDTO

		err := json.Unmarshal(msg.Body, &inputDTO)
		if err != nil {
			panic(err)
		}
		
		outputDTO, err := usecase.Execute(inputDTO)
		if err != nil {
			panic(err)
		}

		msg.Ack(false)
		fmt.Println(outputDTO)
	}
}