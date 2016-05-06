package main

import(

	//"html/template"
	"net/http"

)

func main(){

	http.ListenAndServe("localhost:9000",h())



}

func h() http.Handler{

	return 	http.H("/test/", f)
}

func f(r http.Response, req *http.Request){



}