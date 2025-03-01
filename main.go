package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

func main() {
	// Serve static files from the "assets" directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	// Handle the root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var filepath = path.Join("view", "index.html")
		var tmpl, err = template.ParseFiles(filepath)
		if err != nil {
			http.Error(w, "Error parsing template: "+err.Error(), http.StatusInternalServerError)
			return
		}

		var data = map[string]interface{}{
			"title": "Learning Golang Web",
			"name":  "Batman",
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
			return
		}
	})

	fmt.Println("Server started at localhost:9000")
	// Start the server
	if err := http.ListenAndServe(":9000", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
