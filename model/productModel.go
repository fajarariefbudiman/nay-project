package model

import (
	"database/sql"
	"jar-project/database"
)

type Product struct {
	Id              int     `json:"id"`
	Name            string  `validate:"required" json:"name"`
	Description     string  `validate:"required" json:"description"`
	Slug            string  `validate:"required" json:"slug"`
	Price           int     `validate:"required" json:"price"`
	Quantity        int     `validate:"required" json:"quantity"`
	Sub_Category_Id int     `validate:"required" json:"sub_category_id"`
	Discount        float64 `validate:"required" json:"discount"`
	DiscountedPrice float64
	Created_At      string `json:"created_at"`
	Updated_At      string `json:"updated_at"`
	Body            string
}

func GetProductDiscount() ([]Product, error) {
	cond := database.Database()
	sqlstmt := "SELECT name,description,slug,price,sub_category_id,body FROM products"
	rows, err := cond.Query(sqlstmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.Name, &p.Description, &p.Slug, &p.Price, &p.Sub_Category_Id, &p.Body); err != nil {
			return nil, err
		}
		p.DiscountedPrice = (float64(p.Price) / 100) * p.Discount
		products = append(products, p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func GetAllProducts() ([]Product, error) {
	var get Product
	var getAll []Product
	cond := database.Database()
	sqlstmt := "SELECT * FROM products"
	rows, err := cond.Query(sqlstmt)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&get.Id, &get.Name, &get.Description, &get.Slug, &get.Price, &get.Discount, &get.Quantity, &get.Sub_Category_Id, &get.Created_At, &get.Updated_At, &get.Body)
		if err != nil {
			return nil, err
		}
		getAll = append(getAll, get)
	}
	return getAll, nil
}

func GetProductBySlug(slug string) (Product, error) {
	var get Product
	cond := database.Database()
	sqlstmt := "SELECT * FROM products WHERE slug=?"
	rows, err := cond.Query(sqlstmt, slug)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&get.Id, &get.Name, &get.Description, &get.Slug, &get.Price, &get.Quantity, &get.Sub_Category_Id,
			&get.Discount, &get.Created_At, &get.Updated_At, &get.Body)
		if err == sql.ErrNoRows {
			panic(err)
		}
	}
	return get, nil
}

func GetProductById(id int) (Product, error) {
	var get Product
	cond := database.Database()
	sqlstmt := "SELECT * FROM products WHERE id=?"
	rows, err := cond.Query(sqlstmt, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&get.Id, &get.Name, &get.Description, &get.Slug, &get.Price, &get.Quantity, &get.Sub_Category_Id,
			&get.Discount, &get.Created_At, &get.Updated_At, &get.Body)
		if err == sql.ErrNoRows {
			panic(err)
		}
	}
	return get, nil
}
