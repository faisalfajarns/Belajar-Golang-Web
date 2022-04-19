package belajargolangweb

import (
	"fmt"
	"net/http"
	"testing"
)


func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		//logic web

		//notes Fprint dll untuk di cetak kedalam writer karena sudah terkonversi ke byte dengan Fprint
		fmt.Fprint(writer, "Hello World")
	}

	server := http.Server{
		Addr: "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}