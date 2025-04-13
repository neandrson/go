package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
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
}
