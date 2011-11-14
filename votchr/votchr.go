package votchr

import (
    "appengine"
    //"appengine/datastore"
    "appengine/user"
    "fmt"
    "http"
    //"os"
    //"template"
    //"time"
)

func init() {
    http.HandleFunc("/", hello)
    http.HandleFunc("/votch", votch)
    http.HandleFunc("/_ah/login_required", openIdHandler)
}

func votch(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    u := user.Current(c)
    if u != nil {
        url, err := user.LogoutURL(c, "/")
        if err != nil {
            http.Error(w, err.String(), http.StatusInternalServerError)
            return
        }
        fmt.Fprintf(w, "Hello, %s! (<a href='%s'>Sign out</a>)", u, url)
    } else {
        unauthorized(w, r)
    }

}

func hello(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    u := user.Current(c)
    if u != nil {
        url, err := user.LogoutURL(c, "/")
        if err != nil {
            http.Error(w, err.String(), http.StatusInternalServerError)
            return
        }
        fmt.Fprintf(w, "Hello, %s! (<a href='%s'>Sign out</a>)", u, url)
    } else {
        fmt.Fprintf(w, "Please, <a href='/_ah/login_required'>login</a>.")
    }

}

func openIdHandler(w http.ResponseWriter, r *http.Request) {
    providers := map[string]string {
        "google"   : "gmail.com",
    }

    c := appengine.NewContext(r)
    fmt.Fprintf(w, "Sign in at: ")
    for name, url := range providers {
        login_url, err := user.LoginURLFederated(c, "/", url)
        if err != nil {
            http.Error(w, err.String(), http.StatusInternalServerError)
            return
        }
        fmt.Fprintf(w, "[<a href='%s'>%s</a>]", login_url, name)
    }
}

// return 401 unauthorized
func unauthorized(w http.ResponseWriter, req *http.Request){
    w.Header().Set("Content-Type", "text/plain;" + "charset=utf-8")
    w.WriteHeader(http.StatusUnauthorized)
}
