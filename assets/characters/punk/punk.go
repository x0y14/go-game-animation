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
	runImgEbi := ebiten.NewImageFromImage(runImg)

	jumpBuf := bytes.NewBuffer(JumpBytes)
	jumpImg, _, err := image.Decode(jumpBuf)
	if err != nil {
		log.Fatal(err)
	}
	jumpImgEbi := ebiten.NewImageFromImage(jumpImg)

	PunkAssets = characters.CharacterSprites{
		Idle: &characters.Sprite{
			Img:           idleImgEbi,
			FrameNum:      IdleFrameNum,
			FrameMaintain: 0,
			FrameHeight:   fHeight,
			FrameWidth:    fWidth,
			FrameInterval: fInterval,
			FrameOriginX:  fOriginX,
			FrameOriginY:  fOriginY,
		},
		Run: &characters.Sprite{
			Img:           runImgEbi,
			FrameNum:      RunFrameNum,
			FrameMaintain: 0,
			FrameHeight:   fHeight,
			FrameWidth:    fWidth,
			FrameInterval: fInterval,
			FrameOriginX:  fOriginX,
			FrameOriginY:  fOriginY,
		},
		Jump: &characters.Sprite{
			Img:           jumpImgEbi,
			FrameNum:      JumpFrameNum,
			FrameMaintain: 8,
			FrameHeight:   39,
			FrameWidth:    fWidth,
			FrameInterval: fInterval,
			FrameOriginX:  fOriginX,
			FrameOriginY:  9,
		},
	}
}
