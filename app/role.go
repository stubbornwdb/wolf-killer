package app

type Role interface {
	Action(Player) error
}

type Werewolf interface {
	Role
	Kill(Player) error
}

type Villager interface {
	Role
	Vote(Player) error
}

type Seer interface {
	Role
	Check(Player) error
}

type Witch interface {
	Role
	Save(Player) error
	Poison(Player) error
}

type Hunter interface {
	Role
	Shoot(Player) error
}
