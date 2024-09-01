package handlers

import "github.com/sayed-imran/go-design-pattern/db"

type Repository struct {
	DB db.MongodbRepo
}

var Repo *Repository

func NewRepository(db db.MongodbRepo) *Repository {
	return &Repository{
		DB: db,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}
