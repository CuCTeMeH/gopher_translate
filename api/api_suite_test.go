package api

import (
	"bytes"
	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestAPIRequest struct {
	Recorder *httptest.ResponseRecorder
	hasRun   bool
	headers  map[string]string
	Request  *http.Request
	handler  http.Handler
}

func handler() http.Handler {
	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(func(http.ResponseWriter, *http.Request) {})
	router.HandleFunc("/word", postWord).Methods("POST").Name("get_history")
	router.HandleFunc("/sentence", postSentence).Methods("POST").Name("get_history")
	router.HandleFunc("/history", getHistory).Methods("GET").Name("get_history")

	return router
}

func (r *TestAPIRequest) run() {
	if r.hasRun {
		return
	}
	r.handler.ServeHTTP(r.Recorder, r.Request)
	r.hasRun = true
}

func ExpectRequest(handler http.Handler, method, path string, bodyString string) *TestAPIRequest {
	var body io.Reader
	if bodyString != "" {
		body = bytes.NewReader([]byte(bodyString))
	}
	Request, err := http.NewRequest(method, path, body)
	if err != nil {
		panic(err)
	}
	recorder := httptest.NewRecorder()
	return &TestAPIRequest{
		Recorder: recorder,
		headers:  map[string]string{},
		Request:  Request,
		handler:  handler,
	}
}

func TestAPI(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "API Suite")
}
