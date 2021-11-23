package stage1

import (
	"bytes"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/x0y14/gameAnimation/assets/stages"
	"image"
	_ "image/png"
	"log"
)

//go:embed background.png
var stage1Bytes []byte

var Stage1 stages.Stage

func init() {

	stage1Buf := bytes.NewBuffer(stage1Bytes)
	stage1Img, _, err := image.Decode(stage1Buf)
	if err != nil {
		log.Fatal(err)
	}
	stage1ImgEbi := ebiten.NewImageFromImage(stage1Img)

	Stage1 = stages.Stage{Img: stage1ImgEbi}
}
