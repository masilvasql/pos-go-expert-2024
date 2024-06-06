package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{})

	// category := Category{
	// 	Name: "Eletrônicos",
	// }
	// db.Create(&category)

	// product := Product{
	// 	Name:       "Notebook",
	// 	Price:      1000.00,
	// 	CategoryID: category.ID,
	// }
	// db.Create(&product)

	var products []Product

	db.Preload("Category").Find(&products)

	for _, product := range products {
		fmt.Println(product.Category)
	}
}
