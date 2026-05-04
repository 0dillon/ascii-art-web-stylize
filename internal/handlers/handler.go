package handlers

import (
	"html/template"
	"log"
	"net/http"

	"ascii-art-web-stylize/internal/generator"
)

type pageData struct {
	Result string
	Text   string
}

// Handles the home page of the web server.
func Home(w http.ResponseWriter, r *http.Request) {
	// Parses the HTML template for use.
	tmp, err := template.ParseFiles("./templates/home.html")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Writes to the webpage.
	err = tmp.Execute(w, nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// Handles the ascii-art page of the web server.
func AsciiArt(w http.ResponseWriter, r *http.Request) {
	// Collects data from our HTML form.
	InputText := r.FormValue("inputText")
	bannerType := r.FormValue("bannerType")

	// Generates ASCII version of our data.
	result := generator.AsciiGen(InputText, bannerType)

	// Error handling for bad request.
	if InputText == "" || result == "" {
		http.Error(w, "400 bad request: input text is empty or non-printable characters", http.StatusBadRequest)
		return
	}

	// Initializing the struct with data from user.
	data := pageData{
		Result: result,
		Text:   InputText,
	}

	tmp, err := template.ParseFiles("./templates/home.html")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmp.Execute(w, data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
}
