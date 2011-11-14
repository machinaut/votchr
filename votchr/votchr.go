package votchr

import (
    "appengine"
    "appengine/datastore"
    "appengine/user"
    "fmt"
    "http"
    //"os"
    //"template"
    //"time"
)

func init() {
    http.HandleFunc("/", root)
    http.HandleFunc("/image", image)
    http.HandleFunc("/votch", votch)
    http.HandleFunc("/login", login)
}

func votch(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    u := user.Current(c)
    if u != nil {
        _, err := user.LogoutURL(c, "/")
        if err != nil {
            http.Error(w, err.String(), http.StatusInternalServerError)
            return
        }
        fmt.Fprintf(w, `{"votch_url":"http://i.imgur.com/CW5y1.jpg"}`)
    } else {
        unauthorized(w, r)
    }
}

type View struct {
    Url string
    Viewer string
}

func image(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    v := View {
        Url : r.URL.String(),
        Viewer : r.RemoteAddr,
    }
    _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "image", nil), &v)
    if err != nil {
        http.Error(w, err.String(), http.StatusInternalServerError)
        return
    }
    fmt.Fprintf(w, "%v",v)
}

func root(w http.ResponseWriter, r *http.Request) {
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
        fmt.Fprintf(w, "Please, <a href='/login'>login</a>.")
    }
}

func login(w http.ResponseWriter, r *http.Request) {
    providers := map[string]string {
        "Google"   : "www.google.com/accounts/o8/id",
        "Yahoo"    : "yahoo.com",
        "MySpace"  : "myspace.com",
        "AOL"      : "aol.com",
        "MyOpenID" : "myopenid.com",
    }

    c := appengine.NewContext(r)
    fmt.Fprintf(w, "Hey you there sigin in at: ")
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
func unauthorized(w http.ResponseWriter, req *http.Request) {
    w.WriteHeader(http.StatusUnauthorized)
}
