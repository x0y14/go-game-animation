package physics

func ConvertFrameCountToSec(count int) float64 {
	return float64(count / 60.0)
}

func ConvertSecToFrameCount(sec float64) int {
	return int(sec * 60)
}

func ConvertPixelToMeter(pix int) float64 {
	return float64(pix / 16.0)
}
func ConvertMeterToPixel(m float64) int {
	return int(m * 16.0)
}
