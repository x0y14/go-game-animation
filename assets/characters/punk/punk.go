package punk

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/x0y14/gameAnimation/assets/characters"
	"image"
	"log"
)

var PunkAssets characters.CharacterSprites

const (
	fHeight   = 32
	fWidth    = 32
	fInterval = 16
	fOriginX  = 0
	fOriginY  = 16
)

func init() {
	idleBuf := bytes.NewBuffer(IdleBytes)
	idleImg, _, err := image.Decode(idleBuf)
	if err != nil {
		log.Fatal(err)
	}
	idleImgEbi := ebiten.NewImageFromImage(idleImg)

	runBuf := bytes.NewBuffer(RunBytes)
	runImg, _, err := image.Decode(runBuf)
	if err != nil {
		log.Fatal(err)
	}
	rnuImgEbi := ebiten.NewImageFromImage(runImg)

	PunkAssets = characters.CharacterSprites{
		Idle: &characters.Sprite{
			Img:           idleImgEbi,
			FrameNum:      IdleFrameNum,
			FrameHeight:   fHeight,
			FrameWidth:    fWidth,
			FrameInterval: fInterval,
			FrameOriginX:  fOriginX,
			FrameOriginY:  fOriginY,
		},
		Run: &characters.Sprite{
			Img:           rnuImgEbi,
			FrameNum:      RunFrameNum,
			FrameHeight:   fHeight,
			FrameWidth:    fWidth,
			FrameInterval: fInterval,
			FrameOriginX:  fOriginX,
			FrameOriginY:  fOriginY,
		},
	}
}
