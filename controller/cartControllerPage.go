package controller

import (
	"fmt"
	"html/template"
	"jar-project/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func Cart(c echo.Context) error {
	renderer := &TemplateRenderer{
		Template: template.Must(template.ParseGlob("./template/*.html")),
	}
	product, err := model.GetAllProducts()
	if err != nil {
		return err
	}
	discountProduct, err := model.GetProductDiscount()
	if err != nil {
		return err
	}
	sess, _ := session.Get("session", c)
	isAuthenticated := true
	if userID, ok := sess.Values["user_id"].(int); ok {
		user, err := model.GetUserByID(userID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		cart_items, err := model.GetAllCartItems(userID)
		if err != nil {
			return c.Redirect(http.StatusSeeOther, "/p/:slug")
		}
		cart_item, err := model.GetCartItems(userID)
		if err != nil {
			return c.Redirect(http.StatusSeeOther, "/p/:slug")
		}
		return renderer.Render(c.Response().Writer, "cart.html", map[string]interface{}{
			"products":          product,
			"discountedProduct": discountProduct,
			"Auth":              isAuthenticated,
			"user":              user,
			"cart_items":        cart_items,
			"cart_item":         cart_item,
		}, c)
	}
	return c.Redirect(http.StatusSeeOther, "/login")
}

func AddCart(c echo.Context) error {
	productId := c.FormValue("product_id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		fmt.Println("ProductId")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": err,
		})
	}
	quantity := c.FormValue("quantity")
	qty, err := strconv.Atoi(quantity)
	if err != nil {
		fmt.Println("Quantity")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": err,
		})
	}
	price := c.FormValue("price")
	strPrice, err := strconv.Atoi(price)
	if err != nil {
		fmt.Println("UserId")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": err,
		})
	}
	userid := c.FormValue("user_id")
	userId, err := strconv.Atoi(userid)
	if err != nil {
		fmt.Println("UserId")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": err,
		})
	}
	err = model.CreateCartItems(userId, id, qty, strPrice)
	if err != nil {
		return err
	}
	return c.Redirect(http.StatusSeeOther, "/cart")
}
