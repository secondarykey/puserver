package umlserver

import (
	"net/http"
	"umlserver/handler"

	"golang.org/x/xerrors"
)

func Listen() error {

	err := handler.Register()
	if err != nil {
		return xerrors.Errorf("handler.Register() error: %w", err)
	}

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		return xerrors.Errorf("http.ListenAndServe() error: %w", err)
	}

	return nil
}
