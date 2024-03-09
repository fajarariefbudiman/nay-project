package controller

import (
	"html/template"
	"jar-project/model"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func WalletPage(c echo.Context) error {
	renderer := &TemplateRenderer{
		Template: template.Must(template.ParseGlob("./template/*.html")),
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
		return renderer.Render(c.Response().Writer, "wallet.html", map[string]interface{}{
			"Auth":      isAuthenticated,
			"addresses": address,
			"user":      user,
		}, c)
	}
	return c.Redirect(http.StatusSeeOther, "/login")
}
