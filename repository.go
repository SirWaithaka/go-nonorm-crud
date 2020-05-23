package main

import "database/sql"

type Repository interface {
	FindUserByEmail(string) (User, error)
	FindUserByUsername(string) (User, error)
}

func NewRepository(db *Database) Repository {
	return &repository{db}
}

type repository struct {
	db *Database
}

type ErrNoRecordFound struct {}
func (ErrNoRecordFound) Error() string {
	return "no record found"
}

func (repo repository) FindUserByEmail(email string) (User, error) {

	query := `SELECT id, email, password FROM users WHERE email = $1`
	_, row, err := repo.db.ExecPrepStmts(singleRow, query, email)

	if err != nil {
		return User{}, err
	}

	var user User
	err = row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil && err != sql.ErrNoRows {
		return User{}, err
	}

	if err == sql.ErrNoRows {
		return User{}, ErrNoRecordFound{}
	}

	return user, nil
}

func (repo repository) FindUserByUsername(username string) (User, error) {
	var user User
	result := repo.db.gormConn.Where(User{Username: username}).First(&user)

	if result.RecordNotFound() {
		return User{}, ErrNoRecordFound{}
	}

	if err := result.Error; err != nil {
		return User{}, nil
	}

	return user, nil
}