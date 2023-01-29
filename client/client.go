package client

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/lbwise/curl/cli"
)

func SendRequest(req *cli.Request) (string, error) {
	body := http.NoBody
	httpReq, _ := http.NewRequest(req.Method, req.URL, body)
	cl := &http.Client{}
	resp, err := cl.Do(httpReq)
	if err != nil {
		return "", errors.New("something went wrong")
	}
	respBytes, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return string(respBytes), nil
}