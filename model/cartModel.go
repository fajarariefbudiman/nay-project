package model

import (
	"database/sql"
	"fmt"
	"jar-project/database"
)

type Cart struct {
	Id               int64
	User_Id          int
	Created_At       string
	Updated_At       string
	Base_Total_Price float64
	Grand_Total      float64
	Quantity         int //cart_items
	Product_Id       int
	Products         []Product
}

func CreateCart(userId, product_id, quantity int) error {
	cond := database.Database()
	sqlstmnt := "INSERT INTO carts (user_id, product_id, quantity,base_total_price) VALUES(?, ?, ?)"
	stmt, err := cond.Prepare(sqlstmnt)
	if err != nil {
		fmt.Printf("Error Prepare %v", sqlstmnt)
		return err
	}
	result, err := stmt.Exec(userId, product_id, quantity)
	if err != nil {
		fmt.Printf("Error %v", result)
		return err
	}
	return nil
}

func GetAllCart(user_id int) ([]Cart, error) {
	var getAll []Cart
	cond := database.Database()
	sqlstmt := "SELECT user_id, quantity, product_id FROM carts WHERE user_id = ?"
	rows, err := cond.Query(sqlstmt, user_id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var get Cart
		err := rows.Scan(&get.User_Id, &get.Quantity, &get.Product_Id)
		if err != nil {
			panic(err)
		}
		product, err := getProduct(get.Product_Id)
		if err != nil {
			panic(err)
		}
		get.Products = append(get.Products, product)
		getAll = append(getAll, get)
	}
	return getAll, nil
}

func GetCart(user_id int) (Cart, error) {
	var get Cart
	cond := database.Database()
	sqlstmt := "SELECT user_id, quantity, product_id FROM carts WHERE user_id = ?"
	rows, err := cond.Query(sqlstmt, user_id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&get.User_Id, &get.Quantity, &get.Product_Id)
		if err == sql.ErrNoRows {
			panic(err)
		}
	}
	return get, nil
}
func getProduct(product_id int) (Product, error) {
	var product Product
	cond := database.Database()
	query := "SELECT * FROM products WHERE id = ?"
	row := cond.QueryRow(query, product_id)
	err := row.Scan(&product.Id, &product.Name, &product.Description, &product.Slug, &product.Price, &product.Quantity,
		&product.Sub_Category_Id, &product.Discount, &product.Created_At, &product.Updated_At, &product.Body)
	if err != nil {
		return product, err
	}
	return product, nil
}
