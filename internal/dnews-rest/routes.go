package dnewsrest

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *application) routes() http.Handler {
	// http multiplexer
	router := chi.NewRouter()
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		responseBody := fmt.Sprintln("Application Name:", a.appName)
		_, err := w.Write([]byte(responseBody))
		if err != nil {
			a.errLog.Println("failed to write response body, URI:", r.URL)
		}
	})
	return router
}
