package proxygonanza

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	APIBaseURL = "https://proxybonanza.com/api/v1/"
	packages   = "userpackages.json"
	authips    = "authips.json"
)

// Package contains what we know about a particular proxybonanza package.
type Package struct {
	ID           int
	AuthIPs      []AuthIP
	AllTimeStats PackageStatistics
	HourlyStats  map[time.Time]PackageStatistics
}

// APIClient is a client for ProxyBonanza.com.
type APIClient struct {
	Key           string
	KnownPackages map[int]PackageDetails

	c *http.Client
}

// NewApiClient instantiates a proxybonanza.com API client with the given key.
func NewApiClient(key string) *APIClient {
	return &APIClient{
		Key:           key,
		KnownPackages: make(map[int]PackageDetails),
		c:             http.DefaultClient,
	}
}

func (api *APIClient) newRequest(method, u string) (r *http.Request) {
	r, _ = http.NewRequest(method, u, nil)
	r.Header.Add("accept", "application/json")
	r.Header.Add("Authorization", api.Key)

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
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", api.Key)
	if err != nil {
		return nil, err
	}

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

// GetProxyPackages gets current proxy packages from your account.
func (api *APIClient) GetProxyPackages() ([]UserPackage, error) {
	body, err := api.getReq(packages)
	if err != nil {
		return nil, err
	}

	var packs PackageResponse

	err = json.Unmarshal(body, &packs)
	if err != nil {
		return nil, err
	}

	if !packs.Success {
		return nil, errors.New("ERROR: " + string(body))
	}

	for _, p := range packs.PackageData {
		api.KnownPackages[p.ID] = p.Package
	}

	return packs.PackageData, nil
}

// GetAuthIPs gets all authentication IPs active on your account.
func (api *APIClient) GetAuthIPs() ([]AuthIP, error) {
	body, err := api.getReq(authips)
	if err != nil {
		return nil, err
	}

	var auths AuthIPResponse

	err = json.Unmarshal(body, &auths)
	if err != nil {
		return nil, err
	}

	if !auths.Success {
		return nil, errors.New("ERROR: " + string(body))
	}

	return auths.AuthIPData, nil
}

// DeleteAllAuthIPs deletes all authenticaiton IPs from your account.
func (api *APIClient) DeleteAllAuthIPs() (int, error) {
	aips, err := api.GetAuthIPs()
	if err != nil {
		return 0, err
	}

	todo := len(aips)
	done := 0

	for _, aip := range aips {
		target := int(aip.ID.(float64))
		fmt.Println("deleting ", target)
		if api.DeleteAuthIPByID(target) {
			done++
		}
	}
	if done < todo {
		err = errors.New("failed to delete some IPs")
	}
	return done, err
}

// AddAuthIP adds a new IP to the corresponding/provided proxy package ID.
func (api *APIClient) AddAuthIP(ip net.IP, packageID int) (AddAuthIPResponse, error) {
	failadd := AddAuthIPResponse{Success: false}
	if ip.IsPrivate() || ip.IsUnspecified() || ip.IsLoopback() {
		return failadd, errors.New("ip is private: " + ip.String())
	}

	post := map[string]string{
		"ip":             ip.String(),
		"userpackage_id": strconv.Itoa(packageID),
	}

	body, err := api.postReq(authips, post)
	if err != nil {
		return failadd, err
	}

	var addipres AddAuthIPResponse

	err = json.Unmarshal(body, &addipres)
	if err != nil {
		return failadd, err
	}

	if !addipres.Success {
		return failadd, errors.New("ERROR: " + addipres.Message)
	}

	return addipres, nil
}

// DeleteAuthIPByIP will iterate through all the authips on your account and delete one that matches the given IP.
func (api *APIClient) DeleteAuthIPByIP(ipa net.IP) (err error) {
	if ipa.IsPrivate() || ipa.IsUnspecified() || ipa.IsLoopback() {
		return errors.New("IP is invalid")
	}
	aips, err := api.GetAuthIPs()
	for _, aip := range aips {
		if net.ParseIP(aip.IP).Equal(ipa) {
			target := int(aip.ID.(float64))
			if api.DeleteAuthIPByID(target) {
				return nil
			}
		}
	}
	return errors.New("IP doesn't exist")
}

// AddCurrentIPtoAllPackages adds your current WAN IP to all packages on your account.
// It returns the amount of successful packages that it was applied to.
func (api *APIClient) AddCurrentIPtoAllPackages() (success int) {
	packs, err := api.GetProxyPackages()
	if err != nil {
		fmt.Println(err)
		return
	}

	myip := GetMyIP()
	for _, p := range packs {
		_, err := api.AddAuthIP(myip, p.ID)
		if err == nil {
			success++
		}
	}
	return
}

// DeleteAuthIPByID deletes an authentication IP with the matching ID provided
func (api *APIClient) DeleteAuthIPByID(ipID int) (ok bool) {
	body, err := api.deleteReq("authips/" + strconv.Itoa(ipID) + ".json")
	if err != nil {
		return
	}

	var delipres DelAuthIPResponse
	err = json.Unmarshal(body, &delipres)

	if err != nil {
		return
	}

	if delipres.Success {
		ok = true
	}

	return
}
