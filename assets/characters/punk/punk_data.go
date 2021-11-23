package punk

import (
	_ "embed"
	_ "image/png"
)

//go:embed Punk_idle.png
var IdleBytes []byte

const IdleFrameNum = 4

//go:embed Punk_run.png
var RunBytes []byte

const RunFrameNum = 6
