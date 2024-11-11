package main

import (
	"automation/db"
	"automation/handlers"
	"fmt"
	"log"
	"net/http"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")                                // Разрешаем доступ всем доменам (или замените на конкретный домен)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE") // Указываем разрешенные методы
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")     // Разрешенные заголовки

		// Обрабатываем preflight-запрос OPTIONS
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer database.Close()

	// Применяем CORS middleware к обработчику /login
	http.Handle("/login", enableCORS(http.HandlerFunc(handlers.LoginHandler(database))))

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
