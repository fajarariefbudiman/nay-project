package controller

import (
	"html/template"
	"jar-project/model"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func Checkout(c echo.Context) error {
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
		address, err := model.GetAddress(userID)
		if err != nil {
			return err
		}
		cart_items, err := model.GetAllCartItems(userID)
		if err != nil {
			return c.Redirect(http.StatusSeeOther, "/p/:slug")
		}
		cart_item, err := model.GetCartItems(userID)
		if err != nil {
			return c.Redirect(http.StatusSeeOther, "/p/:slug")
		}
		return renderer.Render(c.Response().Writer, "checkout.html", map[string]interface{}{
			"products":          product,
			"discountedProduct": discountProduct,
			"Auth":              isAuthenticated,
			"addresses":         address,
			"user":              user,
			"cart_items":        cart_items,
			"cart_item":         cart_item,
		}, c)
	}
	return c.Redirect(http.StatusSeeOther, "/login")
}
