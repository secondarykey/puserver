package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	. "umlserver/handler/internal"
	"umlserver/logic"
)

func Register() error {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/convert/image", convertImageHandler)
	return nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := Write(w, nil, "index.tmpl")
	if err != nil {
		log.Println(err)
	}
}

type ImageRequest struct {
	Text string `json:"text"`
}

func convertImageHandler(w http.ResponseWriter, r *http.Request) {

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var req ImageRequest

	err = json.Unmarshal(b, &req)
	if err != nil {
		log.Println(err)
		return
	}

	buf := strings.NewReader(req.Text)

	err = logic.WriteImage(w, buf)
	if err != nil {
		log.Println(err)
		return
	}

}
