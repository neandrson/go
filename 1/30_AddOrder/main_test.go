package main

import (
	"bytes"
	"fmt"
	"log"
	"testing"
)

func TestOrderLogger(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	logger := NewOrderLogger()

	order := Order{1, "Иванов", 100.50}
	logger.AddOrder(order)

	expectedLog := fmt.Sprintf("Добавлен заказ #%d, Имя клиента: %s, Сумма заказа: $%.2f\n", order.OrderNumber, order.CustomerName, order.OrderAmount)
	actualLog := buf.String()

	if actualLog[20:] != expectedLog {
		t.Errorf("Лог не соответствует ожиданиям. Ожидалось:\n%s\nПолучено:\n%s\n", expectedLog, actualLog)
	}
}
