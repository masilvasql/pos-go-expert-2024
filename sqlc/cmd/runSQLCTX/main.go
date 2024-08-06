package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/masilvasql/sqlc/internal/db"
)

type CourseDB struct {
	dbConn *sql.DB
	*db.Queries
}

type CourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
	Price       float64
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
	CategoryID  string
}

func NewCourseDB(dbConn *sql.DB) *CourseDB {
	return &CourseDB{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

func (c *CourseDB) callTX(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := db.New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb error: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

func (c *CourseDB) CreateCourseAndCategory(ctx context.Context, argsCategory CategoryParams, argsCourse CourseParams) error {
	err := c.callTX(ctx, func(q *db.Queries) error {

		err := q.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          argsCategory.ID,
			Name:        argsCategory.Name,
			Description: argsCategory.Description,
		})

		if err != nil {
			return err
		}

		err = q.CreateCourse(ctx, db.CreateCourseParams{
			ID:          argsCourse.ID,
			Name:        argsCourse.Name,
			Description: argsCourse.Description,
			CategoryID:  argsCategory.ID,
			Price:       argsCourse.Price,
		})

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func main() {
	ctx := context.Background()

	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}

	defer dbConn.Close()

	queries := db.New(dbConn)

	//courseDB := NewCourseDB(dbConn)
	//
	//idCategory := uuid.New().String()
	//
	//category := CategoryParams{
	//	ID:          idCategory,
	//	Name:        "Category 5",
	//	Description: sql.NullString{String: "Description", Valid: true},
	//}
	//
	//coruse := CourseParams{
	//	ID:          uuid.New().String(),
	//	Name:        "Course 6",
	//	Description: sql.NullString{String: "Description", Valid: true},
	//	Price:       9.99,
	//}
	//
	//err = courseDB.CreateCourseAndCategory(ctx, category, coruse)
	//
	//if err != nil {
	//	panic(err)
	//}

	courses, err := queries.ListCourses(ctx)
	if err != nil {
		panic(err)
	}

	for _, course := range courses {
		fmt.Printf("ID: %s, Name: %s, Description: %s, Price: %f, CategoryID: %s, CategoryName %s\n", course.ID, course.Name, course.Description.String, course.Price, course.CategoryID, course.CategoryName)
	}

}
