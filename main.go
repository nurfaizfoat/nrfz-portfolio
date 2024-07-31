package main

import (
	"html/template"
	"log"
	"net/http"
)

type PageVariables struct {
	Title string
}

func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/about", AboutPage)
	http.HandleFunc("/works", WorksPage)
	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	vars := PageVariables{Title: "Home"}
	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, vars); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	vars := PageVariables{Title: "About"}
	tmpl, err := template.ParseFiles("templates/about.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, vars); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func WorksPage(w http.ResponseWriter, r *http.Request) {
	vars := PageVariables{Title: "Works"}
	tmpl, err := template.ParseFiles("templates/works.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, vars); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
