package model

import (
	"database/sql"
	"jar-project/database"
	"log"
)

type Category struct {
	Id          int    `json:"id"`
	Name        string `validate:"required" json:"name"`
	Description string `validate:"required" json:"description"`
	Slug        string `validate:"required" json:"slug"`
	Created_At  string `json:"created_at"`
	Updated_At  string `json:"updated_at"`
	SubCategory []SubCategory
}

func GetAllCategories() ([]Category, error) {
	var categories []Category
	cond := database.Database()
	sqlstmt := "SELECT id,name,description,slug,created_at FROM categories"
	rows, err := cond.Query(sqlstmt)
	if err != nil {
		log.Println("No Query Rows")
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var get Category
		err := rows.Scan(&get.Id, &get.Name, &get.Description, &get.Slug, &get.Created_At)
		if err != nil {
			panic(err)
		}
		categories = append(categories, get)
	}
	return categories, nil
}

func GetCategoryBySlug(slug string) (Category, error) {
	var get Category
	cond := database.Database()
	sqlstmt := "SELECT name,description,slug, created_at FROM categories WHERE slug = ?"
	rows, err := cond.Query(sqlstmt, slug)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&get.Name, &get.Description, &get.Slug, &get.Created_At)
		if err == sql.ErrNoRows {
			panic(err)
		}
	}
	return get, nil
}
