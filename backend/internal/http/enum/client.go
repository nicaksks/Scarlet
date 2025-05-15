package enum

type Methods string

const (
	GET    Methods = "GET"
	POST   Methods = "POST"
	PUT    Methods = "PUT"
	DELETE Methods = "DELETE"
)

const (
	PROTOCOL = "http"
	DOMAIN   = "projectlumia.com"
	PATH     = "api"
)

type Region string

const (
	NA   Region = "na-master"       // Central
	EU   Region = "playtest.master" // Finland
	ASIA Region = "as-master"       // Singapore
)

type Port int

const (
	REST      Port = 10800 // Public
	USER      Port = 18000 // ???
	ADMIN     Port = 60696 // Admin
	WebSocket Port = 27000 // Websocket
)
