package api

import (
	"encoding/json"
	"github.com/CuCTeMeH/gopher_translate/translator"
	"github.com/go-chi/render"
	"net/http"
)

func postWord(w http.ResponseWriter, r *http.Request) {
	var word map[string]string

	err := json.NewDecoder(r.Body).Decode(&word)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	gopherWord, err := translator.TranslateWord(word["english-word"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := map[string]string{"gopher-word": gopherWord}
	render.JSON(w, r, result)
}

func postSentence(w http.ResponseWriter, r *http.Request) {
	var sentence map[string]string

	err := json.NewDecoder(r.Body).Decode(&sentence)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	gopherSentence, err := translator.TranslateSentence(sentence["english-sentence"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := map[string]string{"gopher-sentence": gopherSentence}
	render.JSON(w, r, result) // A chi router helper for serializing and returning json
}
