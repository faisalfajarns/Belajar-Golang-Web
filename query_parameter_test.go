package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//menangkap query parameter

func SayHello(writer http.ResponseWriter, request *http.Request) {
	name :=	request.URL.Query().Get("name")

	if name == "" {
		fmt.Fprint(writer, "Hello")
	}else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Faisal", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()

	body,  err  := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}else {
		fmt.Println(string(body))
	}
}