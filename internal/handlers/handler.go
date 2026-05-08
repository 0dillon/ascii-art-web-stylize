package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"ascii-art-web-stylize/internal/generator"
)

type pageData struct {
	Result    template.HTML
	Text      string
	SubString string // Added to hold the substring value
	Style     template.HTMLAttr
}

func Home(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("./templates/home.html")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = tmp.Execute(w, nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func AsciiArt(w http.ResponseWriter, r *http.Request) {
	inputText := r.FormValue("inputText")
	bannerType := r.FormValue("bannerType")
	textColor := r.FormValue("textColor")
	subString := r.FormValue("substring")

	if textColor == "" {
		textColor = "#ffffff" // Fallback to white
	}

	rawResult := generator.AsciiGen(inputText, bannerType, subString, textColor)

	if inputText == "" || rawResult == "" {
		http.Error(w, "400 bad request: input text is empty or non-printable characters", http.StatusBadRequest)
		return
	}

	data := pageData{
		Text:      inputText,
		SubString: subString, // Pass the value back to the template
		Result:    template.HTML(rawResult), 
	}

	// Logic for coloring:
	if subString == "" {
		data.Style = template.HTMLAttr(fmt.Sprintf(`style="--ascii-color: %s;"`, textColor))
	} else {
		data.Style = template.HTMLAttr(`style="--ascii-color: var(--terminal-text);"`)
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