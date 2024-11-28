package main

import (
	"bytes"
	"fmt"
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
	logger.orders = append(logger.orders, order)
	fmt.Println("Order added to logger:", order)
	log.Printf("Добавлен заказ #%d, Имя клиента: %s, Сумма заказа: $%.2f\n", order.OrderNumber, order.CustomerName, order.OrderAmount)
}

func NewOrderLogger() logger {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	return buf
}
func main() {
	logger := NewOrderLogger()

	order := Order{1, "Иванов", 100.50}
	logger.AddOrder(order)
}
