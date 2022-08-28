package logic

import (
	"io"
	"os/exec"
	"strings"

	"golang.org/x/xerrors"
)

const PlantUMLJar = "plantuml-nodot.1.2022.7.jar"

func WriteImage(w io.Writer, buf *strings.Reader) error {

	err := runPlantUML(w, buf, w)
	if err != nil {
		return xerrors.Errorf("runPlantUML() error: %w", err)
	}
	return nil
}

func runPlantUML(w io.Writer, r io.Reader, errW io.Writer) error {

	cmd := exec.Command("java", "-jar", PlantUMLJar, "-pipe")
	cmd.Stdin = r
	cmd.Stdout = w
	cmd.Stderr = errW

	err := cmd.Run()
	if err != nil {
		return xerrors.Errorf("cmd.Run() error: %w", err)
	}

	return nil
}
