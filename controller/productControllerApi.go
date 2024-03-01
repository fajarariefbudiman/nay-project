package controller

import (
	"jar-project/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateProductApi(c echo.Context) error {
	name := c.FormValue("name")
	description := c.FormValue("description")
	slug := c.FormValue("slug")
	price := c.FormValue("price")
	strprice, err := strconv.Atoi(price)
	if err != nil {
		return err
	}
	quantity := c.FormValue("quantity")
	strquantity, err := strconv.Atoi(quantity)
	if err != nil {
		return err
	}
	category := c.FormValue("sub_category_id")
	sub_category_id, err := strconv.Atoi(category)
	discount := c.FormValue("discount")
	strdiscount, err := strconv.ParseFloat(discount, 64)
	if err != nil {
		return err
	}
	result, err := model.CreateProductApi(name, description, slug, strprice, strquantity, sub_category_id, strdiscount)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, result)
}

func GetProductApi(c echo.Context) error {
	result, err := model.GetAllProductsApi()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateProductApi(c echo.Context) error {
	slug := c.Param("slug")
	name := c.FormValue("name")
	description := c.FormValue("description")
	price := c.FormValue("price")
	strprice, err := strconv.Atoi(price)
	if err != nil {
		return err
	}
	quantity := c.FormValue("quantity")
	strquantity, err := strconv.Atoi(quantity)
	category := c.FormValue("sub_category_id")
	sub_category_id, err := strconv.Atoi(category)
	discount := c.FormValue("discount")
	strdiscount, err := strconv.ParseFloat(discount, 64)
	result, err := model.UpdateProductApi(slug, name, description, strprice, strquantity, sub_category_id, strdiscount)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, result)
}

func DeleteProductApi(c echo.Context) error {
	slug := c.Param("slug")
	result, err := model.DeleteProductsApi(slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}
	return c.JSON(http.StatusNoContent, result)
}
