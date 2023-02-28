// Test

package main

import (
	"net/http"

	"github.com/ub2013210024/hello/helpers"
)

// Create our handler functions
func (app *application) Greeting(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Welcome to my webpage"))
	question, err := app.question.Get()
	if err != nil {
		return
	}
	w.Write([]byte(question.Body))
}

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	helpers.RenderTemplates(w, "./static/html/home.page.tmpl")
}

func (app *application) About(w http.ResponseWriter, r *http.Request) {
	//day := time.Now().Weekday()
	//w.Write([]byte(fmt.Sprintf("Have a good %s.", day)))
}

func (app *application) MessageCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		//w.Write([]byte("Method not allowed"))
		w.Header().Set("Allow", "POST")
		//w.WriteHeader(405)
		http.Error(w, "Methold not allowed", http.StatusMethodNotAllowed)
		return
	}
}
