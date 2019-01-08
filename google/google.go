package google

import (
	"fmt"
	"net/http"
	"os"

	"github.com/rs/xid"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	api "google.golang.org/api/oauth2/v2"
)

var (
	googleCallbackURL = "http://%s/callbackGoogle"

	googleConfig      = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"email"},
		Endpoint:     google.Endpoint,
	}
)

// LoginHandler for Google Sign-in login.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Set RedirectURL here because a host changes local or server.
	googleConfig.RedirectURL = fmt.Sprintf(googleCallbackURL, r.Host)

	// Create and set "state" cookie.
	uuid := xid.New().String()
	cookie := &http.Cookie{
		Name:  "googleState",
		Value: uuid,
	}
	http.SetCookie(w, cookie)

	authURL := googleConfig.AuthCodeURL(uuid)

	http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)
}

// CallbackHandler for Google Sign-in callback.
func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	// Validate "state" for CSRF.
	googleState, err := r.Cookie("googleState")
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if googleState.Value != r.FormValue("state") {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// Create token by "code".
	code := r.FormValue("code")
	if code == "" {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	token, err := googleConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Create web client for Google.
	client := googleConfig.Client(oauth2.NoContext, token)

	//  Create OAuth2 service for Google.
	service, err := api.New(client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Get token info.
	tokenInfo, err := service.Tokeninfo().Do()
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Create and set "tokenEmail" cookie.
	cookie := &http.Cookie{
		Name:  "tokenEmail",
		Value: tokenInfo.Email,
	}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/redirectGoogle", http.StatusFound)
}

// RedirectHandler for Google Sign-in callback to redirect.
func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	tokenEmail, err := r.Cookie("tokenEmail")
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	fmt.Fprint(w, "Success: " + tokenEmail.Value)
}
