package main

import (
	"log"
	"sync"
)

type SafeMap struct {
	m   map[string]interface{}
	mux sync.Mutex
}

func (s *SafeMap) Get(key string) interface{} {
	s.mux.Lock()
	data, exists := s.m[key] // теперь доступ к мапе внутри критической секции
	s.mux.Unlock()
	if exists {
		return *m.s
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
	return *data, true
}

func (s *SafeMap) Set(key string, value interface{}) {

}

func NewSafeMap() *SafeMap {

}

func main() {

}
