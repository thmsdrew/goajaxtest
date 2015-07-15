package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Name struct {
	Fname string `json:"fname"`
	Lname string `json:"lname"`
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/ajax", Ajax)
	http.ListenAndServe(":8080", nil)
}

func Ajax(w http.ResponseWriter, r *http.Request) {
	r.ParseForm();
	fname := r.FormValue("fname");
	lname := r.FormValue("lname");
	
	data := Name{fname, lname}
	js, err := json.Marshal(data);
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	fmt.Println(fname, lname);
	
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}