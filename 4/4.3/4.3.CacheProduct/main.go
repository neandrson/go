package main

import (
	"errors"
	"sync"
	"time"
)

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

type Cache struct {
	products map[int]Product // Кэш продуктов
	ttl      time.Duration   // Время жизни записи в кэше
	mutex    sync.RWMutex    // мьютекс для синхронизации конкурентного доступа к кешу
}

// Функция получения информации о товаре:
func getProduct(productId int, db map[int]Product, cache *Cache) (Product, error) {
	if product, found := cache.Get(productId); found {
		// Возвращаем данные из кеша
		return product, nil
	}

	// Данных в кеше нет, загружаем из БД
	if product, found := db[productId]; found {
		// Сохраняем данные в кеш
		cache.Set(productId, product)
		return product, nil
	}
	return Product{}, errors.New("not found")
}

// Функция обновления информации о товаре (фейк-функция выполняющая роль базы данных):
func updateProduct(productId int, newProduct Product, db map[int]Product) error {
	db[productId] = newProduct
	return nil
}

// Кеш продуктов:
func NewCache() *Cache {
	return &Cache{
		products: make(map[int]Product),
		ttl:      2 * time.Minute,
	}
}

// Получение продукта из кэша:
func (c *Cache) Get(productId int) (Product, bool) {
	c.mutex.RLock()         // получение блокировки на запись для эксклюзивного доступа
	defer c.mutex.RUnlock() // снятие блокировки записи при завершении работы функции
	product, ok := c.products[productId]
	return product, ok
}

// Сохранение продукта в кэш:
func (c *Cache) Set(productId int, product Product) {
	c.mutex.Lock()                  // получение блокировки на запись для эксклюзивного доступа
	defer c.mutex.Unlock()          // снятие блокировки записи при завершении работы функции
	c.products[productId] = product // установка значения в кеше по ключу
}

// Инвалидация кэша:
func (c *Cache) Invalidate(productId int) {
	delete(c.products, productId)
}

func main() {

}
