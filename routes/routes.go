package routes

import (
	"jar-project/controller"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func Routes() *echo.Echo {
	//Login Register Validasi
	e := echo.New()
	var sessionStore = sessions.NewCookieStore([]byte("jare-key"))

	e.Use(session.Middleware(sessionStore))
	// Web
	e.Static("/static", "static")
	e.GET("/", controller.Home)
	e.GET("/c/:slug", controller.Category)
	e.GET("/p/:slug", controller.Product)
	e.GET("/dashboard/:id", controller.Dashboard)
	e.GET("/cart", controller.Cart) // unfinished mobile and price
	e.POST("/p/:slug", controller.AddCart)
	e.GET("/checkout", controller.Checkout) //unfinished
	e.GET("/wallet", controller.WalletPage) //unfinished
	e.GET("/cashback", controller.Cashback) //unfinished
	e.GET("/order", controller.Order)       //unfinished
	e.GET("/review", controller.Review)     //unfinished

	//user
	e.POST("/dashboard/:id", controller.AddAddresses)
	e.GET("/forgot-password", controller.ForgotPasswordPage)
	e.POST("/forgot-password", controller.ForgotPassword)
	e.GET("/reset-password", controller.ResetPasswordPage)
	e.PUT("/reset-password", controller.ResetPassword)
	e.GET("/login", controller.Login)
	e.POST("/login", controller.LoginUser)
	e.GET("/logout", controller.Logout)
	e.GET("/register", controller.Register)
	e.POST("/register", controller.CreateUser)

	api := e.Group("/api")
	api.GET("/product", controller.GetProductApi)
	api.POST("/product", controller.CreateProductApi)
	api.PUT("/product/:slug", controller.UpdateProductApi)
	api.DELETE("/product/:slug", controller.DeleteProductApi)

	api.GET("/category", controller.GetCategoryApi)
	api.POST("/category", controller.CreateCategoryApi)
	api.PUT("/category/:slug", controller.UpdateCategoryApi)
	api.DELETE("/category/:slug", controller.DeleteCategoryApi)

	api.GET("/user", controller.GetUserApi)
	api.POST("/user", controller.CreateUserApi)
	api.PUT("/user/:id", controller.UpdateUserApi)
	api.DELETE("/user/:id", controller.DeleteUserApi)
	return e
}
