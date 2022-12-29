package controller

import (
	"marllef/beautiful-api/internal/app/interfaces/presenter"
	"marllef/beautiful-api/internal/app/services"
	logger "marllef/beautiful-api/pkg/mylogger"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type BibleController interface {
	GetAll(res http.ResponseWriter, req *http.Request)
	GetOne(res http.ResponseWriter, req *http.Request)
	GetSingleVerse(res http.ResponseWriter, req *http.Request)
}

type bibleController struct {
	service services.BibleService
	log     *logger.Logger
	// encoder *json.Encoder
	BibleController
}

func NewBibleController(service services.BibleService) *bibleController {
	return &bibleController{
		service: service,
		log:     logger.Default(),
	}
}

func (c *bibleController) GetSingleVerse(res http.ResponseWriter, req *http.Request) {
	response := presenter.NewHttpPresenter(res)
	params := mux.Vars(req)

	book := strings.ToLower(params["book"])

	chapter, err := strconv.ParseInt(params["chapter"], 10, 32)
	if err != nil {
		response.Status(500)
		return
	}

	verse, err := strconv.ParseInt(params["verse"], 10, 32)
	if err != nil {
		response.Status(500)
		return
	}

	txtVerse := c.service.GetSingleVerse(book, chapter, verse)

	response.Send(txtVerse)
}
