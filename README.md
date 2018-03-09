# nexmo
<p align="center">
    <a href="https://godoc.org/github.com/mikeyscode/nexmo"><img src="https://godoc.org/github.com/mikeyscode/nexmo?status.svg" alt="GoDoc"></a>
</p>

Go wrapper for the Nexmo Library

#### Usage

```
client := &http.Client{Timeout: time.Second * 15}
	auth := nexmo.Auth{Key: "<apikey>", Secret: "<apisecret>"}

	n := nexmo.Setup(client, auth)
	message, err := n.SendSMS("4401234567890", "Mikey's Phone", nexmo.SMSOptions{
		Text: "Hello World, this text was sent from Nexmo",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", message)
```
