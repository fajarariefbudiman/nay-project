package model

import (
	"jar-project/database"
)

type SubCategory struct {
	Id          int
	Name        string
	Category_Id int
	Slug        string
	Category    Category
}

func GetSubCategoryByCategorySlug(categorySlug string) ([]SubCategory, error) {
	var subCategories []SubCategory
	cond := database.Database()
	var categoryId int
	query := "SELECT id,slug FROM categories WHERE slug = ?"
	err := cond.QueryRow(query, categorySlug).Scan(&categoryId, &categorySlug)
	if err != nil {
		return nil, err
	}
	query = "SELECT id, name, category_id, slug FROM sub_categories WHERE category_id = ?"
	rows, err := cond.Query(query, categoryId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var subCategory SubCategory
		err := rows.Scan(&subCategory.Id, &subCategory.Name, &subCategory.Category_Id, &subCategory.Slug)
		if err != nil {
			return nil, err
		}
		category, err := GetCategoryBySlug(categorySlug)
		if err != nil {
			return nil, err
		}
		subCategory.Category.Slug = category.Slug
		subCategories = append(subCategories, subCategory)
	}
	return subCategories, nil
}
