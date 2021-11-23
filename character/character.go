package character

import (
	"context"
	"fmt"
	"github.com/x0y14/gameAnimation/assets/characters"
	"github.com/x0y14/gameAnimation/assets/characters/punk"
)

type Character struct {
	Name string
	Situation
	OffsetX int
	OffsetY int
	Direction
	Sprites *characters.CharacterSprites
	//ReDrawRequestSender     chan<- bool // for drawer
	//ReDrawRequestReceiver   <-chan bool // for drawer
	SituationUpdateSender   chan<- Situation
	SituationUpdateReceiver <-chan Situation
	GiveUpSender            <-chan bool
	Count                   int
}

func NewPunkTypeCharacter(name string, x, y int) *Character {
	//reDrawCh := make(chan bool)
	situationCh := make(chan Situation)

	return &Character{
		Name:      name,
		Situation: Idling,
		OffsetX:   x,
		OffsetY:   y,
		Direction: Right,
		Sprites:   &punk.PunkAssets,
		//ReDrawRequestSender:     reDrawCh,
		//ReDrawRequestReceiver:   reDrawCh,
		SituationUpdateSender:   situationCh,
		SituationUpdateReceiver: situationCh,
		Count:                   0,
	}
}
func (c *Character) ListenUpdateSituation(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("状態の監視をやめます")
			return
		case <-c.SituationUpdateReceiver:
			c.Situation = <-c.SituationUpdateReceiver
			fmt.Printf("状態が変更されました: %v\n", c.Situation.String())
			//c.ReDrawRequestSender <- true
			c.Count = 0
		}
	}
}

func (c *Character) UpdateSituation(situation Situation) {
	fmt.Printf("状態変更要請: %v\n\n", situation.String())
	c.SituationUpdateSender <- situation
}
