package main

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"testing"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func TestSqlite(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1) // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	db.Delete(&product, 1)
}
func TestLen(t *testing.T) {
	m := make([]ChatMessage,10)
	fmt.Println("m len ", len(m))
	m = append(m,ChatMessage{})
	m = append(m,ChatMessage{})
	fmt.Println("m len ", len(m))
	m = m[:0]
	fmt.Println("m len ", len(m))

}