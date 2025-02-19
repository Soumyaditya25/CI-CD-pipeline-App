package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func root(w http.ResponseWriter, e *http.Request){
//sample login check if Jenkins works take 4
	person := Person{
		Name:  "Soumyaditya",
		Age:   20,
		Email: "soumyaditya@example.com",
	}

	jsonData, err := json.Marshal(person)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(jsonData)
}

func main(){
	http.HandleFunc("/", root)
	fmt.Println("Serving at port 3001")
	err := http.ListenAndServe(":3001", nil)
	if err != nil{
		fmt.Println("Error Occured!")
	}
	
}