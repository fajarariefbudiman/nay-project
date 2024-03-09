package model

import (
	"fmt"
	"jar-project/database"
)

type Cart_Item struct {
	Id         int
	User_Id    int
	Product_Id int
	Quantity   int //product
	Base_Price float64
	Created_At string
	Updated_At string
	Products   []Product
}

func CreateCartItems(userId, product_id, quantity, price int) error {
	cond := database.Database()
	sqlstmnt := "INSERT INTO cart_items (user_id, product_id, quantity, base_price) VALUES(?, ?, ?, ?)"
	stmt, err := cond.Prepare(sqlstmnt)
	if err != nil {
		fmt.Printf("Error Prepare %v", sqlstmnt)
		return err
	}
	basePrice := price * quantity
	result, err := stmt.Exec(userId, product_id, quantity, basePrice)
	if err != nil {
		fmt.Printf("Error %v", result)
		return err
	}
	return nil
}

func GetAllCartItems(user_id int) ([]Cart_Item, error) {
	var getAll []Cart_Item
	cond := database.Database()
	sqlstmt := "SELECT user_id, quantity, product_id, base_price FROM cart_items WHERE user_id = ?"
	rows, err := cond.Query(sqlstmt, user_id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var get Cart_Item
		err := rows.Scan(&get.User_Id, &get.Quantity, &get.Product_Id, &get.Base_Price)
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

func GetCartItems(user_id int) (Cart_Item, error) {
	var cartItem Cart_Item
	var totalQuantity int
	var totalBasePrice float64
	cond := database.Database()
	sqlstmt := "SELECT quantity, base_price FROM cart_items WHERE user_id = ?"
	rows, err := cond.Query(sqlstmt, user_id)
	if err != nil {
		return cartItem, err
	}
	defer rows.Close()

	for rows.Next() {
		var quantity int
		var basePrice float64
		err = rows.Scan(&quantity, &basePrice)
		if err != nil {
			return cartItem, err
		}
		totalQuantity += quantity
		totalBasePrice += basePrice
	}
	if totalQuantity == 0 {
		return cartItem, err
	}
	cartItem.User_Id = user_id
	cartItem.Quantity = totalQuantity
	cartItem.Base_Price = totalBasePrice

	return cartItem, nil
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
