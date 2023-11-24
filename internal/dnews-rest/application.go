package dnewsrest

import (
	"fmt"
	"log"
	"os"

	"github.com/CloudyKit/jet/v6"
)

type application struct {
	appName string
	server  *server
	debug   bool
	infoLog *log.Logger
	errLog  *log.Logger
	view    *jet.Set
}

type server struct {
	host string
	port string
}

// Entrypoint for dNews application
func StartApplication(version string) {

	appServer := &server{
		host: "localhost",
		port: "9200",
	}

	app := &application{
		server:  appServer,
		appName: "DNews",
		debug:   true,
		infoLog: log.New(os.Stdout, "INFO\t", log.Ltime|log.Ldate|log.Lshortfile),
		errLog:  log.New(os.Stderr, "ERROR\t", log.Ltime|log.Ldate|log.Llongfile),
	}

	fmt.Printf("Starting DNews application...	version '%s'\n", version)

	if app.debug {
		app.view = jet.NewSet(jet.NewOSFileSystemLoader("./web/views"), jet.InDevelopmentMode())
	} else {
		app.view = jet.NewSet(jet.NewOSFileSystemLoader("./web/views"))
	}

	err := app.runServer()

	if err != nil {
		log.Fatal(err)
	}
}
