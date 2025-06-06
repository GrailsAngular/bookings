package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/GrailsAngular/bookings/pkg/config"
	"github.com/GrailsAngular/bookings/pkg/handlers"
	"github.com/GrailsAngular/bookings/pkg/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	//fmt.Println("Hello World!")
	/* 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	   		n, err := fmt.Fprintf(w, "Hello, world!")
	   		if err != nil {
	   			fmt.Println(err)
	   		}
	   		fmt.Println("Bytes written:", n)
	   	})
	*/

	app.InProduction = false

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

	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)
	fmt.Printf("Starting application on port %s\n", portNumber)
	//http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
