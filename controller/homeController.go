package controller

import (
	"html/template"
	"io"
	"jar-project/model"
	"net/http"
	"net/url"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	Template *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}
	return t.Template.ExecuteTemplate(w, name, data)
}

func Home(c echo.Context) error {
	renderer := &TemplateRenderer{
		Template: template.Must(template.ParseGlob("./template/*.html")),
	}
	category, err := model.GetAllCategories()
	if err != nil {
		return err
	}
	discountProduct, err := model.GetProductDiscount()
	if err != nil {
		return err
	}
	search := c.QueryParam("search")
	if search != "" {
		u, _ := url.Parse("/p")
		q := u.Query()
		q.Set("search", search)
		u.RawQuery = q.Encode()
		return c.Redirect(http.StatusSeeOther, u.String())
	}
	sess, _ := session.Get("session", c)
	isAuthenticated := true
	if userID, ok := sess.Values["user_id"].(int); ok {
		user, err := model.GetUserByID(userID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return renderer.Render(c.Response().Writer, "index.html", map[string]interface{}{
			"Auth":       isAuthenticated,
			"user":       user,
			"categories": category,
			"discountedProduct": discountProduct,
			
		}, c)
	}
	return renderer.Render(c.Response().Writer, "index.html", map[string]interface{}{
		"categories": category,
		"discountedProduct": discountProduct,
	
	}, c)
}


