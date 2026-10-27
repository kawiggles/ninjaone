package main

import (
	"fmt"
	"net/http"
)

func main() {
	client := &http.Client{}
	token := GetCred(client)
}

func SetHeader(req *http.Request, token string) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Kawika's CLI Stuff")
}
