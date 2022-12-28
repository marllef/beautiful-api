package controller

import (
	"encoding/json"
	"marllef/beautiful-api/internal/app/interfaces/presenter"
	"marllef/beautiful-api/internal/app/services"
	logger "marllef/beautiful-api/pkg/mylogger"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserController interface {
	GetAll(res http.ResponseWriter, req *http.Request)
	GetOne(res http.ResponseWriter, req *http.Request)
}

type userController struct {
	service services.UserService
	log     *logger.Logger
	encoder *json.Encoder
	ProductController
}

func NewUserController(service services.UserService) *userController {
	return &userController{
		service: service,
		log:     logger.Default(),
	}
}

func (c *userController) GetAll(res http.ResponseWriter, req *http.Request) {
	response := presenter.NewHttpPresenter(res)

	products, err := c.service.GetAllUsers()
	if err != nil {
		response.Status(404).Error(err)
		return
	}

	response.Status(200).Json(products)
	return
}

func (c *userController) GetOne(res http.ResponseWriter, req *http.Request) {
	response := presenter.NewHttpPresenter(res)
	params := mux.Vars(req)

	id, err := strconv.ParseInt(params["id"], 10, 32)
	if err != nil {
		response.Status(500)
		return
	}

	product, err := c.service.GetOneUser(id)
	if err != nil {
		c.log.Errorf("Failed to get products: %v", err)
		response.Status(404).Error(err)
		return
	}

	response.Status(200).Json(product)
	return
}
