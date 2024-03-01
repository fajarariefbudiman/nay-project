package model

import (
	"fmt"
	"jar-project/database"
)

type Addresses struct {
	Id          int
	User_Id     int
	Address1    string
	Subdistrict string
	City        string
	Province    string
	Postcode    string
	Phone       string
	Created_At  string
	Updated_At  string
	Firstname   string
	Lastname    string
}

func AddAddressesModel(user_id int, address1, subdistrict, city, province, postcode, phone, firstname, lastname string) error {
	cond := database.Database()
	sqlstmnt := "INSERT INTO addresses (user_id,address1,subdistrict,city,province,postcode,phone,firstname,lastname) VALUES (?,?,?,?,?,?,?,?,?)"
	stmt, err := cond.Prepare(sqlstmnt)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(user_id, address1, subdistrict, city, province, postcode, phone, firstname, lastname)
	if err != nil {
		fmt.Printf("Error %v", result)
		return err
	}
	return nil
}

func GetAddress(userId int) (Addresses, error) {
	var get Addresses
	cond := database.Database()
	sqlstmnt := "SELECT user_id,address1,subdistrict,city,province,postcode,phone,firstname,lastname FROM addresses WHERE user_id = ?"
	rows, err := cond.Query(sqlstmnt, userId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&get.User_Id, &get.Address1, &get.Subdistrict, &get.City, &get.Province, &get.Postcode, &get.Phone, &get.Firstname, &get.Lastname)
		if err != nil {
			panic(err)
		}
	}
	return get, nil
}
