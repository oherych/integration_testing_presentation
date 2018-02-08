package api

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/labstack/gommon/log"
	"github.com/oherych/integration_testing_presentation/model"
)

func (a *service) getBookListHandler(w http.ResponseWriter, r *http.Request) {
	var data []model.Book
	err := a.postgress.Find(&data).Error
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	render.JSON(w, r, data)
}
