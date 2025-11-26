package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	html := `
		<!DOCTYPE html>
		<html>
		<head>
			<meta charset="UTF-8">
			<title>Mi web en Go</title>
		</head>
		<body>
			<h1>Soy alumno de la UOC</h1>
			<p>Esta es mi primera página web con Go.</p>
			<img src="/static/imagen.JPG" alt="Imagen" style="max-width:300px;">
		</body>
		</html>
	`

	fmt.Fprintln(w, html)
}

func main() {
	// Asociamos la ruta "/" con la función homeHandler
	http.HandleFunc("/", homeHandler)
	
	// Servir archivos estáticos desde "/static"
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	
	// Puerto donde escuchará el servidor
	port := "8080"
	log.Println("Servidor escuchando en http://localhost:" + port)

	// Arrancamos el servidor
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}