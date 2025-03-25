package main

import (
	"html/template"
	"os"
)

//func myFunc(a interface{}) {
// Пример 1 - Nil interface
/*s, ok := a.(string)
if ok {
	fmt.Printf("'%v' is a string\n", s)
} else {
	fmt.Printf("'%v' is not a string\n", a)
}*/

// Пример 2 - Рефлексия
/*t := reflect.TypeOf(a)
fmt.Printf("Type of '%v' is %v\n", a, t)*/
//}

// Пример 3 - Получение значения поля структуры
/*type Person struct {
	Name string
	Age  int
}*/

//func main() {
// Пример 1, 2
/*myFunc("hello")
myFunc(42)*/

// Пример 3 - Получение значения поля структуры
/*p := Person{Name: "John", Age: 30}
v := reflect.ValueOf(p)
name := v.FieldByName("Name").String()
age := v.FieldByName("Age").Int()
fmt.Println(name, age)*/

// Пример 4 - Изменение значения поля структуры
/*p := Person{Name: "John", Age: 30}
v := reflect.ValueOf(&p).Elem()
v.FieldByName("Name").SetString("Jane")
v.FieldByName("Age").SetInt(25)
fmt.Println(p)*/
//}

// Пример 5 - Кодогенерация — это генерация кода на основе шаблонов и данных
type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

func main() {
	t := template.Must(template.ParseFiles("template.go.tmpl"))
	data := struct {
		Package string
		Struct  string
		User    string
	}{
		Package: "db",
		Struct:  "UserRepository",
		User:    "User",
	}
	t.Execute(os.Stdout, data)
}
