package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

type User struct {
	ID       int
	Username string
	Email    string
}

type DB struct {
	conn *sql.DB
}

func NewDB(dataSourceName string) (*DB, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	return &DB{conn: db}, nil
}

func (d *DB) Close() error {
	return d.conn.Close()
}

func (d *DB) InsertUser(user User) error {
	stmt, err := d.conn.Prepare("INSERT INTO users (username, email) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username, user.Email)
	if err != nil {
		return err
	}

	return nil
}

func (d *DB) DeleteUser(userID int) error {
	stmt, err := d.conn.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID)
	if err != nil {
		return err
	}

	return nil
}

func (d *DB) UpdateUser(user User) error {
	stmt, err := d.conn.Prepare("UPDATE users SET username = ?, email = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username, user.Email, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	db, err := NewDB("user:password@tcp(host:port)/database")
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}
	defer db.Close()

	// Example usage
	user := User{
		Username: "john",
		Email:    "john@example.com",
	}

	err = db.InsertUser(user)
	if err != nil {
		fmt.Println("Failed to insert user:", err)
		return
	}

	// Update user
	user.Username = "john_doe"
	err = db.UpdateUser(user)
	if err != nil {
		fmt.Println("Failed to update user:", err)
		return
	}

	// Delete user
	err = db.DeleteUser(user.ID)
	if err != nil {
		fmt.Println("Failed to delete user:", err)
		return
	}
}
