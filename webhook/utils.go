package webhook

import (
	"net/http"
)

// CheckBasicAuth validates the Basic Auth credentials in an HTTP request.
func CheckBasicAuth(r *http.Request, expectedUsername, expectedPassword string) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		return false
	}
	return username == expectedUsername && password == expectedPassword
}
