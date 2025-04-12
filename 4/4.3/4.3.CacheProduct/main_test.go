package main

import (
	"testing"
)

func TestProductCache(t *testing.T) {
	// Имитация базы данных
	db := make(map[int]Product)
	db[1] = Product{ID: 1, Name: "Товар 1", Price: 100, Stock: 50}

	// Создание кэша
	cache := NewCache()

	// Проверка загрузки из базы данных и кэширования
	product, err := getProduct(1, db, cache)
	if err != nil {
		t.Fatalf("Error retrieving product: %s", err)
	}
	if product.ID != 1 {
		t.Errorf("Expected product ID 1, got %d", product.ID)
	}

	// Обновление информации о товаре в базе данных
	newProduct := Product{ID: 1, Name: "Товар 1", Price: 200, Stock: 30}
	err = updateProduct(1, newProduct, db)
	if err != nil {
		t.Fatalf("Error updating product: %s", err)
	}

	// Инвалидация кэша
	cache.Invalidate(1)

	// Проверка загрузки обновленной информации из базы данных
	updatedProduct, err := getProduct(1, db, cache)
	if err != nil {
		t.Fatalf("Error retrieving updated product: %s", err)
	}
	if updatedProduct.Price != 200 {
		t.Errorf("Expected updated price 200, got %f", updatedProduct.Price)
	}
}
