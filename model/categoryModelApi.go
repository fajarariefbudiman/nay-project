package model

import (
	"jar-project/database"
	"log"
	"net/http"

	"github.com/go-playground/validator"
)

func CreateCategoryApi(name, description, slug string) (Response, error) {
	var response Response
	v := validator.New()
	use := Category{
		Name:        name,
		Description: description,
		Slug:        slug,
	}
	err := v.Struct(use)
	if err != nil {
		return response, err
	}
	cond := database.Database()
	sqlstmnt := "INSERT INTO categories (name,description,slug) VALUES (?,?,?)"
	stmt, err := cond.Prepare(sqlstmnt)
	if err != nil {
		return response, err
	}
	result, err := stmt.Exec(name, description, slug)
	if err != nil {
		return response, err
	}
	lastinsert, err := result.LastInsertId()
	if err != nil {
		return response, err
	}
	response.Status = http.StatusCreated
	response.Message = "Success Create"
	response.Data = map[string]int64{
		"lastinsertId": lastinsert,
	}
	return response, nil
}

func GetAllCategoriesApi() (Response, error) {
	var categories []Category
	var res Response
	cond := database.Database()
	sqlstmt := "SELECT name,description,slug,created_at FROM categories"
	rows, err := cond.Query(sqlstmt)
	if err != nil {
		log.Println("No Query Rows")
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var category Category
		err := rows.Scan(&category.Name, &category.Description, &category.Slug, &category.Created_At)
		if err != nil {
			log.Printf("No Rows: %s", err)
		}
		categories = append(categories, category)
	}

	for i := range categories {
		var subCategories []SubCategory
		query := "SELECT * FROM sub_categories WHERE category_id = ?"
		rows, err := cond.Query(query, categories[i].Id)
		if err != nil {
			log.Printf("No rows: %s", err)
		}
		defer rows.Close()
		for rows.Next() {
			var subCategory SubCategory
			err := rows.Scan(&subCategory.Id, &subCategory.Name, &subCategory.Category_Id, &subCategory.Slug)
			if err != nil {
				log.Printf("Error rows Products: %s", err)
			}
			subCategories = append(subCategories, subCategory)
		}

		categories[i].SubCategory = subCategories
		res.Status = http.StatusOK
		res.Message = "Success"
		res.Data = categories
	}
	return res, nil
}

func UpdateCategoryApi(slug, name, description string) (Response, error) {
	var response Response
	cond := database.Database()
	sqlstmt := "UPDATE categories set name = ?, description = ? WHERE slug = ?"
	stmt, err := cond.Prepare(sqlstmt)
	if err != nil {
		return response, err
	}
	result, err := stmt.Exec(name, description, slug)
	if err != nil {
		return response, err
	}
	rowsaffected, err := result.RowsAffected()
	if err != nil {
		return response, err
	}
	response.Status = http.StatusOK
	response.Message = "Success Update"
	response.Data = map[string]int{
		"rowsaffected": int(rowsaffected),
	}
	return response, nil
}

func DeleteCategoryApi(Slug string) (Response, error) {
	var response Response
	cond := database.Database()
	sqlstmnt := "DELETE FROM categories WHERE slug=?"
	stmt, err := cond.Prepare(sqlstmnt)
	if err != nil {
		return response, err
	}

	result, err := stmt.Exec(Slug)
	if err != nil {
		return response, err
	}

	rowsaffected, err := result.RowsAffected()
	if err != nil {
		return response, err
	}

	response.Status = http.StatusNoContent
	response.Message = "Success Delete"
	response.Data = map[string]int64{
		"rowsaffected": rowsaffected,
	}
	return response, err
}
