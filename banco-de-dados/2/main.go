package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})
	// create(db)
	// createBatch(db)
	// findById(1, db)
	// findFirst(db)
	// findByName("Notebook", db)
	// findAll(db)

	// var product Product
	// db.First(&product)
	// product.Name = "Big Table"
	// // update(&product, db)
	// deleteByProduct(&product, db)
	// delete(4, db)

}

func create(db *gorm.DB) {
	db.Create(&Product{
		Name:  "Table",
		Price: 120.00,
	})
}

func createBatch(db *gorm.DB) {
	products := []*Product{
		{Name: "Notebook", Price: 1000.00},
		{Name: "Mouse", Price: 100.00},
		{Name: "Keyboard", Price: 399.00},
	}

	db.Create(&products)
}

func findById(id int, db *gorm.DB) {
	var p Product

	db.First(&p, id)

	fmt.Printf("O produto encontrado foi %v e seu valor é de %.2f ", p.Name, p.Price)
}

func findFirst(db *gorm.DB) {
	var p Product

	db.First(&p)

	fmt.Printf("O produto encontrado foi %v e seu valor é de %.2f ", p.Name, p.Price)
}

func findByName(name string, db *gorm.DB) {
	var p Product

	db.Where("name = ?", name).First(&p)

	fmt.Printf("O produto encontrado foi %v e seu valor é de %.2f ", p.Name, p.Price)
}

func findAll(db *gorm.DB) {
	var products []Product

	db.Find(&products)

	for _, product := range products {
		fmt.Printf("O produto encontrado foi %v e seu valor é de %.2f\n", product.Name, product.Price)
	}

}

func update(product *Product, db *gorm.DB) {
	db.Save(&product)
	findById(product.ID, db)
}

func delete(id int, db *gorm.DB) {
	var p Product
	db.Where("id = ?", id).Delete(&p)
}

func deleteByProduct(p *Product, db *gorm.DB) {
	db.Delete(&p)
}
