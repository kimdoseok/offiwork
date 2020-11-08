package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func FilterCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("CORS_SUPPORTS_CREDENTIALS", "True")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		filterStr := chi.URLParam(r, "filterStr")
		ctx := context.WithValue(r.Context(), "filterStr", filterStr)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func main() {
	db, err := connectSQL()
	if err != nil {
		log.Println(err)
	}
	db.LogMode(true)
	r := getRouter()
	fmt.Println("Starting backend on 8008...")
	http.ListenAndServe(":8008", r)
}
