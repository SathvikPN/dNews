package dnewsrest

import (
	"fmt"
	"net/http"
	"time"
)

func (a *application) runServer() error {
	hostAddress := fmt.Sprintf("%s:%s", a.server.host, a.server.port)
	server := http.Server{
		Addr:        hostAddress,
		Handler:     a.routes(),
		ReadTimeout: 30 * time.Second,
	}

	a.infoLog.Printf("Application REST server listening on '%s'\n", hostAddress)

	return server.ListenAndServe()
}
