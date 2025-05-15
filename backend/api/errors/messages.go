package errors

type Message string

const (
	REGION_MISSING     Message = "missing.region"
	UNSUPPORTED_REGION Message = "unsupported.region"
	TOO_MANY_REQUESTS  Message = "too.many.requests"
)
