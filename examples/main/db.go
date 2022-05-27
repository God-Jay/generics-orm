package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDb() *gorm.DB {
	db, _ := gorm.Open(mysql.Open(
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True", "homestead", "secret", "localhost", 33060, "gorm")),
	)
	return db
}
