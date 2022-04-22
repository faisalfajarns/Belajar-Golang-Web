package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

//handler
func FormPost(w http.ResponseWriter, req *http.Request){

	// step form post 

	// 1. parsing data
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}

	//cara manual
	
	//2. get data post formnya 
	first_name := req.PostForm.Get("first_name")
	last_name := req.PostForm.Get("last_name")

	//cara cepat yang sudah terparsing
	// req.PostFormValue("first_name")
	// req.PostFormValue("last_name")

	fmt.Fprintf(w,"Hello %s %s", first_name, last_name )
}

func TestFormPost(t *testing.T) {
	req_body := strings.NewReader("first_name=Faisal&last_name=Fajar")
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080", req_body)
	
	//note : standar header untuk menggunakan form post : Content-Type" : "application/x-www-form-urlencoded
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	rec := httptest.NewRecorder()

	FormPost(rec, req)

	res := rec.Result()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}else {
		fmt.Println(string(body))
	}

}

