package internal

import (
	"embed"
	"html/template"
	"io/fs"
	"log"
	"net/http"

	"golang.org/x/xerrors"
)

//go:embed _embed/templates
var embedTmpl embed.FS
var tmpl fs.FS

func init() {
	var err error
	tmpl, err = fs.Sub(embedTmpl, "_embed/templates")
	if err != nil {
		log.Println(err)
	}
}

func Write(w http.ResponseWriter, data interface{}, names ...string) error {

	tmpl, err := template.ParseFS(tmpl, names...)
	if err != nil {
		return xerrors.Errorf("template.ParseFS() error: %w", err)
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		return xerrors.Errorf("Template Execute() error: %w", err)
	}

	return nil
}

func WriteJSON(w http.ResponseWriter, data interface{}) error {
	log.Println("Not Implemented")
	return nil
}
