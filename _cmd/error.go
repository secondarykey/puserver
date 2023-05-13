package main

import (
	"os"
	"umlserver/logic"

	"golang.org/x/xerrors"
)

func main() {

	fp, err := os.Create("error_image.png")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	logic.WriteErrorImage(fp, xerrors.Errorf("test"))

}
