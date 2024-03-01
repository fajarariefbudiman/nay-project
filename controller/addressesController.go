package controller

import (
	"fmt"
	"jar-project/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddAddresses(c echo.Context) error {
	user_id := c.Param("id")
	strid, err := strconv.Atoi(user_id)
	if err != nil {
		return err
	}
	firstName := c.FormValue("firstName")
	lastName := c.FormValue("lastName")
	phone := c.FormValue("phone")
	address1 := c.FormValue("address1")
	province := c.FormValue("province")
	city := c.FormValue("city")
	subdistrict := c.FormValue("subdistrict")
	postCode := c.FormValue("postCode")
	err = model.AddAddressesModel(strid, address1, subdistrict, city, province, postCode, phone, firstName, lastName)
	if err != nil {
		return err
	}
	return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/dashboard/%d", strid))
}
