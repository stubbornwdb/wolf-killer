package app

type Player interface {
	GetRole() string
	Vote(Player) error
	IsAlive() bool
	Kill() error
	Save() error
	Poison() error
	Check() error
}
