package logic

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io"
	"log"
	"os/exec"
	"strings"
	"time"
	"umlserver/config"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/xerrors"
)

func WriteImage(w io.Writer, buf *strings.Reader) error {

	err := runPlantUML(w, buf, w)
	if err != nil {
		return xerrors.Errorf("runPlantUML() error: %w", err)
	}
	return nil
}

func runPlantUML(w io.Writer, r io.Reader, errW io.Writer) error {

	conf := config.Get()
	jar := conf.Jar

	cmd := exec.Command("java", "-jar", jar, "-pipe")
	cmd.Stdin = r
	cmd.Stdout = w
	cmd.Stderr = errW

	fmt.Println("Generate Start:", time.Now())
	err := cmd.Run()
	if err != nil {
		return xerrors.Errorf("cmd.Run() error: %w", err)
	}
	fmt.Println("Generate End  :", time.Now())
	return nil
}

func WriteErrorImage(w io.Writer, e error) {

	var err error

	buf := fmt.Sprintf("%+v", e)
	lines := strings.Split(buf, "\n")

	width := 500
	height := 300
	const SIZE = 12
	const SPACING = 1.2
	fg, bg := image.White, image.Black

	rgba := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)

	ft, err := truetype.Parse(goregular.TTF)

	c := freetype.NewContext()
	c.SetFont(ft)
	c.SetFontSize(SIZE)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)

	pt := freetype.Pt(10, 10+int(c.PointToFixed(SIZE)>>6))
	for _, s := range lines {
		_, err = c.DrawString(s, pt)
		if err != nil {
			log.Println(err)
			return
		}
		pt.Y += c.PointToFixed(SIZE * SPACING)
	}

	err = png.Encode(w, rgba)
	if err != nil {
		log.Println(err)
	}
}
