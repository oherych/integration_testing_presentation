package api

import (
	"github.com/go-chi/chi"
	"net/http"

	"github.com/labstack/gommon/log"
	"github.com/oherych/integration_testing_presentation/model"

	"github.com/go-chi/render"
)

func (a *service) createBookHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "book_id")
	name := r.PostFormValue("name")

	book := model.Book{
		BookID: id,
		Name:   name,
	}

	err := a.postgress.Assign(book).FirstOrCreate(&book).Error
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	render.JSON(w, r, map[string]string{
		"status": "ok",
	})
}
