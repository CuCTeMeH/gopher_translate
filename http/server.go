package http

import (
	"fmt"
	"github.com/CuCTeMeH/gopher_translate/api"
	"github.com/go-chi/chi"
	"net/http"
)

func InitServer(port int) error {
	router := chi.NewRouter()
	router.Mount("/", api.Routes())

	address := fmt.Sprintf(":%d", port)
	server := http.Server{Addr: address, Handler: router}
	err := server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
