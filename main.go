package main

import (
	"net/http"

	"project/handlers"
	"project/middleware"
	"project/storage"
	"project/logger"
)

func main() {

	logger.Init(true)

	st := &storage.UserStorage{
		FileName: "data/users.json",
	}

	h := &handlers.UserHandler{
		Storage: st,
	}

	mux := http.NewServeMux()

	mux.Handle(
		"/users",
		middleware.Logging(
			middleware.Auth(
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					if r.Method == http.MethodGet {
						h.GetUsers(w, r)
						return
					}

					if r.Method == http.MethodPost {
						h.CreateUser(w, r)
						return
					}

					http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				}),
			),
		),
	)

	mux.Handle(
		"/users/",
		middleware.Logging(
			middleware.Auth(
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					if r.Method == http.MethodGet {
						h.GetUserByID(w, r)
						return
					}

					if r.Method == http.MethodPut {
						h.UpdateUser(w, r)
						return
					}

					http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				}),
			),
		),
	)


	http.ListenAndServe(":8080", mux)
}
