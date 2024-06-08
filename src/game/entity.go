package game

// Entity is the "base" of mobs and the character
type Entity interface {
	Turn(Entity, int)
	TakeDamage(Entity, int)
	Attack(Entity, int)
	GetInitiative() int
	GetName() string
}
