package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/purna25/go-webapp/pkg/config"
	"github.com/purna25/go-webapp/pkg/handlers"
	"github.com/purna25/go-webapp/pkg/render"
)

const portNumber = ":8500"

func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting application on %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
