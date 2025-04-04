package main

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

/*var counter int32
var wg sync.WaitGroup

func main() {
	// Увеличиваем значение counter с помощью атомарной операции AddInt32
	wg.Add(2)
	go increment()
	go increment()
	wg.Wait()
	fmt.Println("Counter:", counter)
}

func increment() {
	defer wg.Done()
	atomic.AddInt32(&counter, 1)
}*/

/*type Person struct {
	Name string
	Age  int
}*/

/*func main() {
	person := &Person{Name: "Alice", Age: 25}

	// Загружаем указатель на структуру Person
	ptr := unsafe.Pointer(person)
	loadedPtr := atomic.LoadPointer(&ptr)

	// Преобразуем указатель обратно в структуру Person
	loadedPerson := (*Person)(loadedPtr)

	fmt.Println(loadedPerson.Name, loadedPerson.Age)
}*/

/*func main() {
	person := &Person{Name: "Alice", Age: 25}

	// Загружаем указатель на структуру Person
	ptr := unsafe.Pointer(person)

	// Ожидаемое значение указателя
	expectedPtr := ptr

	// Новое значение указателя
	newPtr := unsafe.Pointer(&Person{Name: "Bob", Age: 30})

	// Атомарно сравниваем и заменяем указатель
	swapped := atomic.CompareAndSwapPointer(&ptr, expectedPtr, newPtr)

	if swapped {
		fmt.Println("Pointer was successfully swapped")
	} else {
		fmt.Println("Pointer was not swapped")
	}
}*/

// Node представляет узел в связанном списке.
/*type Node struct {
	value int            // Значение узла
	next  unsafe.Pointer // Указатель на следующий узел в списке
}*/

// List представляет связанный список.
/*type List struct {
	head unsafe.Pointer // Указатель на голову списка
}*/

// Add добавляет новый узел с заданным значением в начало списка.
/*func (l *List) Add(value int) {
	node := &Node{value: value}

	for {
		// Загружаем текущую голову списка
		oldHead := atomic.LoadPointer(&l.head)
		// Устанавливаем новый узел как следующий для добавляемого узла
		node.next = oldHead

		// Если голову можно атомарно заменить на новый узел, то прерываем цикл
		if atomic.CompareAndSwapPointer(&l.head, oldHead, unsafe.Pointer(node)) {
			break
		}
	}
}*/

// Print выводит значения узлов списка.
/*func (l *List) Print() {
	// Загружаем голову списка
	curr := atomic.LoadPointer(&l.head)

	// Проходим по всем узлам в списке
	for curr != nil {
		// Преобразуем указатель в структуру Node
		node := (*Node)(curr)
		// Выводим значение узла
		fmt.Println(node.value)
		// Загружаем указатель на следующий узел
		curr = atomic.LoadPointer(&node.next)
	}
}

func main() {
	// Создаем новый связанный список
	list := &List{}
	// Добавляем узлы со значениями 1, 2, 3 в начало списка
	list.Add(1)
	list.Add(2)
	list.Add(3)
	// Выводим значения узлов списка
	list.Print()
}*/

// Node представляет узел в стеке.
type Node struct {
	value int            // Значение узла
	next  unsafe.Pointer // Указатель на следующий узел в стеке
}

// Stack представляет lock-free стек.
type Stack struct {
	top unsafe.Pointer // Указатель на вершину стека
}

// Push добавляет новый элемент на вершину стека.
func (s *Stack) Push(value int) {
	node := &Node{value: value}

	for {
		// Загружаем текущую вершину стека
		oldTop := atomic.LoadPointer(&s.top)
		// Устанавливаем новый узел как следующий для добавляемого узла
		node.next = oldTop

		// Если вершину можно атомарно заменить на новый узел, то прерываем цикл
		if atomic.CompareAndSwapPointer(&s.top, oldTop, unsafe.Pointer(node)) {
			break
		}
	}
}

// Pop удаляет и возвращает элемент с вершины стека. Если стек пуст, возвращает false.
func (s *Stack) Pop() (int, bool) {
	for {
		// Загружаем текущую вершину стека
		oldTop := atomic.LoadPointer(&s.top)
		// Если стек пуст, возвращаем false
		if oldTop == nil {
			return 0, false
		}

		// Загружаем указатель на следующий узел
		newTop := (*Node)(oldTop).next

		// Выполняем атомарную попытку заменить текущую вершину на следующий узел
		if atomic.CompareAndSwapPointer(&s.top, oldTop, unsafe.Pointer(newTop)) {
			// Возвращаем значение удалённого узла
			return (*Node)(oldTop).value, true
		}
	}
}

func main() {
	// Создаём новый стек
	stack := &Stack{}

	// Добавляем элементы на вершину стека
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// Удаляем элементы с вершины стека и выводим их значения
	for {
		if value, ok := stack.Pop(); ok {
			fmt.Println(value)
		} else {
			break
		}
	}
}
