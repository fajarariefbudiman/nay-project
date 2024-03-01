package controller

import (
	"html/template"
	"jar-project/model"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

var otpStore = make(map[string]string)

func Login(c echo.Context) error {
	renderer := &TemplateRenderer{
		Template: template.Must(template.ParseGlob("./template/*.html")),
	}
	return renderer.Render(c.Response().Writer, "login.html", map[string]interface{}{
		"user": nil,
	}, c)
}

func ForgotPasswordPage(c echo.Context) error {
	renderer := &TemplateRenderer{
		Template: template.Must(template.ParseGlob("./template/*.html")),
	}
	return renderer.Render(c.Response().Writer, "forgot.html", map[string]interface{}{
		"user": nil,
	}, c)
}

func ResetPasswordPage(c echo.Context) error {
	renderer := &TemplateRenderer{
		Template: template.Must(template.ParseGlob("./template/*.html")),
	}
	return renderer.Render(c.Response().Writer, "reset-password.html", map[string]interface{}{
		"user": nil,
	}, c)
}

func Register(c echo.Context) error {
	renderer := &TemplateRenderer{
		Template: template.Must(template.ParseGlob("./template/*.html")),
	}
	return renderer.Render(c.Response().Writer, "register.html", map[string]interface{}{
		"name": "Dolly!",
	}, c)
}

func CreateUser(c echo.Context) error {
	firstName := c.FormValue("firstName")
	lastName := c.FormValue("lastName")
	email := c.FormValue("email")
	password := c.FormValue("password")
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hashed := string(hashedPassword)
	err := model.CreateUser(firstName, lastName, email, hashed)
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/register")
	}
	user, err := model.GetUserByEmail(email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["user_id"] = user.Id
	sess.Values["firstname"] = user.Firstname
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusSeeOther, "/c/men")
}

func LoginUser(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	result, err := model.AuthenticateUser(email, password)
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/login")
	}
	if !result {
		return c.Redirect(http.StatusSeeOther, "/login")
	}
	user, err := model.GetUserByEmail(email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["user_id"] = user.Id
	sess.Values["firstname"] = user.Firstname
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusSeeOther, "/c/men")
}

func Logout(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}
	sess.Options.MaxAge = -1
	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		return err
	}
	return c.Redirect(http.StatusSeeOther, "/c/men")
}

func ForgotPassword(c echo.Context) error {
	email := c.FormValue("email")
	rand.Seed(time.Now().UnixNano())
	otp := strconv.Itoa(rand.Intn(900000) + 100000)
	otpStore[email] = otp
	sendOTPEmail(email, otp)

	return c.Redirect(http.StatusSeeOther, "/reset-password")
}

func sendOTPEmail(email, otp string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "admin@example.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Password Reset OTP")
	m.SetBody("text/plain", "Your OTP for password reset: "+otp)

	d := gomail.NewDialer("smtp.gmail.com", 587, "budimanfajar660@gmail.com", "bwtoyhwtdjwugbwq")

	if err := d.DialAndSend(m); err != nil {
		log.Println("Failed to send OTP email: ", err)
	}
}

func ResetPassword(c echo.Context) error {
	email := c.FormValue("email")
	otp := c.FormValue("otp")
	newPassword := c.FormValue("newPassword")
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	hashed := string(hashedPassword)
	err := model.UpdatePassword(email, hashed)
	if err != nil {
		log.Println("Invalid Email or Password")
		return c.JSON(http.StatusBadRequest, err)
	}
	storedOTP, exists := otpStore[email]
	if !exists || storedOTP != otp {
		return c.JSON(http.StatusUnauthorized, "Invalid OTP.")
	}
	log.Printf("Password reset for email: %s. New password: %s\n", email, newPassword)
	delete(otpStore, email)
	return c.Redirect(http.StatusSeeOther, "/c/men")
}
