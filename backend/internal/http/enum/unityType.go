package enum

type UnitType int

const (
	CHARACTER UnitType = iota + 1
	MONSTER
	EVENT_OBJECT
)
