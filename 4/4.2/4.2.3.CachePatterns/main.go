package main

import (
	"container/list"
	"sync"
)

// Структура кеша
// Обязательно содержит 2 публичных поля
// UpperBound - верхняя граница размера кеша
// LowerBound - нижняя граница размера кеша
// Если len > UpperBound, кеш автоматически вытеснит значения до нижней границы
// Если любое из этих значений 0 - то этого не произойдет
type Cache struct {
	UpperBound int
	LowerBound int
	values     map[string]*list.Element
	list       *list.List
	lock       sync.Mutex
}

type CacheEntry struct {
	count int
	key   string
	value interface{}
}

// Создает инстанс кеша
func New() *Cache {
	return &Cache{
		values: make(map[string]*list.Element),
		list:   list.New(),
	}
}

func (c *Cache) increment(el *list.Element) {
	el.Value.(*CacheEntry).count++
	next := el.Next()
	for next != nil {
		if next.Value.(*CacheEntry).count > el.Value.(*CacheEntry).count {
			break
		}
		next = next.Next()
	}
	if next == nil {
		c.list.MoveToBack(el)
	} else {
		c.list.MoveBefore(el, next)
	}
}

// Проверяет, содержит, ли кэш ключ
func (c *Cache) Has(key string) bool {
	c.lock.Lock()
	defer c.lock.Unlock()
	_, found := c.values[key]
	return found
}

// Возвращает значение по ключу, если оно существует
// Возвращает nil, если не существует
func (c *Cache) Get(key string) interface{} {
	c.lock.Lock()
	defer c.lock.Unlock()

	if el, ok := c.values[key]; ok {
		c.increment(el)
		return el.Value.(*CacheEntry).value
	}
	return nil
}

// Сохраняет значение по ключу
func (c *Cache) Set(key string, value interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if el, found := c.values[key]; found {
		el.Value.(*CacheEntry).value = value
		c.increment(el)
	} else {
		entry := &CacheEntry{key: key, value: value, count: 1}
		newElement := c.list.PushFront(entry)
		c.values[key] = newElement

		if c.list.Len() > c.UpperBound && c.LowerBound != 0 && c.UpperBound != 0 {
			c.Evict(c.list.Len() - c.LowerBound)
		}
	}
}

// Возвращает размер кеша
func (c *Cache) Len() int {
	return c.list.Len()
}

// Возвращает частоту обращений к ключу
func (c *Cache) GetFrequency(key string) int {
	if el, found := c.values[key]; found {
		return el.Value.(*CacheEntry).count
	}
	return 0
}

// Возвращает все ключи в кеше
func (c *Cache) Keys() []string {
	keys := make([]string, 0, c.Len())
	for el := c.list.Front(); el != nil; el = el.Next() {
		keys = append(keys, el.Value.(*CacheEntry).key)
	}
	return keys
}

// Удаляет заданное количество наименее часто используемых элементов элементов
// Возвращает количество удаленных элементов
func (c *Cache) Evict(count int) int {
	n := 0
	for el := c.list.Front(); el != nil && n < count; el = el.Next() {
		delete(c.values, el.Value.(*CacheEntry).key)
		c.list.Remove(el)
		n++
	}
	return n
}

func main() {

}
