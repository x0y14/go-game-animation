package character

import (
	"context"
	"testing"
	"time"
)

func TestCharacter_UpdateSituation(t *testing.T) {
	cha := NewPunkTypeCharacter("tom")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go cha.ListenUpdateSituation(ctx)

	time.Sleep(time.Second)
	cha.UpdateSituation(Idling)
	time.Sleep(time.Second)
	cha.UpdateSituation(Running)
}
