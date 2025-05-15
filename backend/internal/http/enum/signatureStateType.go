package enum

type SignatureStateType int

const (
	UNAVAILABLE SignatureStateType = iota + 1
	AVAILABLE
	OPEN SignatureStateType = iota + 101
	COMPLETE
)
