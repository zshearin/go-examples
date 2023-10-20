package mysql

import (
	"database/sql"

	"github.com/zshearin/go-examples/sql/entity"
	"github.com/zshearin/go-examples/sql/repository"
)

type entityRepository struct {
	db *sql.DB
}

func NewEntityRepository(db *sql.DB) repository.EntityRepository {
	return &entityRepository{db: db}
}

func (e *entityRepository) GetThing() (*entity.Thing, error) {

	return &entity.Thing{
		Field1: "this is the value",
	}, nil
}
