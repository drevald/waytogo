package databases

import (
	"github.com/samber/do"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/ddreval/waytogo/internal/config"
	"fmt"
)

type Database struct {
	db *gorm.DB
}

func New(di *do.Injector) (*Database, error) {    
	cnf, err := do.Invoke[*config.Config](di)
	if err != nil {
		return nil, err
	}
	db, err := gorm.Open(postgres.Open(cnf.DbUrl), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &Database {db}, nil
}