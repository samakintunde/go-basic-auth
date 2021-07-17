package main

import (
	"crypto/sha256"
	"crypto/subtle"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Server config
var (
	port = os.Getenv("PORT")
)

const (
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
	if port == "" {
		port = "5000"
	}

	app := new(App)
	app.auth.username = os.Getenv("AUTH_USERNAME")
	app.auth.password = os.Getenv("AUTH_PASSWORD")

	if app.auth.username == "" || app.auth.password == "" {
		log.Fatalln(`Please, set the "AUTH_USERNAME" and "AUTH_PASSWORD" environment variables`)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", handleHome)                              // unprotected route
	mux.HandleFunc("/dashboard", app.basicAuth(handleDashboard)) // Protected route

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      mux,
		IdleTimeout:  idleTimeout,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	log.Printf("Starting server on port %s", port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func handleHome(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, "The home route")
	if err != nil {
		return
	}
}

func handleDashboard(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, "The protected dashboard route")
	if err != nil {
		return
	}
}

// Basic Auth is a middleware that handles basic auth check on resource access
func (a *App) basicAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()

		if ok {
			usernameHash := sha256.Sum256([]byte(username))
			passwordHash := sha256.Sum256([]byte(password))

			expectedUsernameHash := sha256.Sum256([]byte(a.auth.username))
			expectedPasswordHash := sha256.Sum256([]byte(a.auth.password))

			// Checks the corresponding hashes for a match before returning 1 for equal or 0 for unequal
			// Prevents hackers guessing what exact character is wrong by monitoring the time it takes for
			// the request to fail
			doesUsernameMatch := subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1
			doesPasswordMatch := subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1

			if doesUsernameMatch && doesPasswordMatch {
				next.ServeHTTP(w, r)
				return
			}
		}

		// This tells the browser/client to use Basic Authentication.
		// Browsers will show a native dialog for username and password fields
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}
