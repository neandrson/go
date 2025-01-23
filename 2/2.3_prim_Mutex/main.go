package main

import (
	"log"
	"sync"
)

type DataRetriever interface {
	Retrieve(ID string) (*Data, error)
}

type Data struct {
	ID string // для упрощения содержит только ID
}

// кэш для хранения данных
type Cache struct {
	// Пример 1
	// данные будем хранить здесь
	m  map[string]*Data
	dr DataRetriever

	// Пример 2
	//mu sync.Mutex

	// Пример 3
	mu sync.RWMutex
}

// Пример 4
var m sync.Mutex

func first() {
	m.Lock()
	defer m.Unlock()
	second() // lock
}

func second() {
	m.Lock() // здесь будет вызов m.Lock() второй раз
	defer m.Unlock()
	// далее — основное тело функции 
}

// создаём новый объект
func NewCache(dr DataRetriever) *Cache {
	return &Cache{
		m:  make(map[string]*Data),
		dr: dr,
	}
}

func (c *Cache) Get(ID string) (Data, bool) {
	// Пример 1
	// проверим, есть ли данные в кэше
	/*data, exists := c.m[ID]
	// нашли в мапе — вернём значение
	if exists {
		return *data, true
	}
	// данные не нашли — нужно запросить
	data, err := c.dr.Retrieve(ID)
	if err != nil {
		// ошибка получения данных — запишем в лог
		log.Printf("c.dr.Retrieve(ID): %s", err)
		// вернём пустое значение
		return Data{}, false
	}
	// получили значение — запомним
	c.m[data.ID] = data
	// вернём полученное значение
	return *data, true*/

	//Пример 2
	/*c.mu.Lock()
	data, exists := c.m[ID] // теперь доступ к мапе внутри критической секции
	c.mu.Unlock()*/

	// Пример 3
	c.mu.RLock()
	// теперь в эту секцию могут зайти несколько горутин
	data, exists := c.m[ID]
	c.mu.RUnlock()
	// далее без изменений

	// нашли в мапе — вернём значение
	if exists {
		return *data, true
	}
	// запрос данных из базы — не в критической секции
	data, err := c.dr.Retrieve(ID)
	if err != nil {
		// ошибка получения данных — запишем в лог
		log.Printf("c.dr.Retrieve(ID): %s", err)
		// вернём пустое значение
		return Data{}, false
	}
	// перед обращением к мапе снова заблокируем мьютекс
	c.mu.Lock()
	// разблокируем при выходе из функциии
	defer c.mu.Unlock()
	// внутри критической секции нужно снова проверить на наличие значения в мапе
	data, exists = c.m[data.ID]
	if exists {
		return *data, true
	}
	// получили значение — запомним
	c.m[data.ID] = data
	// вернём полученное значение
	return *data, true
}

