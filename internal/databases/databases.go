package databases

import (
	"github.com/ddreval/waytogo/internal/config"
	"github.com/samber/do"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/xo/dburl"
	"github.com/sirupsen/logrus"
)

func New(di *do.Injector) (*gorm.DB, error) {	
	logger, err := do.Invoke[*logrus.Logger](di)
	if err != nil {
		return nil, err
	}
	cnf, err := do.Invoke[*config.Config](di)
	if err != nil {
		return nil, err
	}
	url, err := dburl.Parse(cnf.DbUrl)
	if err != nil {
		logger.Errorf("Url error %s", err)
		return nil, err
	}
	db, err := gorm.Open(postgres.Open(url.DSN), &gorm.Config{})
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	err = db.AutoMigrate(&User{})
    if err != nil {
        logger.Errorf("failed to migrate database: %v", err)
    }
	return db, nil
}
