package main

import (
	"container/list"
	"sync"
)

// структура MRU кеша
type MRUCache struct {
	capacity int
	cache    map[string]*list.Element
	list     *list.List
	mutex    sync.Mutex
}

type CacheEntry struct {
	key   string
	value string
}

// возвращает новый инстанс кеша размером capacity
func NewMRUCache(capacity int) *MRUCache {
	return &MRUCache{
		capacity: capacity,
		cache:    make(map[string]*list.Element),
		list:     list.New(),
	}
}

// устанавливает значени value ключу key
func (c *MRUCache) Set(key, value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// проверка того, существует ли ключ уже в кэше
	if element, ok := c.cache[key]; ok {
		// обновляем существующую запись и перемещаем ее в начало (использовалась в последний раз)
		element.Value.(*CacheEntry).value = value
		c.list.MoveToFront(element)
	} else {
		// добавление новой записи в кеш
		entry := &CacheEntry{key: key, value: value}
		element := c.list.PushFront(entry)
		c.cache[key] = element

		// проверяем, заполнен ли кеш (есть ли место), при необходимости удаляем наименее недавно использованный элемент
		if c.list.Len() > c.capacity {
			oldest := c.list.Back()
			if oldest != nil {
				delete(c.cache, oldest.Value.(*CacheEntry).key)
				c.list.Remove(oldest)
			}
		}
	}
}

// получает значение и флаг его начличия по ключу key
func (c *MRUCache) Get(key string) (string, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if element, ok := c.cache[key]; ok {
		// перемещение доступного элемента в начало списка (последний раз использованный)
		c.list.MoveToFront(element)
		return element.Value.(*CacheEntry).value, true
	}

	return "", false
}

func main() {

}
