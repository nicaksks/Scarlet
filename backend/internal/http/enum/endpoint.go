package enum

type Endpoint string

const (
	PROFILE        Endpoint = "/users/{0}"
	HISTORIES      Endpoint = "/users/{0}/histories"
	PLAYCOUNT      Endpoint = "/users/playCount"
	CHARACTERSINFO Endpoint = "/characters/info/{0}"
)
