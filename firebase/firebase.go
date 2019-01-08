package firebase

import (
	"fmt"
	"net/http"
)

// LoginHandler for Firebase Auth login.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Firebase Auth Test")
}
