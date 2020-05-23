package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"

	_ "github.com/lib/pq"
)

const (
	invalidData = "datastore: invalid %s found : '%v'"

	// noReturnVal indicates the sql query being executed should not
	// return any value. Return value not expected.
	noReturnVal int = iota

	// singleRow indicates the sql query bieng executed should only
	// return a single row of the expected result set.
	singleRow

	// multiRows indicates the sql query being executed should return
	// multiple rows of the expected result set.
	multiRows
)

type DBConnConfig struct {
	Host         string
	Port         int
	DBName         string
	User         string
	Password         string
	RootCert       string
	ClientCert string
	ClientKey string
}

func (db DBConnConfig) String() string {
	return fmt.Sprintf(""+
		"user=%s "+
		"password=%s "+
		"host=%s "+
		"port=%d "+
		"dbname=%s "+
		"sslcert=%s " +
		"sslkey=%s " +
		"rootcert=%s " +
		"", db.User, db.Password, db.Host, db.Port, db.DBName, db.ClientCert, db.ClientKey, db.RootCert)
}


type Database struct {
	dbConn *sql.DB
	gormConn *gorm.DB
}

func (db Database) ExecPrepStmts(queryType int, sqlQuery string, val ...interface{}) (*sql.Rows, *sql.Row, error) {
	stmt, err := db.dbConn.Prepare(sqlQuery)
	if err != nil {
		return nil, nil, err
	}

	defer stmt.Close()

	switch queryType {
	case noReturnVal:
		_, err = db.dbConn.Exec(sqlQuery, val...)
		return nil, nil, err

	case singleRow:
		row := db.dbConn.QueryRow(sqlQuery, val...)
		return nil, row, nil

	case multiRows:
		rows, err := db.dbConn.Query(sqlQuery, val...)
		return rows, nil, err

	default:
		return nil, nil,
			fmt.Errorf(invalidData, "queryType", queryType)
	}
}

func NewConnection() (*sql.DB, error) {
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

	conn, err := sql.Open("postgres", cfg.String())
	if err != nil {
		return nil, err
	}

	// ping db to check if alive
	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func NewDatabase(connection *sql.DB, gormConnection *gorm.DB) *Database {
	return &Database{dbConn: connection, gormConn: gormConnection}
}