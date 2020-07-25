package api

import "github.com/go-chi/chi"

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/word", postWord)
	router.Post("/sentence", postSentence)

	router.Get("/history", getHistory)

	return router
}
