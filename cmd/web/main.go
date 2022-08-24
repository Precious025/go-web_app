package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Precious025/go-web_app/pkg/config"
	"github.com/Precious025/go-web_app/pkg/handler"
	"github.com/Precious025/go-web_app/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false

	session = scs.New()
	session.Lifetime = time.Hour * 24 * 7
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplate()

	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handler.NewRepo(&app)

	handler.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("This is starting server on port%s\n", portNumber)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
