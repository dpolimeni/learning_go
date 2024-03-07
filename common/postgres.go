package common

import (
	"fmt"

	"github.com/dpolimeni/fiber_app/ent"
)

type Postgres struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

func (p *Postgres) DataBase() (*ent.Client, error) {
	connection := fmt.Sprintf(
		"host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		p.Host, p.Port, p.User, p.Password, p.Dbname)
	client, err := ent.Open("postgres", connection)
	return client, err
}
