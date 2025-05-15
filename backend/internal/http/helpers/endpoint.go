package helpers

import (
	"bytes"
	"fmt"
	"scarlet/internal/http/enum"
	"strings"
)

func FormatEndpoint(e enum.Endpoint, values ...string) string {
	endpoint := string(e)

	for i, v := range values {
		variable := fmt.Sprintf("{%d}", i)
		endpoint = strings.Replace(endpoint, variable, v, -1)
	}

	return endpoint
}

func DefaultBody() *bytes.Reader {
	return bytes.NewReader([]byte("{}"))
}
