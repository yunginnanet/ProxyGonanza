package proxygonanza

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"
)

func (api *APIClient) debugPrintf(format string, obj ...interface{}) {
	if api.Debug {
		fmt.Println(fmt.Sprintf(format, obj...))
	}
}

// NewAPIClient instantiates a proxybonanza.com API client with the given key using golang's default http client.
func NewAPIClient(key string) *APIClient {
	return &APIClient{
		Key:           key,
		KnownPackages: make(map[int]PackageDetails),
		c:             http.DefaultClient,
	}
}

// NewCustomClient insantiates a proxybonanza API client with the given key and the given http.Client.
func NewCustomClient(key string, client *http.Client) *APIClient {
	return &APIClient{
		Key:           key,
		KnownPackages: make(map[int]PackageDetails),
		c:             client,
	}
}

// GetAllSOCKSIPsAndPorts will return a slice of IP:Port formatted proxy strings
func (api *APIClient) GetAllSOCKSIPsAndPorts() ([]string, error) {
	packs, err := api.GetProxyPackages()
	if err != nil {
		return []string{}, err
	}

	var results []string
	for _, pack := range packs {
		packsocks, err := api.GetPackageSOCKS(pack.ID)
		if err != nil {
			return results, err
		}
		results = append(results, packsocks...)
	}
	return results, nil
}

// fun fact, this wasn't in their api docs.
// proxybonanza.com/api/v1/userpackages_ippacks.csv?_delimiter=%3A&userpackage_id=xxxxx&_csvNoHeader=0&_csvFields=ip%2Cport_socks

// GetPackageSOCKS returns a specified packages SOCKS5 proxies in host:port format.
func (api *APIClient) GetPackageSOCKS(packageid int) ([]string, error) {
	body, err := api.getReq("userpackages_ippacks.csv?_delimiter=%3A&userpackage_id=" + strconv.Itoa(packageid) + "&_csvNoHeader=0&_csvFields=ip%2Cport_socks")
	if err != nil {
		return []string{}, err
	}
	var results []string
	scanner := bufio.NewScanner(bytes.NewReader(body))
	for i := 0; scanner.Scan(); i++ {
		if i == 0 {
			continue
		}
		results = append(results, scanner.Text())
	}
	return results, nil
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
		return nil, fmt.Errorf("failed to get auth IPs: %w", err)
	}

	var auths AuthIPResponse

	err = json.Unmarshal(body, &auths)
	if err != nil {
		print(string(body))
		return nil, fmt.Errorf("failed to unmarshal auth ips: %w", err)
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

// DeleteOtherAuthIPs deletes all authenticaiton IPs from your account's packages that do not match your current IP address.
// Returns a slice of authentication IP IDs that were deleted and any errors that occurred.
func (api *APIClient) DeleteOtherAuthIPs() ([]int, error) {
	myip, err := getMyIP()
	if err != nil {
		return []int{}, fmt.Errorf("failed to get my IP: %w", err)
	}
	aips, err := api.GetAuthIPs()
	if err != nil {
		return []int{}, fmt.Errorf("failed to get auth IPs: %w", err)
	}
	todo := len(aips)
	done := 0
	var deleted []int
	for _, aip := range aips {
		if net.ParseIP(aip.IP).Equal(myip) {
			todo--
			continue
		}
		if api.DeleteAuthIPByID(int(aip.ID.(float64))) {
			done++
			deleted = append(deleted, int(aip.ID.(float64)))
		}
	}
	if done < todo {
		err = errors.New("failed to delete some IPs")
	}
	return deleted, err
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
	var aips []AuthIP
	aips, err = api.GetAuthIPs()
	if err != nil {
		return
	}
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
// It will skip packages that are already using the current IP.
func (api *APIClient) AddCurrentIPtoAllPackages() (success int) {
	myip, myIPErr := getMyIP()
	if myIPErr != nil {
		// TODO: probably don't handle any of these errors like this - avoiding breaking changes for now.
		println("ProxyGonanza failed to retrieve own IP: " + myIPErr.Error())
		return
	}
	aips, aipErr := api.GetAuthIPs()
	if aipErr != nil {
		println("ProxyGonanza failed to get auth IPs: " + aipErr.Error())
		return
	}
	packs, packErr := api.GetProxyPackages()
	if packErr != nil {
		println("ProxyGonanza failed to get proxy packages: " + packErr.Error())
		return 0
	}
	var skips []int
	for _, aip := range aips {
		aipIP := net.ParseIP(aip.IP)
		if aipIP.Equal(myip) {
			skips = append(skips, aip.UserpackageID)
		}
	}
	for _, pack := range packs {
		for _, skip := range skips {
			if pack.ID == skip {
				continue
			}
		}
		_, err := api.AddAuthIP(myip, pack.ID)
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
