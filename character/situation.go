package character

type Situation int

const (
	_ Situation = iota
	Idling
	Running
	Jumping
)

var situation = [...]string{
	Idling:  "Idling",
	Running: "Running",
	Jumping: "OnJumping",
}

func (s Situation) String() string {
	return situation[s]
}
