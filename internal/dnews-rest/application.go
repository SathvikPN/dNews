package dnewsrest

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/CloudyKit/jet/v6"
	"github.com/alexedwards/scs/v2"
)

type application struct {
	appName string
	server  *server
	debug   bool
	infoLog *log.Logger
	errLog  *log.Logger
	view    *jet.Set
	session *scs.SessionManager
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

	// init Jet Template Engine
	if app.debug {
		app.view = jet.NewSet(jet.NewOSFileSystemLoader("./web/views"), jet.InDevelopmentMode())
	} else {
		app.view = jet.NewSet(jet.NewOSFileSystemLoader("./web/views"))
	}

	// init session
	app.session = scs.New()
	app.session.Lifetime = 24 * time.Hour
	app.session.Cookie.Persist = true
	app.session.Cookie.Name = app.appName
	app.session.Cookie.Domain = app.server.host
	app.session.Cookie.SameSite = http.SameSiteStrictMode

	err := app.runServer()

	if err != nil {
		log.Fatal(err)
	}
}
