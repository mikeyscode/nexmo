# nexmo

<a href="https://godoc.org/github.com/mikeyscode/nexmo"><img src="https://godoc.org/github.com/mikeyscode/nexmo?status.svg" alt="GoDoc"></a>

An API wrapper for the Nexmo library crafted in Go. 

## Usage

**Sending an SMS Message**
```
auth := auth.New("<key>", "<secret>")
sms.Auth(auth)

message, err := sms.Send("<to>", "<from>", sms.Options{Text: "Hello Gophers"})
if err != nil {
	panic(err)
}

fmt.Printf("%+v\n", message)
```
