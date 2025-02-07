package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"rps/cmd/rps"
	"strconv"
)

type Player struct {
	Name string
}

// Creando jugador
var player Player

const (
	templateDir  = "cmd/templates/"
	templateBase = templateDir + "base.html"
)

func Index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", nil)
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "new-game.html", nil)
}

func Game(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		// Leer los datos del formulario
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		player.Name = r.Form.Get("name")
	}

	// Redirecíon a otra ruta
	if player.Name == "" {
		http.Redirect(w, r, "/new", http.StatusFound)
	}

	renderTemplate(w, "game.html", player)
}

func Play(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := rps.PlayRound(playerChoice)

	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func About(w http.ResponseWriter, r *http.Request) {
	restartValue()
	renderTemplate(w, "about.html", nil)
}

// Manejo de páginas de error
var errorTemplates = template.Must(template.ParseGlob("cmd/templates/**/*.html"))

func handlerError(w http.ResponseWriter, name string, status int) {
	w.WriteHeader(status)
	errorTemplates.ExecuteTemplate(w, name, nil)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	// Devolver un error personalizado para páginas no encontradas
	handlerError(w, "404", http.StatusNotFound)
}

func renderTemplate(w http.ResponseWriter, page string, data any) {
	tpl := template.Must(template.ParseFiles(templateBase, templateDir+page))

	err := tpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tpl.Execute(w, nil)
}

// Reiniciar valores
func restartValue() {
	player.Name = ""
	rps.ComputerScore = 0
	rps.PlayerScore = 0
}
