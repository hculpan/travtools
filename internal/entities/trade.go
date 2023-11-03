package entities

const (
	PASSENGER int = iota
	FREIGHT
)

type Trade interface {
	Description() string
}
