package controller

import (
	"html/template"
	"jar-project/model"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func Product(c echo.Context) error {
	slug := c.Param("slug")
	renderer := &TemplateRenderer{
		Template: template.Must(template.ParseGlob("./template/*html")),
	}
	product, err := model.GetProductBySlug(slug)
	if err != nil {
		return err
	}
	categories, err := model.GetAllCategories()
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
		return renderer.Render(c.Response().Writer, "products.html", map[string]interface{}{
			"user":       user,
			"categories": categories,
			"Auth":       isAuthenticated,
			"product":    product,
		}, c)
	}
	return renderer.Render(c.Response().Writer, "products.html", map[string]interface{}{
		"categories": categories,
		"product":    product,
	}, c)
}
