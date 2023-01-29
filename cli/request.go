package cli

import (
	"os"
	"strings"
	"errors"
)

type Request struct {
	URL string
	Method string
	Cookies string
	UserAgent string
	Help bool
	Headers bool
	Time bool
}


func ParseCmd() (*Request, error) {
	if len(os.Args) == 1 {
		return nil, errors.New("please enter a URL to request")
	}
	args := os.Args[1:]
	var req *Request = &Request{}
	var keys []string
	for _, arg := range args {
		keys = strings.Split(arg, "=")
		switch keys[0] {
			case "cookies":
				req.Cookies = keys[1]
			
			case "method":
				req.Method = keys[1]

			case "agent":
				req.UserAgent = keys[1]
			
			case "-h":
				req.Help = true
			
			case "-H":
				req.Headers = true
			
			case "-T":
				req.Time = true

			default:
				req.URL = keys[0]
		}
	}
	return req, nil
}