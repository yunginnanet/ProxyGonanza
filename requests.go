package proxygonanza

import (
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"

	"git.tcp.direct/kayos/proxygonanza/internal"
)

func (api *APIClient) newRequest(method, u string) (r *http.Request) {
	r, _ = http.NewRequest(method, u, nil)
	if !strings.Contains(u, ".csv") {
		r.Header.Add("accept", "application/json")
	}
	r.Header.Add("Authorization", api.Key)

	api.debugPrintf("[%s] %s (Headers: %v)", method, u, r.Header)

	return
}

func (api *APIClient) getReq(endpoint string) ([]byte, error) {
	res, err := api.c.Do(api.newRequest("GET", APIBaseURL+endpoint))
	body, err := processBody(res)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (api *APIClient) postReq(endpoint string, post map[string]string) ([]byte, error) {
	params := url.Values{}
	for k, v := range post {
		params.Set(k, v)
	}
	enc := params.Encode()
	req, err := http.NewRequest("POST", APIBaseURL+endpoint, strings.NewReader(enc))
	if err != nil {
		return nil, err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", api.Key)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := api.c.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := processBody(res)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (api *APIClient) deleteReq(endpoint string) ([]byte, error) {
	req, err := http.NewRequest("DELETE", APIBaseURL+endpoint, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", api.Key)

	res, err := api.c.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := processBody(res)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func processBody(res *http.Response) ([]byte, error) {
	defer internal.CloseBody(res)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func getMyIP() (final net.IP, err error) {
	endpoint := "https://wtfismyip.com/text"
	envopt := os.Getenv("PROXYBONANZA_GETIP")
	if envopt != "" {
		endpoint = envopt
	}
	var res *http.Response
	res, err = http.DefaultClient.Get(endpoint)
	defer internal.CloseBody(res)
	if err != nil {
		return
	}
	var body []byte
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	final = net.ParseIP(strings.TrimSpace(string(body)))
	return
}
