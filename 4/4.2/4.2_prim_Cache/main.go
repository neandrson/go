package main

import (
	"fmt"
	"sync"
	"time"
)

// создание простого кеша ключ-значение с потокобезопасными операциями чтения и записи
/*type Cache struct {
	data  map[string]interface{} // мапа для хранения пар ключ-значение
	mutex sync.RWMutex           // мьютекс для синхронизации конкурентного доступа к кешу
}

// создание нового экземпляра кеша с инициализированной мапой данных
func NewCache() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}

// извлекает значение, связанное с данным ключом, из кеша
// Get() возвращает значение и признак, указывающий, был ли найден ключ
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mutex.RLock()         // acquire a read lock to allow multiple readers simultaneously
	defer c.mutex.RUnlock() // release the read lock when the function exits
	value, ok := c.data[key]
	return value, ok
}

// установка значения, связанного с данным ключом в кеше
// Set() получает блокировку на запись для обеспечения эксклюзивного доступа во время обновления
func (c *Cache) Set(key string, value interface{}) {
	c.mutex.Lock()         // получение блокировки на запись для эксклюзивного доступа
	defer c.mutex.Unlock() // снятие блокировки записи при завершении работы функции
	c.data[key] = value    // установка значения в кеше по ключу
}

func main() {
	cache := NewCache()

	cache.Set("username", "yandexlyceum")
	cache.Set("year", 2024)

	if value, ok := cache.Get("username"); ok {
		fmt.Println("Value for username:", value)
	} else {
		fmt.Println("username not found in the cache.")
	}

	if value, ok := cache.Get("year"); ok {
		fmt.Println("Value for year:", value)
	} else {
		fmt.Println("year not found in the cache.")
	}

	time.Sleep(2 * time.Second)

	if value, ok := cache.Get("username"); ok {
		fmt.Println("Value for username (after some time):", value)
	} else {
		fmt.Println("username not found in the cache after some time.")
	}

	if value, ok := cache.Get("year"); ok {
		fmt.Println("Value for year (after some time):", value)
	} else {
		fmt.Println("year not found in the cache after some time.")
	}
}*/

// -----------------
// создание LruCache, представляющего кеш LRU
/*type LRUCache struct {
	capacity int
	cache    map[string]*list.Element
	list     *list.List
	mutex    sync.Mutex
}

// создание записи кеша, представляющей запись в кеше LRU
type CacheEntry struct {
	key   string
	value interface{}
}

// NewLRUCache создает новый экземпляр LruCache с указанной ёмкостью
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[string]*list.Element),
		list:     list.New(),
	}
}

// извлечение значения, связанного с данным ключом, из кеша
func (lru *LRUCache) Get(key string) (interface{}, bool) {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()

	if element, ok := lru.cache[key]; ok {
		// перемещение доступного элемента в начало списка (последний раз использованный)
		lru.list.MoveToFront(element)
		return element.Value.(*CacheEntry).value, true
	}

	return nil, false
}

// добавление или обновление пары ключ-значение в кеше
func (lru *LRUCache) Set(key string, value interface{}) {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()

	// проверка того, существует ли ключ уже в кэше
	if element, ok := lru.cache[key]; ok {
		// обновляем существующую запись и перемещаем ее в начало (использовалась в последний раз)
		element.Value.(*CacheEntry).value = value
		lru.list.MoveToFront(element)
	} else {
		// добавление новой записи в кеш
		entry := &CacheEntry{key: key, value: value}
		element := lru.list.PushFront(entry)
		lru.cache[key] = element

		// проверяем, заполнен ли кеш (есть ли место), при необходимости удаляем наименее недавно использованный элемент
		if lru.list.Len() > lru.capacity {
			oldest := lru.list.Back()
			if oldest != nil {
				delete(lru.cache, oldest.Value.(*CacheEntry).key)
				lru.list.Remove(oldest)
			}
		}
	}
}

// PrintCache, который печатает текущее содержимое кеша
func (lru *LRUCache) PrintCache() {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()

	fmt.Printf("LRU Cache (Capacity: %d, Size: %d): [", lru.capacity, lru.list.Len())
	for element := lru.list.Front(); element != nil; element = element.Next() {
		entry := element.Value.(*CacheEntry)
		fmt.Printf("(%s: %v) ", entry.key, entry.value)
	}
	fmt.Println("]")
}

func main() {
	// создание кэша LRU ёмкостью 3
	lruCache := NewLRUCache(3)

	// задаём пары ключ-занчение
	lruCache.Set("company", "Yandex")
	lruCache.Set("division", "Yandex Lyceum")
	lruCache.Set("course", "Golang")

	lruCache.PrintCache()

	if value, ok := lruCache.Get("company"); ok {
		fmt.Println("Value for company:", value)
	} else {
		fmt.Println("company not found in the cache.")
	}

	// установка дополнительных пар ключ-значение для инициирования вытеснения
	lruCache.Set("year", 2024)
	lruCache.Set("age", "13-17yrs")

	lruCache.PrintCache()
}*/

// --------------------
// создание записи кеша, представляющей запись в TTLCache со значением и временем истечения срока действия
/*type CacheEntry struct {
	value      interface{} // значение, связанное с записью в кеше
	expiration time.Time   // время истечения срока действия записи в кеше
}*/

// создание TTLCache в качестве потокобезопасного кеша с функциональностью time-to-live
/*type TTLCache struct {
	data  map[string]CacheEntry // мапа для хранения пар ключ-значение со временем истечения срока действия
	mutex sync.RWMutex          // мьютекс для синхронизации конкурентного доступа к кешу
}

// создание NewTTLCache нового экземпляра TTLCache с инициализированной картой данных
func NewTTLCache() *TTLCache {
	return &TTLCache{
		data: make(map[string]CacheEntry),
	}
}

// извлечение значения, связанного с данным ключом, из кеша
// Get() возвращает значение и признак, указывающий, был ли ключ найден и не истёк ли срок его действия
func (c *TTLCache) Get(key string) (interface{}, bool) {
	c.mutex.RLock()         // получение блокировки чтения, чтобы разрешить одновременное чтение нескольким устройствам чтения
	defer c.mutex.RUnlock() // снятие блокировки чтения, когда функция завершит работу
	entry, ok := c.data[key]
	if !ok || time.Now().After(entry.expiration) {
		// если ключ не найден или срок действия истёк, возвращаем nil и false
		return nil, false
	}
	return entry.value, true
}

// установка значения, связанного с данным ключом в кеше
// Set() получает блокировку на запись для обеспечения эксклюзивного доступа во время обновления
func (c *TTLCache) Set(key string, value interface{}, ttl time.Duration) {
	c.mutex.Lock()         // получение блокировки на запись для эксклюзивного доступа
	defer c.mutex.Unlock() // снятие блокировки записи при завершении работы функции
	c.data[key] = CacheEntry{
		value:      value,
		expiration: time.Now().Add(ttl),
	}
}

func main() {
	ttlCache := NewTTLCache()

	// установка пар ключ-значение в кеше TTL с разным временем
	ttlCache.Set("company", "yandex", 5*time.Second) // ttl 5 seconds
	ttlCache.Set("year", 2024, 10*time.Second)       //ttl 10 seconds

	// извлечение значений из кеша TTL
	if value, ok := ttlCache.Get("company"); ok {
		fmt.Println("Value for company:", value)
	} else {
		fmt.Println("company not found in the TTL cache.")
	}

	if value, ok := ttlCache.Get("year"); ok {
		fmt.Println("Value for year:", value)
	} else {
		fmt.Println("year not found in the TTL cache.")
	}

	// создание ожидания с помощью функции Sleep() для имитации использования кеша
	// изменяем время ожидания на другое количество секунд, чтобы увидеть эффект
	time.Sleep(7 * time.Second)

	// повторное извлечение значений через некоторое время
	if value, ok := ttlCache.Get("company"); ok {
		fmt.Println("Value for company (after some time):", value)
	} else {
		fmt.Println("company not found in the TTL cache after some time.")
	}

	if value, ok := ttlCache.Get("year"); ok {
		fmt.Println("Value for year (after some time):", value)
	} else {
		fmt.Println("year not found in the TTL cache after some time.")
	}
}*/

type CacheEntry struct {
	value    interface{}
	expireAt int64
}

func NewCacheEntry(value interface{}, expireAt int64) CacheEntry {
	return CacheEntry{
		value:    value,
		expireAt: expireAt,
	}
}

func (ce CacheEntry) IsExpired() bool {
	return ce.expireAt < time.Now().UnixNano()
}

type Cache struct {
	kvstore  map[string]CacheEntry
	locker   sync.RWMutex
	interval time.Duration
	stop     chan struct{}
}

func NewCache(cleanUpInterval time.Duration) *Cache {
	cache := &Cache{
		kvstore:  make(map[string]CacheEntry),
		interval: cleanUpInterval,
		stop:     make(chan struct{}),
	}

	if cleanUpInterval > 0 {
		go cache.cleaning()
	}
	return cache
}

func (c *Cache) cleaning() {
	fmt.Println("cleaner starting...")
	ticker := time.NewTicker(c.interval)
	fmt.Println("cleaner was started")
	for {
		select {
		case <-ticker.C:
			c.purge()
		case <-c.stop:
			ticker.Stop()
			fmt.Println("cleaner was stopped")
			return
		}
	}
}

func (c *Cache) purge() {
	c.locker.Lock()
	defer c.locker.Unlock()
	for key, data := range c.kvstore {
		if data.IsExpired() {
			delete(c.kvstore, key)
		}
	}
}

func (c *Cache) Set(key string, value interface{}, expiryDuration time.Duration) {
	expireAt := time.Now().Add(expiryDuration).UnixNano()
	c.locker.Lock()
	defer c.locker.Unlock()
	c.kvstore[key] = NewCacheEntry(value, expireAt)

}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.locker.RLock()
	defer c.locker.RUnlock()
	data, found := c.kvstore[key]
	if !found || data.IsExpired() {
		return nil, false
	}

	return data.value, true
}

func (c *Cache) Close() {
	close(c.stop)
}

func main() {
	cache := NewCache(time.Second)
	defer cache.Close()
	cache.Set("foo", "bar", 2*time.Second)
	for i := 0; i < 3; i++ {
		value, found := cache.Get("foo")
		if found {
			fmt.Println("value for key foo is ", value)
		} else {
			fmt.Println("value for key foo is not found")
			break
		}

		fmt.Println("waiting for 1 second...")
		time.Sleep(time.Second)
	}
}
