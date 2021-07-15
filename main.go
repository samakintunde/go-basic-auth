package main

import (
	"log"
	"net/http"
	"time"
)

// Auth credentials
const (
	authUsername = "Samakintunde"
	authPassword = "WelcomeBoss"
)

// Server config
const (
	port         = ":5000"
	idleTimeout  = time.Minute
	readTimeout  = time.Second * 10
	writeTimeout = time.Second * 10
)

type App struct {
	auth struct {
		username string
		password string
	}
}

func main() {
	app := new(App)
	app.auth.username = authUsername
	app.auth.password = authPassword

	mux := http.NewServeMux()

	mux.HandleFunc("/", handleHome)
	mux.HandleFunc("/register", handleSignup)
	mux.HandleFunc("/login", handleSignin)
	mux.HandleFunc("/dashboard", app.basicAuth(handleDashboard))

	srv := &http.Server{
		Addr:         port,
		Handler:      mux,
		IdleTimeout:  idleTimeout,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	log.Printf("Starting server on port %s", port)
	if err := srv.ListenAndServeTLS("./localhost.pem", "./localhost-key.pem"); err != nil {
		log.Fatal(err)
	}
}

func handleSignup(w http.ResponseWriter, r *http.Request) {
	_ = w
	_ = r
}

func handleSignin(w http.ResponseWriter, r *http.Request) {
	_ = w
	_ = r
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	_ = w
	_ = r
}

func handleDashboard(w http.ResponseWriter, r *http.Request) {
	_ = w
	_ = r
}

// Basic Auth is a middleware that handles basic auth check on resource access
func (a *App) basicAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = w
		_ = r
		_ = next
		// Basic Auth checker
	}
}
