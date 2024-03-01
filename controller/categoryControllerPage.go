package controller

import (
	"fmt"
	"html/template"
	"jar-project/model"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func Category(c echo.Context) error {
	slug := c.Param("slug")
	renderer := &TemplateRenderer{
		Template: template.Must(template.ParseGlob("./template/*.html")),
	}
	category, err := model.GetCategoryBySlug(slug)
	if err != nil {
		return err
	}
	product, err := model.GetAllProducts()
	if err != nil {
		return err
	}
	discountProduct, err := model.GetProductDiscount()
	if err != nil {
		return err
	}
	categories, err := model.GetAllCategories()
	if err != nil {
		return err
	}
	subCategory, err := model.GetSubCategoryByCategorySlug(slug)
	if err != nil {
		fmt.Println("Error Get Sub Category")
		return err
	}
	sess, _ := session.Get("session", c)
	isAuthenticated := true
	if userID, ok := sess.Values["user_id"].(int); ok {
		user, err := model.GetUserByID(userID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return renderer.Render(c.Response().Writer, "categories.html", map[string]interface{}{
			"user":              user,
			"subCategory":       subCategory,
			"categories":        categories,
			"Auth":              isAuthenticated,
			"products":          product,
			"category":          category,
			"discountedProduct": discountProduct,
		}, c)
	}
	return renderer.Render(c.Response().Writer, "categories.html", map[string]interface{}{
		"products":          product,
		"categories":        categories,
		"discountedProduct": discountProduct,
	}, c)
}
