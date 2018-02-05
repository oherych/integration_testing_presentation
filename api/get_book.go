package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/labstack/gommon/log"
	"github.com/oherych/integration_testing_presentation/model"
)

func (a *service) getBookHandler(w http.ResponseWriter, r *http.Request) {
	bookID := chi.URLParam(r, "book_id")

	var data model.Book
	db := a.postgress.Where("book_id = ?", bookID).Find(&data)
	if db.RecordNotFound() {
		log.Error("Not found")
		http.Error(w, http.StatusText(404), 404)
		return
	}

	if err := db.Error; err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	render.JSON(w, r, data)
}
