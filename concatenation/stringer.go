package concatenation

import (
	"bytes"
	"fmt"
	"net/url"
	"time"
)

type Client struct {
	Name string
	Addr url.URL
}

type Request struct {
	ID string
	Client
}

func Concatenate(request Request) string {
	s := request.ID
	s += " " + request.Client.Addr.String()
	s += " " + time.Now().String()
	return s
}

func Fprintfate(request Request) string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "%s %v %v", request.ID, request.Client.Addr, time.Now())
	return b.String()
}

func Sprintfate(request Request) string {
	return fmt.Sprintf("%s %v %v", request.ID, request.Client.Addr, time.Now())
}

func Strvconv(request Request) string {
	b := make([]byte, 0, 40)
	b = append(b, request.ID...)
	b = append(b, ' ')
	b = append(b, request.Client.Addr.String()...)
	return string(b)
}
