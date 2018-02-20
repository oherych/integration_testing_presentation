package api

import (
	"net/http"

	"github.com/satori/go.uuid"

	"github.com/labstack/gommon/log"
	"github.com/oherych/integration_testing_presentation/model"

	"github.com/go-chi/render"
)

func (a *service) createBookHandler(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")

	book := model.Book{
		BookID: uuid.NewV4().String(),
		Name:   name,
	}

	err := a.postgress.Create(&book).Error
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	render.JSON(w, r, map[string]string{
		"status": "ok",
	})
}
