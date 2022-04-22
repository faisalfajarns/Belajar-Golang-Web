package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(w http.ResponseWriter, req *http.Request) {
	content_type := req.Header.Get("content-type")
	fmt.Fprint(w, content_type)
}

func TestHeader(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil)
	req.Header.Add("content-type", "application/json")

	rec := httptest.NewRecorder()

	RequestHeader(rec, req)

	res := rec.Result()

	body, err := io.ReadAll(res.Body)
	if err != nil{
		panic(err)
	}else {
		fmt.Println(string(body))
	}
}

func ResponseHeader(w http.ResponseWriter, req *http.Request){
	w.Header().Add("X-Powered-By", "Faisal Fajar")
	fmt.Fprint(w, "Ok")
}

func TestResponseHeader(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil)
	req.Header.Add("content-type", "application/json")

	rec := httptest.NewRecorder()

	ResponseHeader(rec, req)

	res := rec.Result()

	body, err := io.ReadAll(res.Body)
	if err != nil{
		panic(err)
	}else {
		fmt.Println(string(body))
		fmt.Println(res.Header.Get("X-Powered-By"))
	}
}