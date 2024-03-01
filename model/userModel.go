package model

import (
	"database/sql"
	"fmt"
	"jar-project/database"

	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id         int    `json:"id"`
	Firstname  string `validate:"required" json:"firstname"`
	Lastname   string `validate:"required" json:"lastname"`
	Email      string `validate:"required|email" json:"email"`
	Password   string `validate:"required" json:"password"`
	Created_At string
	Updated_At string
}

type Login struct {
	Email    string
	Password string
}

func CreateUser(firstname, lastname, email, password string) error {
	v := validator.New()
	Use := User{
		Firstname: firstname,
		Lastname:  lastname,
		Email:     email,
		Password:  password,
	}
	err := v.Struct(Use)
	if err != nil {
		fmt.Println("Validation Error")
		return err
	}
	cond := database.Database()
	sqlstmnt := "INSERT INTO users (firstname,lastname,email,password) VALUES(?,?,?,?)"
	stmt, err := cond.Prepare(sqlstmnt)
	if err != nil {
		fmt.Printf("Error Prepare %v", sqlstmnt)
		return err
	}
	_, err = stmt.Exec(firstname, lastname, email, password)
	if err != nil {
		fmt.Println("Exec Error")
		return err
	}

	return nil
}

func AuthenticateUser(email, password string) (bool, error) {
	v := validator.New()
	Use := Login{
		Email:    email,
		Password: password,
	}
	err := v.Struct(Use)
	if err != nil {
		fmt.Println("Validation Error")
		return false, err
	}
	var user User
	con := database.Database()
	sqlstmt := "SELECT * FROM users WHERE email = ?"
	err = con.QueryRow(sqlstmt, email).Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Email, &user.Password, &user.Created_At, &user.Updated_At)
	if err != nil {
		fmt.Printf("Email: %v", email)
		fmt.Println("No Row Selected")
		return false, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Println("Compare Password Error")
		return false, err
	}

	return true, nil
}

func GetUserByEmail(email string) (User, error) {
	var get User
	cond := database.Database()
	sqlstmt := "SELECT id,firstname,lastname,email,password,created_at,updated_at FROM users WHERE email=?"
	rows, err := cond.Query(sqlstmt, email)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&get.Id, &get.Firstname, &get.Lastname, &get.Email, &get.Password, &get.Created_At, &get.Updated_At)
		if err == sql.ErrNoRows {
			panic(err)
		}
	}
	return get, nil
}

func GetUserByID(userID int) (User, error) {
	var get User
	cond := database.Database()
	sqlstmt := "SELECT id,firstname,lastname,email,password,created_at,updated_at FROM users WHERE id=?"
	rows, err := cond.Query(sqlstmt, userID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&get.Id, &get.Firstname, &get.Lastname, &get.Email, &get.Password, &get.Created_At, &get.Updated_At)
		if err == sql.ErrNoRows {
			panic(err)
		}
	}
	return get, nil
}

func UpdatePassword(email, newPassword string) error {
	db := database.Database()
	query := "UPDATE users SET password = ? WHERE email = ?"
	_, err := db.Exec(query, newPassword, email)
	if err != nil {
		return err
	}
	return nil
}
