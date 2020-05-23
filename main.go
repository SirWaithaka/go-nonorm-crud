package main

import (
	"log"
	"os"
)

func main() {
	conn, err := NewConnection()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	gConn, err := NewGormConn()
	db := NewDatabase(conn, gConn)

	repo := NewRepository(db)
	u, err := repo.FindUserByEmail("admin@tospay.net")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println(u)

	user, err := repo.FindUserByUsername("admin")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println(user)
}
