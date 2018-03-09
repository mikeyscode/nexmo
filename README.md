# nexmo
<p align="center">
    <a href="https://godoc.org/github.com/mikeyscode/nexmo"><img src="https://godoc.org/github.com/mikeyscode/nexmo?status.svg" alt="GoDoc"></a>
</p>

Go wrapper for the Nexmo Library

#### Usage

```
auth := nexmo.Auth{Key: "<apikey>", Secret: "<apisecret>"}

err := nexmo.Setup(auth)
if err != nil {
	panic(err)
}
message, err := nexmo.Nexmo.SendSMS("4401234567890", "Mikey's Phone", nexmo.SMSOptions{
	Text: "Hello World, this text was sent from Nexmo",
})
if err != nil {
	panic(err)
}

fmt.Printf("%+v\n", message)
```
