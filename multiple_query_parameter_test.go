package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func MultipleQueryParameter(writer http.ResponseWriter, request *http.Request) {
	first_name := request.URL.Query().Get("first_name")
	last_name := request.URL.Query().Get("last_name")
	
		fmt.Fprintf(writer, "Hello %s %s", first_name, last_name)
	
}

func TestMultipleQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?first_name=Faisal&last_name=Fajar", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParameter(recorder, request)

	response := recorder.Result()

	body,  err  := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}else {
		fmt.Println(string(body))
	}
}