package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/qt-luigi/go-auth-sample/firebase"
	"github.com/qt-luigi/go-auth-sample/google"
	"github.com/qt-luigi/go-auth-sample/index"
	"github.com/qt-luigi/go-auth-sample/openid"
)

func init() {
	// Set Index page handlers
	http.HandleFunc("/", index.IndexHandler)

	// Set Google Sign-in handlers
	http.HandleFunc("/loginGoogle", google.LoginHandler)
	http.HandleFunc("/callbackGoogle", google.CallbackHandler)
	http.HandleFunc("/redirectGoogle", google.RedirectHandler)

	// Set Firebase Auth handlers
	http.HandleFunc("/loginFirebase", firebase.LoginHandler)

	// Set OpenID Connect handlers
	http.HandleFunc("/loginOpenID", openid.LoginHandler)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
