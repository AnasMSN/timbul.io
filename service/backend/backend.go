package backend

import (
	"timbul.io/repository"
)

type Service struct {
	DB repository.Database
}

func New(db repository.Database) Service {
	return Service{DB: db}
}
