package common

import "github.com/dpolimeni/fiber_app/ent"

type Storage interface {
	DataBase() (*ent.Client, error)
}
