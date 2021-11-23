package character

type Situation int

const (
	_ Situation = iota
	Idling
	Running
)

var situation = [...]string{
	Idling:  "Idling",
	Running: "Running",
}

func (s Situation) String() string {
	return situation[s]
}
