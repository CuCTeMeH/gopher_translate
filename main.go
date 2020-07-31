package main

import (
	"flag"
	"github.com/CuCTeMeH/gopher_translate/http"
	"github.com/CuCTeMeH/gopher_translate/translator"
	"github.com/sirupsen/logrus"
)

var port = flag.Int("port", 8080, "")

func main() {
	flag.Parse()

	translator.InitHistoryStorage()

	err := http.InitServer(*port)
	if err != nil {
		logrus.Fatal(err)
	}
}
