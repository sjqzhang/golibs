package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sjqzhang/golibs/dsnparse"
	"sync"
)

var globalGormV1 *gorm.DB
var once sync.Once

func InitGlobalGormDB(dsn string) (*gorm.DB, error) {
	var err error
	once.Do(func() {
		globalGormV1, err = NewGormV1(dsn)

	})
	if err != nil {
		return nil, err
	}
	return globalGormV1, nil
}

func NewGormV1(dsn string) (*gorm.DB, error) {
	dp, err := dsnparse.Parse(dsn)
	if err != nil {
		return nil, err
	}
	var db *gorm.DB
	if dp.Scheme() == "sqlite3" {
		db, err = gorm.Open(dp.Scheme(), dp.Host())
	} else {
		db, err = gorm.Open(dp.Scheme(), fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dp.Username(), dp.Password(), dp.Host(), dp.Port(), dp.DatabaseName()))
	}
	if err != nil {
		return nil, err
	}
	db.DB().SetMaxIdleConns(10)
	db.SingularTable(true)
	return db, nil
}
