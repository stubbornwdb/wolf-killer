package app

type Game interface {
	Start() error
	End() error
	Vote(Player) error
	Night() error
	Day() error
}
