package nexmo

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	makeCallEndpoint = "https://api.nexmo.com/v1/calls"
	makeCallMethod   = "POST"
	callPayload      = `{"to":[{"type": "phone","number": 447985101848}],
      "from": {"type": "phone","number": 12345678901},
	  "answer_url":["https://nexmo-community.github.io/ncco-examples/first_call_talk.json"]}`

	applicationIDHardcoded = ""

	privateKeyHardcoded = ``
)

// TextCallOptions for calls
type TextCallOptions struct {
}

// TextCallResponseInterface of what a response from generating a phone call would give you
type TextCallResponseInterface interface {
}

// DispatchTextCall will dispatch call to a specified number
func (n *Nexmo) DispatchTextCall(to, from string, options TextCallOptions) (TextCallResponseInterface, error) {
	rand.Seed(time.Now().UnixNano())
	token := jwt.New(jwt.SigningMethodRS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["iat"] = time.Now().UTC().Unix()
	claims["application_id"] = applicationIDHardcoded
	claims["jti"] = rand.Intn(100)

	var privateKey *rsa.PrivateKey

	block, data := pem.Decode([]byte(privateKeyHardcoded))

	dataString := string(data)
	fmt.Println(dataString)
	fmt.Println(block)

	privateKey, err := x509.ParsePKCS1PrivateKey(data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %v", err)
	}
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token string: %v", err)
	}

	fmt.Println(tokenString)

	client := &http.Client{Timeout: time.Second * 15}

	req, err := http.NewRequest("POST", makeCallEndpoint, bytes.NewBuffer([]byte(callPayload)))
	if err != nil {
		return nil, fmt.Errorf("unable to create post request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Bearer", tokenString)

	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("unable to process post request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read response body: %v", err)
	}

	fmt.Println(string(respBody))

	return string(respBody), fmt.Errorf(tokenString)
}
