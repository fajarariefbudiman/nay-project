package model

import (
	"jar-project/database"
	"net/http"

	"github.com/go-playground/validator"
)

func CreateProductApi(name, description, slug string, price, quantity, sub_category_id int, discount float64) (Response, error) {
	var response Response
	v := validator.New()
	Use := Product{
		Name:            name,
		Description:     description,
		Slug:            slug,
		Price:           price,
		Quantity:        quantity,
		Sub_Category_Id: sub_category_id,
		Discount:        discount,
	}
	err := v.Struct(Use)
	if err != nil {
		return response, err
	}
	cond := database.Database()
	sqlstmnt := "INSERT INTO products (name,description,slug,price,quantity,sub_category_id,discount) VALUES(?,?,?,?,?,?,?)"
	stmt, err := cond.Prepare(sqlstmnt)
	if err != nil {
		return response, err
	}
	result, err := stmt.Exec(name, description, slug, price, quantity, sub_category_id, discount)
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

func GetAllProductsApi() (Response, error) {
	var get Product
	var getAll []Product
	var res Response

	cond := database.Database()
	sqlstmt := "SELECT * FROM products"
	rows, err := cond.Query(sqlstmt)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&get.Id, &get.Name, &get.Slug, &get.Description, &get.Price, &get.Discount, &get.Quantity, &get.Sub_Category_Id, &get.Created_At, &get.Updated_At)
		if err != nil {
			return res, err
		}
		getAll = append(getAll, get)
	}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = getAll
	return res, nil
}

func UpdateProductApi(slug, name, description string, price, quantity, sub_category_id int, discount float64) (Response, error) {
	var response Response
	cond := database.Database()
	sqlstmnt := "UPDATE products set name = ?, description = ?, price = ?, quantity = ?, sub_category_id = ?, discount = ? WHERE slug = ?"
	stmt, err := cond.Prepare(sqlstmnt)
	if err != nil {
		return response, err
	}
	result, err := stmt.Exec(name, description, price, quantity, sub_category_id, discount, slug)
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

func DeleteProductsApi(Slug string) (Response, error) {
	var response Response
	cond := database.Database()
	sqlstmnt := "DELETE FROM products WHERE slug=?"
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
