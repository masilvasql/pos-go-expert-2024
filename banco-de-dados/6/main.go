package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Numero    string
	ProductId int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// category := Category{
	// 	Name: "Eletrônicos",
	// }
	// db.Create(&category)

	// product := Product{
	// 	Name:       "Notebook",
	// 	Price:      9999.00,
	// 	CategoryID: category.ID,
	// }
	// db.Create(&product)

	// db.Create(&SerialNumber{
	// 	Numero:    "123456",
	// 	ProductId: 2,
	// })

	var categories []Category

	err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println("- ", product.Name, " - ", product.SerialNumber.Numero)
		}

	}
}
