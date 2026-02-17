package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

// PageData holds the data to be rendered in the template
type PageData struct {
	Time      string
	UserAgent string
	OS        string
	Browser   string
	BgColor   string
	TextColor string
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handler)
	log.Println("Server executing on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	userAgent := r.Header.Get("User-Agent")
	os, browser, bgColor, textColor := detectDetails(userAgent)

	data := PageData{
		Time:      time.Now().Format("2006-01-02 15:04:05"),
		UserAgent: userAgent,
		OS:        os,
		Browser:   browser,
		BgColor:   bgColor,
		TextColor: textColor,
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
