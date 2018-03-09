package nexmo

import "fmt"

// TextCallOptions for calls
type TextCallOptions struct {
}

// TextCallResponseInterface of what a response from generating a phone call would give you
type TextCallResponseInterface interface {
}

// DispatchTextCall will dispatch call to a specified number
func (n *Nexmo) DispatchTextCall(to, from string, options TextCallOptions) (TextCallResponseInterface, error) {
	return nil, fmt.Errorf("error: not implemented")
}
