package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"order-service/internal/models"

	_ "github.com/lib/pq"
)

type DB struct {
	conn *sql.DB
}

func NewDB(host, port, user, password, dbname string) (*DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	db := &DB{conn: conn}
	if err := db.createTables(); err != nil {
		return nil, err
	}

	return db, nil
}

func (db *DB) createTables() error {
	query := `
	CREATE TABLE IF NOT EXISTS orders (
		order_uid VARCHAR(255) PRIMARY KEY,
		data JSONB NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`

	_, err := db.conn.Exec(query)
	return err
}

func (db *DB) SaveOrder(order *models.Order) error {
	data, err := json.Marshal(order)
	if err != nil {
		return err
	}

	query := `INSERT INTO orders (order_uid, data) VALUES ($1, $2) 
              ON CONFLICT (order_uid) DO UPDATE SET data = $2`

	_, err = db.conn.Exec(query, order.OrderUID, data)
	return err
}

func (db *DB) GetOrder(orderUID string) (*models.Order, error) {
	var data []byte
	query := `SELECT data FROM orders WHERE order_uid = $1`

	err := db.conn.QueryRow(query, orderUID).Scan(&data)
	if err != nil {
		return nil, err
	}

	var order models.Order
	if err := json.Unmarshal(data, &order); err != nil {
		return nil, err
	}

	return &order, nil
}

func (db *DB) GetAllOrders() (map[string]*models.Order, error) {
	rows, err := db.conn.Query(`SELECT data FROM orders`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := make(map[string]*models.Order)
	for rows.Next() {
		var data []byte
		if err := rows.Scan(&data); err != nil {
			continue
		}

		var order models.Order
		if err := json.Unmarshal(data, &order); err != nil {
			continue
		}

		orders[order.OrderUID] = &order
	}

	return orders, nil
}

func (db *DB) Close() error {
	return db.conn.Close()
}
