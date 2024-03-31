package controller

import (
	"html/template"
	"jar-project/model"
	"net/http"
	"net/url"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func GetProductBySubCategory(c echo.Context) error {
	category := c.Param("category")
	subCategory := c.Param("subCategory")
	renderer := &TemplateRenderer{
		Template: template.Must(template.ParseGlob("./template/*.html")),
	}
	search := c.QueryParam("search")
	if search != "" {
		u, _ := url.Parse("/p")
		q := u.Query()
		q.Set("search", search)
		u.RawQuery = q.Encode()
		return c.Redirect(http.StatusSeeOther, u.String())
	}
	products,err := model.SubCategoryProduct(category,subCategory)
	if err != nil {
		return err
	}
	categories, err := model.GetAllCategories()
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
		return renderer.Render(c.Response().Writer, "products.html", map[string]interface{}{
			"user":              user,
			"categories":        categories,
			"Auth":              isAuthenticated,
			"products":          products,
			"discountedProduct": discountProduct,
		}, c)
	}
	return renderer.Render(c.Response().Writer,"products.html",map[string]interface{}{
		"products":products,
		"discountedProduct": discountProduct,
		"categories":        categories,
	},c)
}