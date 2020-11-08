package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func getRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("officework rest"))
	})
	r.Route("/rest", func(r chi.Router) {
		r.Route("/client", func(r chi.Router) {
			r.Route("/0", func(r chi.Router) {
				r.Route("/{filterStr}", func(r chi.Router) {
					r.Use(FilterCtx)
					r.Post("/", getClientList)
				})
				r.Route("/detail/{filterStr}", func(r chi.Router) {
					r.Use(FilterCtx)
					r.Post("/", getClientDetail)
				})
				r.Route("/delete/{filterStr}", func(r chi.Router) {
					r.Use(FilterCtx)
					r.Post("/", getClientDelete)
				})
				r.Route("/save/{filterStr}", func(r chi.Router) {
					r.Use(FilterCtx)
					r.Post("/", getClientSave)
				})
			})
			r.Route("/1", func(r chi.Router) {
				r.Route("/{filterStr}", func(r chi.Router) {
					r.Use(FilterCtx)
					r.Post("/", getClientList)
				})
				r.Route("/detail/{filterStr}", func(r chi.Router) {
					r.Use(FilterCtx)
					r.Post("/", getClientDetail)
				})
				r.Route("/delete/{filterStr}", func(r chi.Router) {
					r.Use(FilterCtx)
					r.Post("/", getClientDelete)
				})
				r.Route("/save/{filterStr}", func(r chi.Router) {
					r.Use(FilterCtx)
					r.Post("/", getClientSave)
				})
			})
		})
	})

	return r
}
