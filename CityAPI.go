package main

import (
	"encoding/json"
	"net/http"
	"time"
	"strconv"
	"log"
)

type city struct {
	Name string
	Area uint64
}


//Middleware to check content types as JSON
func filterContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
		log.Println("Checking the content inside Middleware")
		// Filter requests by MIME type
		if r.Header.Get("Content-type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("415 - Unsupported Media Type. Please Send JSON\n"))
			return
		}
		handler.ServeHTTP(w,r)
	})
}


//Middleware to add server timestamp for response cookie
func setServerTimeCookie(handler http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w,r)
		cookie := http.Cookie{Name: "Server-Time(UTC)",Value: strconv.FormatInt(time.Now().Unix(),10)}
		http.SetCookie(w,&cookie)
		log.Println("Currently setting the  server time")
	})
}

func mainLogic(w http.ResponseWriter,r *http.Request) {
	if r.Method == "POST" {
		var tempCity city
		decoder := json.NewDecoder(r.Body)
		err:= decoder.Decode(&tempCity)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()
		log.Printf("Got %s city with area of %d sq miles!\n",tempCity.Name,tempCity.Area)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("201- Created\n"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405- Method Not Allowed\n"))
	}
}

func main() {
	mainLogicHandler := http.HandlerFunc(mainLogic)
	http.Handle("/city",filterContentType(setServerTimeCookie(mainLogicHandler)))
	http.ListenAndServe(":8000", nil)
}
/*
func main() {
	mainLogicHandler := http.HandlerFunc(mainLogic)
	http.HandleFunc("/city",filterContentType(setServerTimeCookie(mainLogicHandler)))
	
	http.ListenAndServe(":8000",nil)
}*/
