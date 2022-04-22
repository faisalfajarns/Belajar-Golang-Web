package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func MultipleValueQueryParameter(writer http.ResponseWriter, request *http.Request){
	query := request.URL.Query()

	names := query["name"]
	fmt.Fprint(writer, strings.Join(names, " "))
}

func TestMultipleValueQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Faisal&name=Fajar&name=Nursaid", nil)
	recorder := httptest.NewRecorder()

	MultipleValueQueryParameter(recorder, request)

	response := recorder.Result()

	body,  err  := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}else {
		fmt.Println(string(body))
	}
}