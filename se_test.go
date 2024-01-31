package debug

import (
 //   "bytes"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
)

var routerIP = "10.133.35.158"

func main() {

    username := "admin"
    password := "tcs123"

    formData := url.Values{}
    formData.Set("username", username)
    formData.Set("password", password)

    resp, err := http.PostForm("http://"+routerIP+"/login", formData)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
        return
    }

    if resp.StatusCode == http.StatusOK {
        fmt.Println("Login successful!")

    } else {
        fmt.Printf("Login failed with status code: %d\n", resp.StatusCode)
        fmt.Println("Response body:", string(body))
    }
}

func Login() error {

    username := "admin"
    password := "tcs123"

    formData := url.Values{}
    formData.Set("username", username)
    formData.Set("password", password)

    resp, err := http.PostForm("http://"+routerIP+"/login", formData)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("login failed with status code: %d", resp.StatusCode)
    }

    return nil
}
