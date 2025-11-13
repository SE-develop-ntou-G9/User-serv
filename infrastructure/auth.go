package infrastructure

import (
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("GoogleClientSecret")))

func InitAuth() {
	GoogleClientID := os.Getenv("GoogleClientID")
	GoogleClientSecret := os.Getenv("GoogleClientSecret")

	callBackURL := os.Getenv("oAuthCallBackURL")

	goth.UseProviders(
		google.New(GoogleClientID, GoogleClientSecret, callBackURL, "email", "profile"),
	)

	gothic.Store = store
}
