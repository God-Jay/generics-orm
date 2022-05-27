package main

import (
	"context"
	"github.com/god-jay/generics-orm/gorm"
	"log"
)

func main() {
	ctx := context.Background()

	db := NewDb()

	user, err := gorm.NewModel(ctx, db, User{}).Create()
	if err != nil {
		return
	}
	log.Println(user)

}
