package main

import (
	"fmt"
	"html/template"
	"net/http"
	"math/rand"
)

var numbers = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

func generateLuckyNum() int {
	return rand.Intn(len(numbers))
}

func translateLuckyNum(num int) string {
	if num < 0 || num >= len(numbers) {
		return ""
	}
	
	return numbers[num]
}

func greetingHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	
	funcMap := template.FuncMap {
		"generate": generateLuckyNum,
		"translate": translateLuckyNum,
	}
	
	type Info struct {
		Name string
	}
	
	info := Info{name}
	
	t, err := template.New("").Funcs(funcMap).ParseFiles("template.html")
	
	if err != nil {
		fmt.Print(err)
		return
	}
	
	err = t.ExecuteTemplate(w, "template.html", info)
	
	if err != nil {
		fmt.Print(err)
		return
	}
}

func main() {
	http.HandleFunc("/greeting", greetingHandler)
	http.ListenAndServe(":8080", nil)
}