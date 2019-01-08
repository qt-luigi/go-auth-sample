package openid

import (
	"fmt"
	"net/http"
)

// LoginHandler for OAuth 2.0 and OpenID Connect login.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OAuth 2.0 and OpenID Connect")
}
