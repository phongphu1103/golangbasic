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

type application struct {
    auth struct {
        username string
        password string
    }
}

func (app *application) protectedHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "This is the protected handler")
}

func (app *application) unprotectedHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "This is the unprotected handler")
}

func (app *application) basicAuth(next http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        username, password, ok := r.BasicAuth()
        if ok {
            usernameHash := sha256.Sum256([]byte(username))
            passwordHash := sha256.Sum256([]byte(password))
            expectedUsernameHash := sha256.Sum256([]byte(app.auth.username))
            expectedPasswordHash := sha256.Sum256([]byte(app.auth.password))
            usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
            passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)

            if usernameMatch && passwordMatch {
                next.ServeHTTP(w, r)
                return
            }
        }

        w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
    })
}

func main() {
    app := new(application)
    
    app.auth.username = os.Getenv("USERNAME")
    if password := os.Getenv("PASSWORD"); len(password) > 0 {
        app.auth.password = password
    } else {
        app.auth.password = "$jSZ9y})haSLp>LB"
    }
    
    fmt.Println(app.auth)
    // initializing server
    mux := http.NewServeMux()
    mux.HandleFunc("/unprotected", app.unprotectedHandler)
    mux.HandleFunc("/protected", app.basicAuth(app.protectedHandler))

    srv := &http.Server{
        Addr:         ":8000",
        Handler:      mux,
        IdleTimeout:  time.Minute,
        ReadTimeout:  10 * time.Second,
        WriteTimeout: 30 * time.Second,
    }

    log.Printf("starting server on %s", srv.Addr)
    err := srv.ListenAndServe()
    log.Fatal(err)
}