
package hello

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func init() {
    // Register a couple of routes.
    r := mux.NewRouter()

    r.HandleFunc("/", homeHandler)
    r.HandleFunc("/{salutation}/{name}", helloHandler)

    // Send all incoming requests to mux.DefaultRouter.
    http.Handle("/", r)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", "text/html")
    fmt.Fprint(w, "Try a <a href='/Hello/world'>hello</a>.")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", "text/html")
    vars := mux.Vars(r)
    phrase := vars["salutation"] + ", " + vars["name"] + "!"
    fmt.Fprint(w, phrase)
}
