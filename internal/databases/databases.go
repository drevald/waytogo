package databases

import (
	"fmt"
	"github.com/ddreval/waytogo/internal/config"
	"github.com/samber/do"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(di *do.Injector) (*gorm.DB, error) {
	cnf, err := do.Invoke[*config.Config](di)
	if err != nil {
		return nil, err
	}
	db, err := gorm.Open(postgres.Open(cnf.DbUrl), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	err = db.AutoMigrate(&User{})
    if err != nil {
        fmt.Printf("failed to migrate database: %v", err)
    }
	return db, nil
}
