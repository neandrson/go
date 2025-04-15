package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

/*type User struct {
	ID       int64
	Name     string
	Password string
}

func (u User) ComparePassword(u2 User) error {
	if u.Password == u2.Password {
		log.Println("auth success")
		return nil
	}
	log.Println("auth fail")
	return fmt.Errorf("passwords don't match")
}

func createTable(ctx context.Context, db *sql.DB) error {
	const usersTable = `
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT UNIQUE,
		password TEXT
	);`

	if _, err := db.ExecContext(ctx, usersTable); err != nil {
		return err
	}

	return nil
}

func insertUser(ctx context.Context, db *sql.DB, user *User) (int64, error) {
	var q = `
	INSERT INTO users (name, password) values ($1, $2)
	`
	result, err := db.ExecContext(ctx, q, user.Name, user.Password)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func selectUser(ctx context.Context, db *sql.DB, name string) (User, error) {
	var (
		user User
		err  error
	)

	var q = "SELECT id, name, password FROM users WHERE name=$1"
	err = db.QueryRowContext(ctx, q, name).Scan(&user.ID, &user.Name, &user.Password)
	return user, err
}

// Криптографические хеш-функции
// Наряду с каждым хешем есть случайная строка байтов, известная, как соль.
func generate(s string) (string, error) {
	saltedBytes := []byte(s)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hash := string(hashedBytes[:])
	return hash, nil
}

// Напишем функцию сравнения хэша с паролем.
func compare(hash string, s string) error {
	incoming := []byte(s)
	existing := []byte(hash)
	return bcrypt.CompareHashAndPassword(existing, incoming)
}

func main() {
	ctx := context.TODO()

	db, err := sql.Open("sqlite3", "store.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.PingContext(ctx)
	if err != nil {
		panic(err)
	}

	if err = createTable(ctx, db); err != nil {
		panic(err)
	}

	user := &User{
		Name:     "Name",
		Password: "qwertyqwerty",
	}
	userID, err := insertUser(ctx, db, user)
	if err != nil {
		log.Println("user already exists")
	} else {
		user.ID = userID
	}

	userFromDB, err := selectUser(ctx, db, user.Name)
	if err != nil {
		panic(err)
	}

	user.ComparePassword(userFromDB)
	user.Password = "fail passsword"
	user.ComparePassword(userFromDB)
}*/

// Итоговая
type User struct {
	ID             int64
	Name           string
	Password       string
	OriginPassword string
}

func (u User) ComparePassword(u2 User) error {
	err := compare(u2.Password, u.OriginPassword)
	if err != nil {
		log.Println("auth fail")
		return err
	}

	log.Println("auth success")
	return nil
}

func createTable(ctx context.Context, db *sql.DB) error {
	const usersTable = `
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		name TEXT UNIQUE,
		password TEXT
	);`

	if _, err := db.ExecContext(ctx, usersTable); err != nil {
		return err
	}

	return nil
}

func insertUser(ctx context.Context, db *sql.DB, user *User) (int64, error) {
	var q = `
	INSERT INTO users (name, password) values ($1, $2)
	`
	result, err := db.ExecContext(ctx, q, user.Name, user.Password)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func selectUser(ctx context.Context, db *sql.DB, name string) (User, error) {
	var (
		user User
		err  error
	)

	var q = "SELECT id, name, password FROM users WHERE name=$1"
	err = db.QueryRowContext(ctx, q, name).Scan(&user.ID, &user.Name, &user.Password)
	return user, err
}

func generate(s string) (string, error) {
	saltedBytes := []byte(s)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hash := string(hashedBytes[:])
	return hash, nil
}

func compare(hash string, s string) error {
	incoming := []byte(s)
	existing := []byte(hash)
	return bcrypt.CompareHashAndPassword(existing, incoming)
}

func main() {
	ctx := context.TODO()

	db, err := sql.Open("sqlite3", "store.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.PingContext(ctx)
	if err != nil {
		panic(err)
	}

	if err = createTable(ctx, db); err != nil {
		panic(err)
	}

	password, err := generate("qwertyqwerty")
	if err != nil {
		panic(err)
	}

	user := &User{
		Name:           "New Name",
		Password:       password,
		OriginPassword: "qwertyqwerty",
	}
	userID, err := insertUser(ctx, db, user)
	if err != nil {
		log.Println("user already exists")
	} else {
		user.ID = userID
	}

	userFromDB, err := selectUser(ctx, db, user.Name)
	if err != nil {
		panic(err)
	}

	user.ComparePassword(userFromDB)
	user.Password, err = generate("fail passsword")
	if err != nil {
		panic(err)
	}
	user.OriginPassword = "fail passsword"
	user.ComparePassword(userFromDB)
}
