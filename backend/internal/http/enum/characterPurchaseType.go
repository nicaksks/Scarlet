package enum

type CharacterPurchaseType int

const (
	TRIAL CharacterPurchaseType = iota + 1
	FREE
	PURCHASED
)
