package enum

type AuthProvider int

const (
	DISCORD AuthProvider = 20
	GUEST   AuthProvider = 25
	BYPASS  AuthProvider = 26
)

type BuildPlatform string

const NN BuildPlatform = "NIMBLE_NEURON"

type OsType int

const (
	ANDROID OsType = iota + 1
	IOS
	WINDOWS64
	WINDOWS_DMM
	ETC OsType = 99
)

type SessionKey int

const DEFAULT SessionKey = 582703210599415808

type Market int

const (
	APP_STORE Market = iota + 1
	GOOGLE
	STEAM  Market = 6
	DMM    Market = 8
	XSOLLA        = 10
)
