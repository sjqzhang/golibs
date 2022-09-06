package dsnparse

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	d, e := Parse("mysql://root:123456@x@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if e != nil {
		t.Fail()
	}
	if d.User().Username() != "root" {
		t.Fail()
	}
	if v, ok := d.User().Password(); v != "123456@x" || !ok {
		t.Fail()
	}
	if d.Port() != "3306" {
		t.Fail()
	}
	fmt.Println(d.Host())
	if !d.GetBool("parseTime", false) {
		t.Fail()
	}

}

func TestParseSQlite3(t *testing.T) {
	d, e := Parse("sqlite3://test.db")
	if e != nil {
		t.Fail()
	}

	if d.Scheme() != "sqlite3" {
		t.Fail()
	}

	if d.Host() != "test.db" {
		t.Fail()
	}

}
