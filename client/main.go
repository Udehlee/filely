package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func index(w http.ResponseWriter, req *http.Request) {

	err := tmpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		log.Println("template didn't execute", nil)
	}
}

// upload  uploads and  send  file to a tcp server
func upload(w http.ResponseWriter, req *http.Request) {
	file, fileHeader, err := req.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	//ReadAll reads the content of the file
	fileInfo, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//  Writes write the fileInfo to the dst
	dst, err := os.Create(filepath.Join("./uploads/", fileHeader.Filename))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = dst.Write(fileInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// // Open the uploaded file
	uploadedFile, err := os.Open("./uploads/" + fileHeader.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer uploadedFile.Close()

	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	// Copy the uploadedfile data to the TCP connection

	_, err = io.Copy(conn, uploadedFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "success.gohtml", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {

	mux := http.NewServeMux()

	mux.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))

	mux.HandleFunc("/", index)

	mux.HandleFunc("/upload", upload)

	fmt.Println("listening on port :8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
