package main

import (
	"io"
	//"log"
	"net/http"
	"fmt"
	"io/ioutil"
	//"html/template"
)
/*
var tpl *template.Template
func init() {
	tpl = template.Must(template.ParseGlob("tempates/*"))
}
*/
type person struct {
	FirstName string
	LastName string
	Subscribed bool
}

func main() {
	http.HandleFunc("/",ReadFiles)
	http.HandleFunc("/set",set)
	http.HandleFunc("/read",read)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)
}

func set(w http.ResponseWriter,req *http.Request) {
	http.SetCookie(w,&http.Cookie{
		Name: "my-cookie",
		Value: "some value",
	})
	fmt.Fprintln(w, "COOKIE WRITTEN-CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome goto: dev tools / application / cookies")
}

func read(w http.ResponseWriter,req *http.Request) {
	c,err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(w,err.Error(),http.StatusNoContent)
		return
	}
	fmt.Fprintln(w,"YOUR COOKIE: ",c) 
}

/*
func PersonDataRead(w http.ResponseWriter , req  *http.Request) {
	bs := make([]byte, req.ContentLength)
	req.Body.Read(bs)
	body := string(bs)
	err := tpl.ExecuteTemplate(w, "index.gohtml",tpl)
	if err != nil {
		http.Error(w,err.Error(),500)
		log.Fatalln(err)
	}
	fmt.Println(body)
}
*/

func ReadFiles(w http.ResponseWriter , req *http.Request) {
	var s string
	fmt.Println(req.Method)
	if req.Method == http.MethodPost {
		f,h,err := req.FormFile("q")
		if err != nil {
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return
		}
		defer f.Close()
		 
		//logging Information to Console
		fmt.Println("\nFile : ",f,"\nHeader : ",h,"\nerr : ",err)
		
		//read
		bs,err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return
		}
		s = string(bs)
	}
	w.Header().Set("Content-Type","text/html; charset=utf-8")
	io.WriteString(w,`
	<form method="POST"  enctype="multipart/form-data">
	<input type="file" name="q">
	<input type="submit">
	</form>
	<br>`+s)
	
}
