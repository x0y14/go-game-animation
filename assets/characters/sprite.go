package characters

import "github.com/hajimehoshi/ebiten/v2"

type Sprite struct {
	//Bytes         []byte
	Img           *ebiten.Image
	FrameNum      int
	FrameOriginX  int
	FrameOriginY  int
	FrameInterval int
	FrameWidth    int
	FrameHeight   int
}
