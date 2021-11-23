package main

import (
	"context"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/x0y14/gameAnimation/assets/characters"
	"github.com/x0y14/gameAnimation/assets/stages/stage1"
	"github.com/x0y14/gameAnimation/character"
	"image"
	_ "image/png"
	"log"
)

const screenWidth = 320
const screenHeight = 240

type Game struct {
	keys  []ebiten.Key
	count int
}

var mrPunk *character.Character

func (g *Game) Update() error {
	g.count++
	mrPunk.Count++

	g.keys = inpututil.AppendPressedKeys(g.keys[:0])

	if len(g.keys) == 0 && mrPunk.Situation != character.Idling && mrPunk.CountSituationMaintain < g.count {
		mrPunk.UpdateSituation(character.Idling)
	}
	for _, key := range g.keys {
		if key == ebiten.KeyArrowRight {
			if mrPunk.OffsetX < screenWidth {
				mrPunk.OffsetX += 2
			}
			if mrPunk.Direction != character.Right {
				mrPunk.Direction = character.Right
			}
			if mrPunk.Situation != character.Running {
				mrPunk.UpdateSituation(character.Running)
			}
		} else if key == ebiten.KeyArrowLeft {
			if 0 < mrPunk.OffsetX {
				mrPunk.OffsetX -= 2
			}
			if mrPunk.Direction != character.Left {
				mrPunk.Direction = character.Left
			}
			if mrPunk.Situation != character.Running {
				mrPunk.UpdateSituation(character.Running)
			}
		} else if key == ebiten.KeySpace && mrPunk.Situation != character.Jumping {
			mrPunk.CountSituationMaintain = g.count + mrPunk.Sprites.Jump.FrameMaintain*mrPunk.Sprites.Jump.FrameNum
			mrPunk.UpdateSituation(character.Jumping)
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(stage1.Stage1.Img, &ebiten.DrawImageOptions{})
	g.DrawCharacter(screen, mrPunk)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) DrawCharacter(screen *ebiten.Image, c *character.Character) {
	op := &ebiten.DrawImageOptions{}

	var sprite *characters.Sprite

	switch c.Situation {
	case character.Idling:
		sprite = c.Sprites.Idle
	case character.Running:
		sprite = c.Sprites.Run
	case character.Jumping:
		sprite = c.Sprites.Jump
	}

	op.GeoM.Translate(-float64(sprite.FrameWidth)/2, -float64(sprite.FrameHeight)/2)
	if mrPunk.Direction == character.Left {
		op.GeoM.Scale(-1, 1)
	}
	op.GeoM.Translate(float64(mrPunk.OffsetX), float64(mrPunk.OffsetY))
	i := (c.Count / 10) % sprite.FrameNum
	sx, sy := sprite.FrameOriginX+(i*sprite.FrameWidth+i*sprite.FrameInterval), sprite.FrameOriginY

	img := sprite.Img.SubImage(image.Rect(sx, sy, sx+sprite.FrameWidth, sy+sprite.FrameHeight)).(*ebiten.Image)

	screen.DrawImage(img, op)
}

func Gravity(c *character.Character) {
	for {
		break
	}
}

func main() {
	mrPunk = character.NewPunkTypeCharacter("mr", screenWidth/2, screenHeight-10-16)
	punkCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go mrPunk.ListenUpdateSituation(punkCtx)

	go Gravity(mrPunk)

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("character animation")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
