package api

import (
	"github.com/CuCTeMeH/gopher_translate/translator"
	"github.com/go-chi/render"
	"net/http"
)

func getHistory(w http.ResponseWriter, r *http.Request) {
	h := translator.Storage.GetOrderedMap()
	render.JSON(w, r, h)
}
