package repository

import "github.com/zshearin/go-examples/sql/entity"

type EntityRepository interface {
	GetThing() (*entity.Thing, error)
}
