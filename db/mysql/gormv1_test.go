package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestInitGlobalGormDB(t *testing.T) {
	dsn := "mysql://root:mock@tcp(localhost:63307)/mock?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := InitGlobalGormDB(dsn)
	if err != nil {
		t.Fail()
	}
	if db == nil {
		t.Fail()
	}
}
