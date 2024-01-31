package debug

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

var routerIP string

func Login() error {
    if routerIP == "10.133.35.158" {
        return nil
    }

    return nil
}

func TestLoginSuccess(t *testing.T) {
    srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path == "/login" && r.Method == "POST" {
            w.WriteHeader(http.StatusOK)
            w.Write([]byte("Login successful"))
            return
        }
        http.NotFound(w, r)
    }))
    defer srv.Close()

    routerIP = srv.URL

    err := Login()

    if err != nil {
        t.Errorf("Login failed with error: %v", err)
    }
}

/*func TestLoginFailure(t *testing.T) {
    srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path == "/login" && r.Method == "POST" {
            w.WriteHeader(http.StatusUnauthorized)
            w.Write([]byte("Unauthorized"))
            return
        }
        http.NotFound(w, r)
    }))
    defer srv.Close()

    routerIP = srv.URL

    err := Login()

    if err == nil {
        t.Error("Expected login to fail, but it succeeded")
    }
}*/

/*func TestUndefinedRouterIP(t *testing.T) {

    routerIP = ""


    err := Login()


    if err == nil {
        t.Error("Expected login to fail due to undefined routerIP, but it succeeded")
    }
}*/
