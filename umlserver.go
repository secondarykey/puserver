package umlserver

import (
	"fmt"
	"net/http"
	"umlserver/config"
	"umlserver/handler"

	"golang.org/x/xerrors"
)

func Listen(opts ...config.Option) error {

	err := config.Set(opts...)
	if err != nil {
		return xerrors.Errorf("config.Set() error: %w", err)
	}

	err = handler.Register()
	if err != nil {
		return xerrors.Errorf("handler.Register() error: %w", err)
	}

	conf := config.Get()
	srv := fmt.Sprintf("%s:%d", conf.Server, conf.Port)

	fmt.Printf("Listen -> [%s]\n", srv)

	err = http.ListenAndServe(srv, nil)
	if err != nil {
		return xerrors.Errorf("http.ListenAndServe() error: %w", err)
	}

	return nil
}
