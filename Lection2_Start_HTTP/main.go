package main

import (
	"fmt"
	"log"
	"net/http"
)

// w - responseWriter (куда записывать ответ)
// r = request (откуда брать запрос)

// Обработчик
func GetGreet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello, I'm new web server! </h1>")
	fmt.Println("Method:", r.Method)
	fmt.Println("URL", r.URL)
}

func main() {
	http.HandleFunc("GET /", GetGreet)           // Если придёт запрос по адресу(ресурсу) "/"(root), то вызывай функцию GetGreet
	log.Fatal(http.ListenAndServe(":8080", nil)) // Запускаем веб сервер для "слушания"
}
