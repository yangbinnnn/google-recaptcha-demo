package core

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	sitekey   = "6LcLcXkUAAAAAJniZZQramraMQ6rUjby0MA5d4dC"
	secret    = "6LcLcXkUAAAAAB467eqJgKCmzTUbkxCzAq9mC77r"
	verifyAPI = "https://www.google.com/recaptcha/api/siteverify"
)

func Verify(recaptcha string) []byte {
	body := fmt.Sprintf("secret=%s&response=%s", secret, recaptcha)
	client := &http.Client{}
	fmt.Println(body)
	req, _ := http.NewRequest("POST", verifyAPI, strings.NewReader(body))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, _ := client.Do(req)
	res, _ := ioutil.ReadAll(resp.Body)
	return res
}
