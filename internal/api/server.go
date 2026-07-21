package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewServer() *chi.Mux {
	return chi.NewRouter()
}

func Start() {
	port := ":3000"
	r := NewServer()

	ConfigurateRouter(r)

	fmt.Println("Server is served on http://127.0.0.1" + port + "/")
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}
