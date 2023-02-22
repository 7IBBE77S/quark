package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/7IBBE77S/web-app/internal/config"
	"github.com/7IBBE77S/web-app/internal/handlers"
	"github.com/7IBBE77S/web-app/internal/render"
	"log"
	"net/http"
	"time"
)
type portColor struct {
	BrightRed string
	Red       string
	Green     string
	Yellow    string
	Blue      string
	Magenta   string
	Purple    string
	Cyan      string
	White     string
	Grey      string
	Clean     string
}

var colors = portColor{
	BrightRed: "\033[1;31m",
	Red:       "\x1b[31m",
	Green:     "\x1b[32m",
	Yellow:    "\x1b[33m",
	Blue:      "\x1b[34m",
	Magenta:   "\x1b[35m",
	Purple:    "\x1b[36m",
	Cyan:      "\x1b[36m",
	White:     "\x1b[37m",
	Grey:      "\x1b[90m",
	Clean:     "\x1b[0m",
}
// var reader = bufio.NewReader(os.Stdin)

// var port, _ = getPortInput()
const port = ":8080"

var app config.AppConfig
var session *scs.SessionManager

//TODO: REMOVE 3RD PARTY PACKAGES AND REPLACE WITH OWN CODE!
func main() {
		// change this to true when in production
		app.InProduction = false

		// set up the session
		session = scs.New()
		session.Lifetime = 24 * time.Hour
		session.Cookie.Persist = true
		session.Cookie.SameSite = http.SameSiteLaxMode
		session.Cookie.Secure = app.InProduction
	
		app.Session = session
	
		tc, err := render.CreateTemplateCache()
		if err != nil {
			log.Fatal("cannot create template cache")
		}
	
		app.TemplateCache = tc
		app.UseCache = false
	
		repo := handlers.NewRepo(&app)
		handlers.NewHandlers(repo)
	
		render.NewTemplates(&app)
	
		fmt.Println(fmt.Sprintf("Staring application on port %s", port))
	
		srv := &http.Server{
			Addr:    port,
			Handler: routes(&app),
		}
	
		err = srv.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}

	}