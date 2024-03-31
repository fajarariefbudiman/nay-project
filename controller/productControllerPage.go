package controller

import (
	"html/template"
	"jar-project/model"
	"net/http"
	"net/url"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func Product(c echo.Context) error {
	slug := c.Param("slug")
	renderer := &TemplateRenderer{
		Template: template.Must(template.ParseGlob("./template/*html")),
	}
	search := c.QueryParam("search")
	if search != "" {
		u, _ := url.Parse("/p")
		q := u.Query()
		q.Set("search", search)
		u.RawQuery = q.Encode()
		return c.Redirect(http.StatusSeeOther, u.String())
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
		return renderer.Render(c.Response().Writer, "product.html", map[string]interface{}{
			"user":       user,
			"categories": categories,
			"Auth":       isAuthenticated,
			"product":    product,
		}, c)
	}
	return renderer.Render(c.Response().Writer, "product.html", map[string]interface{}{
		"categories": categories,
		"product":    product,
	}, c)
}


func ListProduct(c echo.Context) error {
    renderer := &TemplateRenderer{
        Template: template.Must(template.ParseGlob("./template/*html")),
    }
    search := c.QueryParam("search")
    categories, err := model.GetAllCategories()
    if err != nil {
        return err
    }
	discountProduct, err := model.GetProductDiscount()
	if err != nil {
		return err
	}
    searchQuery, err := model.SearchProduct(search)
    if err != nil {
        return renderer.Render(c.Response().Writer, "list.html", map[string]interface{}{
            "product":    searchQuery,
            "categories": categories,
			"discountedProduct": discountProduct,
            "errorText":  err.Error(),
        }, c)
    }
    if len(searchQuery) == 0 { 
        return renderer.Render(c.Response().Writer, "list.html", map[string]interface{}{
            "product":    searchQuery,
            "categories": categories,
			"discountedProduct": discountProduct,
            "errorText":  "Produk tidak ditemukan", 
        }, c)
    }
    return renderer.Render(c.Response().Writer, "list.html", map[string]interface{}{
        "product":    searchQuery,
        "categories": categories,
		"discountedProduct": discountProduct,
        "errorText":  "", 
    }, c)
}

