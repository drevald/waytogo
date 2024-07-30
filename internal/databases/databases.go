package databases

import (
	"fmt"
	"github.com/ddreval/waytogo/internal/config"
	"github.com/samber/do"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/xo/dburl"
)

func New(di *do.Injector) (*gorm.DB, error) {
	cnf, err := do.Invoke[*config.Config](di)
	if err != nil {
		fmt.Println(fmt.Scanf("!!!!!!!!!!!!!!!!!  config error %s", err.Error()))
		return nil, err
	}
	url, err := dburl.Parse(cnf.DbUrl)
	fmt.Println(fmt.Scanf("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$ cnf.DbUrl = %s", cnf.DbUrl));
	if err != nil {
		fmt.Println(fmt.Scanf("#################### url error %s", err.Error()))
		return nil, err
	}
	fmt.Println(fmt.Scanf("@@@@@@@@@@@@@@@@@@@@@@@@@ url.DSN = %s", url.DSN));
	db, err := gorm.Open(postgres.Open(url.DSN), &gorm.Config{})
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
