package main

import (
	"log"
)

type OrderLogger struct {
}

type Order struct {
	OrderNumber  int
	CustomerName string
	OrderAmount  float64
}

func (logger *OrderLogger) AddOrder(order Order) {
	log.Printf("Добавлен заказ #%d, Имя клиента: %s, Сумма заказа: $%.2f\n", order.OrderNumber, order.CustomerName, order.OrderAmount)
}

func NewOrderLogger() *OrderLogger {
	return &OrderLogger{}
}
func main() {
	logger := NewOrderLogger()

	order := Order{1, "Иванов", 100.50}
	logger.AddOrder(order)
}
