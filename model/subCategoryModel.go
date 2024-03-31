package model

import (
	"fmt"
	"jar-project/database"
)

type SubCategory struct {
	Id          int
	Name        string
	Category_Id int
	Slug        string
	Category    Category
	Product 	[]Product
}

func SubCategoryProduct(slugCategory, slugSubCategory string) ([]Product, error) {
	var products []Product
	cond := database.Database()
	var categoryId int
	query := "SELECT id FROM categories WHERE slug = ?"
	err := cond.QueryRow(query, slugCategory).Scan(&categoryId)
	if err != nil {
		return nil, err
	}
	var subCategoryId int
	query = "SELECT id FROM sub_categories WHERE category_id = ? AND slug = ?"
	err = cond.QueryRow(query, categoryId, slugSubCategory).Scan(&subCategoryId)
	if err != nil {
		return nil, err
	}
	query = "SELECT * FROM products WHERE sub_category_id = ?"
	rows, err := cond.Query(query, subCategoryId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Slug, &product.Price, &product.Quantity,
			&product.Discount, &product.Sub_Category_Id, &product.Created_At, &product.Updated_At, &product.Body)
		if err != nil {
			fmt.Println("Error scanning product:", err)
			return nil, err
		}
		product.DiscountedPrice = (float64(product.Price) / 100) * product.Discount
		products = append(products, product)
	}
	return products, nil
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



