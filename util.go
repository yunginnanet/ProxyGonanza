package proxygonanza

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

func processBody(res *http.Response) ([]byte, error) {
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func GetMyIP() net.IP {
	res, err := http.DefaultClient.Get("https://wtfismyip.com/text")
	if err != nil {
		fmt.Println(err)
		return net.IP{}
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return net.IP{}
	}
	return net.ParseIP(strings.TrimSpace(string(body)))
}
