package model

import (
	"jar-project/database"
	"net/http"

	"github.com/go-playground/validator"
)

func CreateUserApi(firstname, lastname, email, password string) (Response, error) {
	var response Response
	v := validator.New()
	Use := User{
		Firstname: firstname,
		Lastname:  lastname,
		Email:     email,
		Password:  password,
	}
	err := v.Struct(Use)
	if err != nil {
		return response, err
	}
	cond := database.Database()
	sqlstmnt := "INSERT INTO users (firstname,lastname,email,password) VALUES(?,?,?,?)"
	stmt, err := cond.Prepare(sqlstmnt)
	if err != nil {
		return response, err
	}
	result, err := stmt.Exec(firstname, lastname, email, password)
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

func GetAllUsersApi() (Response, error) {
	var get User
	var getAll []User
	var res Response
	cond := database.Database()
	sqlstmt := "SELECT * FROM users"
	rows, err := cond.Query(sqlstmt)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&get.Id, &get.Firstname, &get.Lastname, &get.Email, &get.Password, &get.Created_At, &get.Updated_At)
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

func UpdateUserApi(id int, firstname, lastname, email, password string) (Response, error) {
	var response Response
	cond := database.Database()
	sqlstmnt := "UPDATE users set firstname=?, lastname=?, email=?, password=? WHERE id=?"
	stmt, err := cond.Prepare(sqlstmnt)
	if err != nil {
		return response, err
	}
	result, err := stmt.Exec(firstname, lastname, email, password, id)
	if err != nil {
		return response, err
	}
	rowsaffected, err := result.RowsAffected()
	if err != nil {
		return response, err
	}
	response.Status = http.StatusOK
	response.Message = "Succes Update"
	response.Data = map[string]int64{
		"rowsaffected": rowsaffected,
	}
	return response, err
}

func DeleteUserApi(id int) (Response, error) {
	var response Response
	cond := database.Database()
	sqlstmnt := "DELETE FROM users WHERE id=?"
	stmt, err := cond.Prepare(sqlstmnt)
	if err != nil {
		return response, err
	}

	result, err := stmt.Exec(id)
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
