package main

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/masilvasql/sqlc/internal/db"
)

func main() {
	ctx := context.Background()

	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}

	defer dbConn.Close()

	queries := db.New(dbConn)
	//
	//idCategory := uuid.New().String()
	//err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	//	ID:          idCategory,
	//	Name:        "Category 3",
	//	Description: sql.NullString{String: "Description", Valid: true},
	//})
	//
	//if err != nil {
	//	panic(err)
	//}
	//categories, err := queries.ListCategories(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//
	//for _, category := range categories {
	//	c := fmt.Sprintf("Category: %v\n", category)
	//	fmt.Println(c)
	//}
	//
	//err = queries.CreateCourse(ctx, db.CreateCourseParams{
	//	ID:          uuid.New().String(),
	//	Name:        "Course 2",
	//	Description: sql.NullString{String: "Description", Valid: true},
	//	Price:       "9.99",
	//	CategoryID:  idCategory,
	//})
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//courses, err := queries.ListCourses(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//
	//for _, course := range courses {
	//	c := fmt.Sprintf("Course: %v\n", course)
	//	fmt.Println(c)
	//}

	//err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
	//	ID:          "1f81f7e1-9663-43f5-bbe1-b6ad52012faa",
	//	Description: sql.NullString{String: "Description Updated", Valid: true},
	//	Name:        "Category Updated",
	//})
	//
	//if err != nil {
	//	panic(err)
	//}

	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(ctg)
	//ctg, err := queries.GetCategory(ctx, "1f81f7e1-9663-43f5-bbe1-b6ad52012faa")

	err = queries.DeleteCategory(ctx, "3594aa77-670b-4311-91e1-d3f85718fda2")
	if err != nil {
		panic(err)
	}

}
