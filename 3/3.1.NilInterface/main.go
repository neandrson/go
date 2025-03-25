package main

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
/*type User struct {
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
}*/

// Пример 6 - Дженерики позволяют писать код, который сможет работать с различными типами данных
// Repository — обобщённый тип для работы с базой данных
/*type Repository[T any] struct {
	db *sqlx.DB
}

// NewRepository — создаёт новый экземпляр репозитория
func NewRepository[T any](db *sqlx.DB) *Repository[T] {
	return &Repository[T]{
		db: db,
	}
}

// Add — добавляет запись в базу данных
func (r *Repository[T]) Add(entity T, ctx context.Context) error {
	_, err := r.db.NamedExecContext(ctx, "INSERT INTO your_table_name_here (field1, field2) VALUES (:field1, :field2)", entity)
	return err
}

// GetById — получает запись из базы данных по идентификатору
func (r *Repository[T]) GetById(id int, ctx context.Context) (T, error) {
	var entity T
	err := r.db.GetContext(ctx, &entity, "SELECT * FROM your_table_name_here WHERE id = ? AND is_active = ?", id, true)
	if err != nil {
		return entity, err
	}

	return entity, nil
}

// Get — выполняет поиск записи в базе данных по параметрам
func (r *Repository[T]) Get(params T, ctx context.Context) T {
	var entity T
	r.db.GetContext(ctx, &entity, "SELECT * FROM your_table_name_here WHERE field1 = :field1 AND field2 = :field2", params)
	return entity
}

// GetAll — получает все записи из базы данных
func (r *Repository[T]) GetAll(ctx context.Context) ([]T, error) {
	var entities []T
	err := r.db.SelectContext(ctx, &entities, "SELECT * FROM your_table_name_here")
	if err != nil {
		return entities, err
	}
	return entities, nil
}

// Update — обновляет запись в базе данных
func (r *Repository[T]) Update(entity T, ctx context.Context) error {
	_, err := r.db.NamedExecContext(ctx, "UPDATE your_table_name_here SET field1 = :field1, field2 = :field2 WHERE id = :id", entity)
	return err
}*/
