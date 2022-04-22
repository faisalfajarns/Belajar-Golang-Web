package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request){
	fmt.Fprintln(writer, "Hello World")
}


func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	
	//untuk membaca semua isi respons
	body, err := io.ReadAll(response.Body) // return dari readall adalah byte dan error. karena ada return byte maka harus dikonversi kebentuk string dengan string(variable) 

	if err != nil {
		panic(err)
	}else {
		result := string(body)
		fmt.Println(result)
	}
}