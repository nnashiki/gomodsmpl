package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type User struct {
	gorm.Model
	Name      string
	CompanyID int
	Company   Company
}

type Company struct {
	ID   int
	Name string
}

func main() {
	// 詳細は https://github.com/go-sql-driver/mysql#dsn-data-source-name を参照
	dsn := "baseball_user:baseball_pass@tcp(127.0.0.1:3306)/baseball_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Migrate the schema
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Company{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})
	db.Create(&User{Name: "nashiki", Company: Company{Name: "サイバーダイン"}})

	// Read
	var product Product
	db.First(&product, "code = ?", "D42")
	print(product.Price)

	/*
		// Update - update product's price to 200
		db.Model(&product).Update("Price", 200)
		// Update - update multiple fields
		db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
		db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

		// Delete - delete product
		db.Delete(&product, 1)

	*/
}
