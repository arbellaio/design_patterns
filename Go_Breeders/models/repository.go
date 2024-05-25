package models

import "database/sql"

type Repository interface {
	AllDogBreeds() ([]*DogBreed, error)
}

type mySqlRepository struct {
	DB *sql.DB
}

func newMySqlRepository(conn *sql.DB) Repository {
	return &mySqlRepository{
		DB: conn,
	}
}
