package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type Config struct {
	Port string `json:"port"`
}

type PageData struct {
	Time      string
	UserAgent string
	OS        string
	Browser   string
	BgColor   string
	TextColor string
}

func loadConfig() Config {
	var config Config
	file, err := os.Open("config.json")
	if err != nil {
		log.Println("Could not open config.json, using default port 8080")
		return Config{Port: "8080"}
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Println("Could not parse config.json, using default port 8080")
		return Config{Port: "8080"}
	}

	if config.Port == "" {
		config.Port = "8080"
	}
	return config
}

func main() {
	config := loadConfig()

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handler)
	log.Println("Server executing on http://localhost:" + config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
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
