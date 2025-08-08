package db

import (
	"SantiagoBobrik/iot-home/domain"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	conn *sql.DB
}

// New initializes and returns a new DB instance.
func New(dbPath string) (*DB, error) {
	conn, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Optional: test the connection
	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return &DB{conn: conn}, nil
}

// Close the DB connection.
func (d *DB) Close() error {
	return d.conn.Close()
}

// InitSchema creates the schema if it doesn't exist.
func (d *DB) InitSchema() error {
	query := `
	CREATE TABLE IF NOT EXISTS data (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		device_id text,
		temperature integer NOT NULL,
		humidity integer NOT NULL,
        created_at DATETIME NOT NULL
	);`
	_, err := d.conn.Exec(query)
	return err
}

// InsertUser inserts a new user into the DB.
func (d *DB) InsertData(data domain.Data) error {
	_, err := d.conn.Exec("INSERT INTO data (temperature, humidity, created_at, device_id) VALUES (?, ?, ?, ?)", data.Temperature, data.Humidity, data.CreatedAt, data.DeviceID)
	return err
}

// GetUsers returns all users.
func (d *DB) GetData() ([]domain.Data, error) {
	rows, err := d.conn.Query("SELECT id, device_id, temperature, humidity, created_at FROM data ORDER BY created_at DESC")
	if err != nil {
		return []domain.Data{}, err
	}
	defer rows.Close()
	data := []domain.Data{}

	for rows.Next() {
		var u domain.Data
		if err := rows.Scan(&u.ID, &u.DeviceID, &u.Temperature, &u.Humidity, &u.CreatedAt); err != nil {
			return nil, err
		}
		data = append(data, u)
	}
	return data, nil
}
