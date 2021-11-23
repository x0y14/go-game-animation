package character

type Direction int

const (
	_           Direction = iota
	Up                    // 上
	UpperRight            // 右上
	UpperLeft             // 左上
	Right                 // 右
	Left                  // 左
	Bottom                // 下
	BottomRight           // 右下
	BottomLeft            // 左下
)
