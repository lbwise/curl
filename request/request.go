package request

import (
	"os"
	"strings"
)

type Request struct {
	URL string
	Cookies string
	UserAgent string
	Help bool
	Headers bool
	Time bool
}


func ParseCmd() *Request {
	args := os.Args[1:]
	var req *Request = &Request{}
	var keys []string
	for _, arg := range args {
		keys = strings.Split(arg, "=")
		switch keys[0] {
			case "cookies":
				req.Cookies = keys[1]
			
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
	return req
}