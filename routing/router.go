package routing

import (
	"dev/interview-craft/products"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

type Router interface {
	AttachRoutes(listener *echo.Echo)
	AddProduct(c echo.Context) error
	DisplayInventory(c echo.Context) error
}

type RouterManager struct {
	Manager products.ProductManager
}

func NewRouter(mgr products.ProductManager) Router {
	return &RouterManager{
		Manager: mgr,
	}
}

func (r *RouterManager) AttachRoutes(listener *echo.Echo) {
	listener.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	listener.POST("/addProduct", r.AddProduct)
	listener.GET("/inventory", r.DisplayInventory)
}

func (r *RouterManager) AddProduct(c echo.Context) error {
	log.Info("router-add-product")
	defer log.Info("router-add-product-complete")

	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Error("router-add-product-error ", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	var product products.ProductDetails

	err = json.Unmarshal(body, &product)
	if err != nil {
		log.Error("router-add-product-json-unmarshal-error ", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	err = r.Manager.AddNewProduct(product)
	if err != nil {
		log.Error("router-add-product-mongo-add-product-error ", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, nil)
}

func (r *RouterManager) DisplayInventory(c echo.Context) error {
	log.Info("router-display-inventory")
	defer log.Info("router-display-inventory-complete")

	products, err := r.Manager.GetInventory()
	if err != nil {
		log.Error("router-get-inventory-error ", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, products)
}
