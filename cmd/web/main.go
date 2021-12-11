package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/kristiansigston/goappbookings/pkg/config"
	"github.com/kristiansigston/goappbookings/pkg/handlers"
	"github.com/kristiansigston/goappbookings/pkg/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {

	//change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("error cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers((repo))
	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %d", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

// http.HandleFunc("/divide", Divide)

// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 	n, err := fmt.Fprintf(w, "Hello, world")

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(fmt.Sprintf("Number of Bytes written %d", n))
// })

// func divideValues(x, y float32) (float32, error) {
// 	if y == 0 {
// 		err := errors.New("cannot divide by 0")
// 		return 0, err
// 	}
// 	result := x / y
// 	return result, nil
// }

// // adds two integers and returns the sum
// func addValues(x, y int) int {
// 	return x + y
// }

// sum := addValues(2, 2)
// _, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page with %d", sum))

// 	func Divide(w http.ResponseWriter, r *http.Request) {
// 	f, err := divideValues(100.0, 0.0)
// 	if err != nil {
// 		fmt.Fprintf(w, "Cannot divide by 0")
// 		return
// 	}

// 	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 10.0, f))
// }
