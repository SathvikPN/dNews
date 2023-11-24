package dnewsrest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func (a *application) routes() http.Handler {
	// http multiplexer
	router := chi.NewRouter()
	// A good base middleware stack
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		err := a.render(w, r, "index", nil)
		if err != nil {
			log.Fatalln(err)
		}
	})

	router.Get("/welcome", func(w http.ResponseWriter, r *http.Request) {
		responseBody := fmt.Sprintln(a.appName, "welcomes you!")
		_, err := w.Write([]byte(responseBody))
		if err != nil {
			a.errLog.Println("failed to write response body, URI:", r.URL)
		}
	})
	return router
}
