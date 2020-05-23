package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// NewDatabase creates a new Database object
func NewGormConn() (*gorm.DB, error) {
	cfg := DBConnConfig{
		Host:       "localhost",
		Port:       26257,
		DBName:     "test_db",
		User:       "waithaka",
		Password:   "kennedy",
		RootCert:   "/home/waithaka/.cockroach/certs/ca.crt",
		ClientCert: "/home/waithaka/.cockroach/certs/client.waithaka.crt",
		ClientKey:  "/home/waithaka/.cockroach/certs/client.waithaka.key",
	}

	var err error

	var conn *gorm.DB
	conn, err = gorm.Open("postgres", cfg.String())

	if err != nil {
		return nil, err
	}

	return conn, nil
}
