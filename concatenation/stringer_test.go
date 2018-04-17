package concatenation

import (
	"net/url"
	"testing"
)

func createRequest() Request {
	return Request{
		"ID",
		Client{
			"Name",
			url.URL{
				Scheme:   "https",
				User:     url.UserPassword("me", "pass"),
				Host:     "example.com",
				Path:     "foo/bar",
				RawQuery: "x=1&y=2",
				Fragment: "anchor",
			}}}
}

func BenchmarkConcatenate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Concatenate(createRequest())
	}
}

func BenchmarkFprintfate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fprintfate(createRequest())
	}
}

func BenchmarkSprintfate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Sprintfate(createRequest())
	}
}

func BenchmarkStrvconv(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Strvconv(createRequest())
	}
}
