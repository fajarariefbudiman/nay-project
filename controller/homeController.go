package controller

import (
	"fmt"
	"html/template"
	"io"
	"jar-project/model"
	"net/http"
	"strconv"

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
		}, c)
	}
	search := c.QueryParam("search")
	strsearch, _ := strconv.ParseBool(search)
	if strsearch {
		fmt.Println(strsearch)
	}
	result, err := model.SearchProduct(search)
	if err != nil {
		return renderer.Render(c.Response().Writer, "index.html", map[string]interface{}{
			"user":      nil,
			"errorText": "Tidak Produk Yang Sesuai",
			"result":    result,
		}, c)
	}
	return renderer.Render(c.Response().Writer, "index.html", map[string]interface{}{
		"categories": category,
	}, c)
}
