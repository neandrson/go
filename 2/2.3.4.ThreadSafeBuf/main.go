package main

import "sync"

type Bufer struct {
	Buf []int
	mu  sync.Mutex
}

func (b *Bufer) Write(num int) {

}

func (b *Bufer) Consume() int {

}

func main() {

}
