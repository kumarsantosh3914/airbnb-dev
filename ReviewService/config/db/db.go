package config

import (
	env "ReviewService/config/env"
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func SetupDB() (*sql.DB, error) {
	cfg := mysql.NewConfig()

	cfg.User = env.GetString("DB_USER", "root")
	cfg.Passwd = env.GetString("DB_PASSWORD", "root")
	cfg.Net = env.GetString("DB_NET", "tcp")
	cfg.Addr = env.GetString("DB_ADDR", "127.0.0.1:3306")
	cfg.DBName = env.GetString("DBName", "airbnb_dev")

	fmt.Println("Connecting to database:", cfg.DBName)

	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		fmt.Println("Error opening database:", err)
	}

	fmt.Println("Trying to connect database...")
	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println("Error connecting to database:", pingErr)
		return nil, pingErr
	}

	fmt.Println("Connected to database successfully:", cfg.DBName)

	return db, nil
}
