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

	applicationIDHardcoded = "fc26cd5d-3171-4b42-9487-8c793b12dd0c"

	privateKeyHardcoded = `-----BEGIN RSA PRIVATE KEY-----
	MIIEowIBAAKCAQEAvLj9y5ySvGYXZvwH0yZlB48rFuadPPWOhrfbu/5j5/xTVyAd
	VjR73lIN92a3Sbc4+XaGlfw6NdzQAPu3RzEp+Yab7CS84VjJ48+IWEUqX6z/gi39
	U/Y/pFrpXXXfhIA1P1ApUFC59uWO2Avz7IaUqXIMWCEABFi6ZqUERR2jqfPMFQrU
	c8WPpg624Wpn3zRXFoGDWqtGE6UOeqAOG8Dh67clloQBSpFa9Izz/hQB8uTAn5Yu
	RPTNINjv0kLGSPBMBvvg+4g/v+aiSGhyeOeBpgvUsuATAi/wAM0zECJDThHytxMK
	jls05K3I4BZGMzFM2wqKP8VHFhaMZavrSzIzRQIDAQABAoIBAC4HMF8ooOEyRSLo
	9T+abamaUXgUZuUnPsu8q+r98H/0Gp91RbJwuoVOnflpI+rmtQ6iydBq5AefA1w6
	CElkxEgHfJ/rleWgMh1N0IM2207acrbdYJvJw1vikGgrB4jZfCMk+e6Mwc5lzqEC
	yUs2x6tMFZao9cgZm9zNAm2Y/QHF+IQpi60/voncs+rEXk4Ftr4WKskUMWr30CV4
	J+IdBxkYRDAOYDFtqjWTjxAcjJRypaORUjgPsqnDvUQk7Z5fwipKXawZgP/ubsBC
	rWXmXMRL3hUHvL3CMpQDQGJ2UwI99WHTQ0/eV01ZqqfPu+QQYnc2Ec0cs+fuAODC
	k4O+z0ECgYEA8+HuV37RsR5dUpYrqTcFGsa7gatHkqszNFSbXYzBEyOVmFyqtoy+
	RYM8zCuzyTGV4IoH1vCKCDpIXYZXv58Zy5r8jimP6z26Tu2cUb2/o8LGvvyTl8VJ
	kNUNdKr+dRul7thVH6cd13g8MLv18fdQhXn92/E+JkluxTJtnNPUNxMCgYEAxhlz
	38NP7sbO5ZVtPPP41TxY9KDMmQbrZVbUjk/hraBr6NUZWh65RZEf57LtNzOcXANA
	Fx2MKr00dn6fJVx9knSCZjEus/Dii37oGOfwupwmCiqLv5H+dFc4HR5OQaGfJO48
	K24BB6WSx4Zcvwgshq5NXo9id5ia0wH8ou35/0cCgYEAxd9YlurRTah6RUiMMiXu
	4VO+zK1gS9LVn67Jw7Qw7stfU5hT9frpYdLiIDGRFDtEBENZqv3MsHJBRoh6Z4G5
	1yVvphR4rX+Oyv0kaHnQpBijUk/xnCE41+bUnQUjoXaGQeyJ3D2mC62FAHFwUhq7
	3SAmZS4to7jOw/ZvUt/XfbsCgYAM4sT1zjeZ2ZbFulWTvG74N+e1aexFG/0d52sG
	Is4URDYgvBPdF9iHXOxNXwctKw9FsPRvTH28nfgWqR/jB0QnzapZyWM3Uzj5R/UD
	AbtX+CZFTQUwHegGW1IMGteOT1wRw+loDczFWZDVp7jKuFZlIFtqBjuqeePATAXJ
	917mqwKBgFDncKzJ1Liu/jlhUrQ+5cRUHxYduzwhh3xvPucKhLogeAgRfzK45EBu
	YHoRQLy6N53Lu/eWYBUsE5P5L2kmOBUNMkOBKPoyPRZjnPQmGss4IbEfM7gLaacC
	MdCMX5r1hubdt3mhEUaRXsQHFuqsI5fB+6EHgMfLZT7Cf1jAt6pm
	-----END RSA PRIVATE KEY-----`
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

	t := time.Now().UTC()
	d := 24 * time.Hour
	t.Truncate(d)

	claims["iat"] = t.Unix()
	claims["application_id"] = applicationIDHardcoded
	claims["jti"] = rand.Intn(100)

	var privateKey *rsa.PrivateKey
	b, err := ioutil.ReadFile("path to private key")
	if err != nil {
		return nil, err
	}

	// block, data := pem.Decode([]byte(privateKeyHardcoded))

	// dataString := string(data)
	// fmt.Println(dataString)
	// fmt.Println(block)

	block, _ := pem.Decode(b)
	// der, err := x509.DecryptPEMBlock(block, []byte(""))
	// if err != nil {
	// 	return nil, err
	// }
	if len(block.Bytes) == 0 {
		return nil, fmt.Errorf("could not decode private key")
	}
	privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %v", err)
	}

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token string: %v", err)
	}

	client := &http.Client{Timeout: time.Second * 15}

	req, err := http.NewRequest("POST", makeCallEndpoint, bytes.NewBuffer([]byte(callPayload)))
	if err != nil {
		return nil, fmt.Errorf("unable to create post request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+tokenString)

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
