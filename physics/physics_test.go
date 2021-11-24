package physics

import (
	"fmt"
	"testing"
)

func TestConvertFrameCountToSec(t *testing.T) {
	s := ConvertFrameCountToSec(120)
	fmt.Println(s)
}

func TestCalcGravity(t *testing.T) {
	p := 0
	for i := 0; i < 120; i++ {
		y := 0.5 * 9.8 * (ConvertFrameCountToSec(i) + 1) * (ConvertFrameCountToSec(i) + 1)

		p += ConvertMeterToPixel(y) / 60
		fmt.Printf("[frameCount %03d] %v\n", i, p)
	}

}
