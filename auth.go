package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Credential struct {
	Token		string `json:"access_token"`
	TTL			int64 `json:"expires_in"`
	TokenType	string `json:"token_type"`
	Refresh		string `json:"refresh_token"`
	Scope		string `json:"scope"`
}

func GetCred(client *http.Client) string {
	vals := url.Values{}
	vals.Set("grant_type", "client_credentials")
	vals.Set("client_id", os.Getenv("NINJAONE_ID"))
	vals.Set("client_secret", os.Getenv("NINJAONE_SECRET"))
	vals.Set("scope", "monitoring management control")

	req, err := http.NewRequest(
		http.MethodPost, 
		os.Getenv("NINJAONE_URL")+"/ws/oauth/token",
		strings.NewReader(vals.Encode()),
	)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("authentication response error: %v\n", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var cred Credential
	err = json.Unmarshal(bodyBytes, &cred)
	if err != nil {
		log.Fatalf("failed to unmarshal authentication response: %v\n", err)
	}

	return cred.Token
}
