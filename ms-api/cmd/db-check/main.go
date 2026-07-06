package main

import (
	"fmt"
	"log"

	"ms-api/app/models"
)


func main() {
	db := models.DbConnection
	if db == nil {
		log.Fatalln("DbConnection is nil")
	}

	var n int
	if err := db.Raw("SELECT 1").Scan(&n).Error; err != nil {
		log.Fatalln("SELECT 1 failed: ", err)
	}
	fmt.Println("OK: SELECT 1 =>",n)

	var count int64

	if err := db.Raw("SELECT COUNT(*) FROM videos").Scan(&count).Error; err != nil {
		fmt.Println("Failed to count videos",n)
		return
	}
	fmt.Println("OK: videos table count =>",count)
}
