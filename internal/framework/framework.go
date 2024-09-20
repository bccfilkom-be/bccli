package framework

type Framework int

const (
	ECHO Framework = iota
	FIBER
	GIN
	MUX
	NET
)

type Generator interface {
	Main() error
	Interface() error
}

func NewGenerator(framework Framework) Generator {
	switch framework {
	//TODO: default to net/http
	default:
		return &echo{}
	}
}
