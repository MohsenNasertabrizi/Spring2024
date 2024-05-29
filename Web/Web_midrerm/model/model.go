package model

import (
	"database/sql"
	"log"
	"time"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type Basket struct {
	ID        int       `json:"id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func InitDB(filepath string) {
	var err error
	db, err = sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatal(err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS baskets (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		status TEXT,
		created_at DATETIME,
		updated_at DATETIME
	);
	`

	if _, err := db.Exec(createTable); err != nil {
		log.Fatal(err)
	}
}

func CloseDB() {
	db.Close()
}

func GetAllBaskets() ([]Basket, error) {
	rows, err := db.Query("SELECT id, status, created_at, updated_at FROM baskets")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var baskets []Basket
	for rows.Next() {
		var basket Basket
		if err := rows.Scan(&basket.ID, &basket.Status, &basket.CreatedAt, &basket.UpdatedAt); err != nil {
			return nil, err
		}
		baskets = append(baskets, basket)
	}
	return baskets, nil
}

func GetBasketByID(id int) (*Basket, error) {
	row := db.QueryRow("SELECT id, status, created_at, updated_at FROM baskets WHERE id = ?", id)

	var basket Basket
	if err := row.Scan(&basket.ID, &basket.Status, &basket.CreatedAt, &basket.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &basket, nil
}

func CreateBasket(basket *Basket) error {
	now := time.Now()
	res, err := db.Exec("INSERT INTO baskets (status, created_at, updated_at) VALUES (?, ?, ?)", basket.Status, now, now)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	basket.ID = int(id)
	basket.CreatedAt = now
	basket.UpdatedAt = now
	return nil
}

func UpdateBasket(basket *Basket) error {
	now := time.Now()
	_, err := db.Exec("UPDATE baskets SET status = ?, updated_at = ? WHERE id = ? AND status != 'COMPLETED'", basket.Status, now, basket.ID)
	if err != nil {
		return err
	}
	basket.UpdatedAt = now
	return nil
}

func DeleteBasket(id int) error {
	_, err := db.Exec("DELETE FROM baskets WHERE id = ?", id)
	return err
}
