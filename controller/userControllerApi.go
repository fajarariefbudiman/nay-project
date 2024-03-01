package controller

import (
	"io/ioutil"
	"jar-project/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

func CreateUserApi(c echo.Context) error {
	firstname := c.FormValue("firstname")
	lastname := c.FormValue("lastname")
	email := c.FormValue("email")
	password := c.FormValue("password")
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hashed := string(hashedPassword)
	result, err := model.CreateUserApi(firstname, lastname, email, hashed)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	htmlcontent, _ := ioutil.ReadFile("../template/email.html")
	m := gomail.NewMessage()
	m.SetHeader("From", "admin123@example.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "From ")
	m.SetBody("text/html", string(htmlcontent))
	d := gomail.NewDialer("smtp.gmail.com", 587, "budimanfajar660@gmail.com", "bwtoyhwtdjwugbwq")

	if err := d.DialAndSend(m); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}

func GetUserApi(c echo.Context) error {
	result, err := model.GetAllUsersApi()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, result)
}

func UpdateUserApi(c echo.Context) error {
	firstname := c.FormValue("firstname")
	lastname := c.FormValue("lastname")
	email := c.FormValue("email")
	password := c.FormValue("password")
	id := c.Param("id")
	strid, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	result, err := model.UpdateUserApi(strid, firstname, lastname, email, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteUserApi(c echo.Context) error {
	id := c.Param("id")
	strid, err := strconv.Atoi(id)
	result, err := model.DeleteUserApi(strid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}
	return c.JSON(http.StatusNoContent, result)
}
