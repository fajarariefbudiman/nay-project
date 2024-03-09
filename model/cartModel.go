package model

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Cart struct {
	Items map[int]CartItem
}

type CartItem struct {
	ProductID int
	Quantity  int
	Price     int
}

func NewCart() *Cart {
	return &Cart{
		Items: make(map[int]CartItem),
	}
}

func GetCart(c echo.Context) (*Cart, error) {
	cookie, err := c.Cookie("cart")
	if err != nil {
		if err == http.ErrNoCookie {
			return NewCart(), nil
		}
		return nil, err
	}
	var cart Cart
	if err := json.Unmarshal([]byte(cookie.Value), &cart); err != nil {
		return nil, err
	}
	return &cart, nil
}

func SaveCart(c echo.Context, cart *Cart) error {
	cartJSON, err := json.Marshal(cart)
	if err != nil {
		return err
	}
	c.SetCookie(&http.Cookie{
		Name:  "cart",
		Value: string(cartJSON),
		Path:  "/",
	})
	return nil
}

// func AddToCart(c echo.Context) error {
// 	productId := c.FormValue("product_id")
// 	id, err := strconv.Atoi(productId)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 			"Message": "Invalid Product ID",
// 		})
// 	}
// 	quantity := c.FormValue("quantity")
// 	qty, err := strconv.Atoi(quantity)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 			"Message": "Invalid Quantity",
// 		})
// 	}
// 	price := c.FormValue("price")
// 	strPrice, err := strconv.Atoi(price)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 			"Message": "Invalid Price",
// 		})
// 	}
// 	cart, err := model.GetCart(c)
// 	if err != nil {
// 		cart = &model.Cart{
// 			Items: make(map[int]model.CartItem),
// 		}
// 	}
// 	cart.Items[id] = model.CartItem{
// 		ProductID: id,
// 		Quantity:  qty,
// 		Price:     strPrice,
// 	}
// 	model.SaveCart(c, cart)
// 	return c.Redirect(http.StatusSeeOther, "/cart")
// }
