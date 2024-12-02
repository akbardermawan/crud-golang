package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB adalah variabel global untuk koneksi database
var DB *gorm.DB

func DatabaseInit() {

	var err error

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	const MYSQl = "root:@tcp(127.0.0.1:3306)/gofiber?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := MYSQl

	// Membuka koneksi ke database
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	
	if err  != nil {
		log.Panic("Koneksi database gagal: %v\n", err)
	}

	fmt.Println("Berhasil terhubung ke database")
	
}