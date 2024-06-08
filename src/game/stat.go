package game

type Stat struct {
	Name string
	Value int
}

func NewStat(name string, value int)(*Stat) {
	stat := new(Stat)

	stat.Name = name
	stat.Value = value

	return stat
}
