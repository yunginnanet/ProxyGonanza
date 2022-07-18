# ProxyGonanza
[![GoDoc](https://godoc.org/git.tcp.direct/kayos/proxygonanza?status.svg)](https://godoc.org/git.tcp.direct/kayos/proxygonanza)
[![Go Report Card](https://goreportcard.com/badge/github.com/yunginnanet/proxygonanza)](https://goreportcard.com/report/github.com/yunginnanet/proxygonanza)

# Table of Contents

1. [ProxyGonanza](#proxygonanza)
   1. [Import Path](#import-path)
   1. [Constants and Variables](#constants-and-variables)
   1. [API Types](#api-types)
      1. [type AddAuthIPResponse](#type-addauthipresponse)
      1. [type AuthIP](#type-authip)
      1. [type AuthIPResponse](#type-authipresponse)
      1. [type DelAuthIPResponse](#type-delauthipresponse)
      1. [type Package](#type-package)
      1. [type PackageDetails](#type-packagedetails)
      1. [type PackageResponse](#type-packageresponse)
      1. [type PackageStatistics](#type-packagestatistics)
      1. [type UserPackage](#type-userpackage)
   1. [Client](#client)
      1. [type APIClient](#type-apiclient)
      1. [Constructors](#constructors)
         1. [func  NewAPIClient](#func--newapiclient)
         1. [func  NewCustomClient](#func--newcustomclient)
      1. [Methods](#methods)
         1. [func (*APIClient) AddAuthIP](#func-apiclient-addauthip)
         1. [func (*APIClient) AddCurrentIPtoAllPackages](#func-apiclient-addcurrentiptoallpackages)
         1. [func (*APIClient) DeleteAllAuthIPs](#func-apiclient-deleteallauthips)
         1. [func (*APIClient) DeleteAuthIPByID](#func-apiclient-deleteauthipbyid)
         1. [func (*APIClient) DeleteAuthIPByIP](#func-apiclient-deleteauthipbyip)
         1. [func (*APIClient) DeleteOtherAuthIPs](#func-apiclient-deleteotherauthips)
         1. [func (*APIClient) GetAllSOCKSIPsAndPorts](#func-apiclient-getallsocksipsandports)
         1. [func (*APIClient) GetAuthIPs](#func-apiclient-getauthips)
         1. [func (*APIClient) GetPackageSOCKS](#func-apiclient-getpackagesocks)
         1. [func (*APIClient) GetProxyPackages](#func-apiclient-getproxypackages)

## Import Path

`import "git.tcp.direct/kayos/proxygonanza"`

## Constants and Variables

```go
const (
	APIBaseURL = "https://proxybonanza.com/api/v1/"
)
```

## API Types

### type AddAuthIPResponse

```go
type AddAuthIPResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		ID int `json:"id"`
	} `json:"data"`
}
```

AddAuthIPResponse represents an API response from proxybonanza.com.

### type AuthIP

```go
type AuthIP struct {
	UserpackageID int         `json:"userpackage_id"`
	ID            interface{} `json:"id"`
	IP            string      `json:"ip"`
}
```

AuthIP is an IP address authorized to use the proxies in the related package.

### type AuthIPResponse

```go
type AuthIPResponse struct {
	Success    bool       `json:"success"`
	Message    string     `json:"message"`
	AuthIPData []AuthIP   `json:"data"`
	Pages      pagination `json:"pagination"`
}
```

AuthIPResponse represents an API response from proxybonanza.com.

### type DelAuthIPResponse

```go
type DelAuthIPResponse struct {
	Success bool `json:"success"`
}
```

DelAuthIPResponse represents an API response from proxybonanza.com.

### type Package

```go
type Package struct {
	ID           int
	AuthIPs      []AuthIP
	AllTimeStats PackageStatistics
	HourlyStats  map[time.Time]PackageStatistics
}
```

Package contains what we know about a particular proxybonanza package.

### type PackageDetails

```go
type PackageDetails struct {
	Name               string      `json:"name"`
	Bandwidth          int64       `json:"bandwidth"`
	Price              interface{} `json:"price"`
	HowmanyIPs         int         `json:"howmany_ips"`
	PricePerGig        interface{} `json:"price_per_gig"`
	PackageType        string      `json:"package_type"`
	HowmanyAuthips     int         `json:"howmany_authips"`
	IPType             int         `json:"ip_type"`
	PriceUserFormatted string      `json:"price_user_formatted"`
}
```

PackageDetails represents an API response from proxybonanza.com containing proxy
package information.

### type PackageResponse

```go
type PackageResponse struct {
	Success     bool          `json:"success"`
	Message     string        `json:"message"`
	PackageData []UserPackage `json:"data"`
	Pages       pagination    `json:"pagination"`
}
```

PackageResponse represents an API response from proxybonanza.com containing
proxy package information.

### type PackageStatistics

```go
type PackageStatistics struct {
	UserpackageID int    `json:"userpackage_id"`
	Date          string `json:"date"`
	BndHTTP       int    `json:"bnd_http"`
	ConnHTTP      int    `json:"conn_http"`
	BndSocks      int    `json:"bnd_socks"`
	ConnSocks     int    `json:"conn_socks"`
	BndTotal      int    `json:"bnd_total"`
	ConnTotal     int    `json:"conn_total"`
}
```

PackageStatistics represents the statistics for the related proxy package.

### type UserPackage

```go
type UserPackage struct {
	ID         int         `json:"id"`
	CustomName interface{} `json:"custom_name"`
	Login      string      `json:"login"`
	Password   string      `json:"password"`
	Expires    time.Time   `json:"expires"`
	Bandwidth  int64       `json:"bandwidth"`

	LowBanwidthNotificationPercent int            `json:"low_banwidth_notification_percent"`
	Package                        PackageDetails `json:"package"`
	BandwidthGb                    float64        `json:"bandwidth_gb"`
	AdditionalBandwidthGb          float64        `json:"additional_bandwidth_gb"`
	BandwidthPercentLeftHuman      string         `json:"bandwidth_percent_left_human"`
	ExpirationDateHuman            string         `json:"expiration_date_human"`
	Name                           string         `json:"name"`
}
```

UserPackage represents a proxy package purchased from proxybonanza.com.

## Client

### type APIClient

```go
type APIClient struct {
	Key           string
	KnownPackages map[int]PackageDetails
	Debug         bool
}
```

APIClient is a client for ProxyBonanza.com.

### Constructors

#### func  NewAPIClient

```go
func NewAPIClient(key string) *APIClient
```
NewAPIClient instantiates a proxybonanza.com API client with the given key using
golang's default http client.

#### func  NewCustomClient

```go
func NewCustomClient(key string, client *http.Client) *APIClient
```
NewCustomClient insantiates a proxybonanza API client with the given key and the
given http.Client.

### Methods

#### func (*APIClient) AddAuthIP

```go
func (api *APIClient) AddAuthIP(ip net.IP, packageID int) (AddAuthIPResponse, error)
```
AddAuthIP adds a new IP to the corresponding/provided proxy package ID.

#### func (*APIClient) AddCurrentIPtoAllPackages

```go
func (api *APIClient) AddCurrentIPtoAllPackages() (success int)
```
AddCurrentIPtoAllPackages adds your current WAN IP to all packages on your
account. It returns the amount of successful packages that it was applied to. It
will skip packages that are already using the current IP.

#### func (*APIClient) DeleteAllAuthIPs

```go
func (api *APIClient) DeleteAllAuthIPs() (int, error)
```
DeleteAllAuthIPs deletes all authenticaiton IPs from your account.

#### func (*APIClient) DeleteAuthIPByID

```go
func (api *APIClient) DeleteAuthIPByID(ipID int) (ok bool)
```
DeleteAuthIPByID deletes an authentication IP with the matching ID provided

#### func (*APIClient) DeleteAuthIPByIP

```go
func (api *APIClient) DeleteAuthIPByIP(ipa net.IP) (err error)
```
DeleteAuthIPByIP will iterate through all the authips on your account and delete
one that matches the given IP.

#### func (*APIClient) DeleteOtherAuthIPs

```go
func (api *APIClient) DeleteOtherAuthIPs() ([]int, error)
```
DeleteOtherAuthIPs deletes all authenticaiton IPs from your account's packages
that do not match your current IP address. Returns a slice of authentication IP
IDs that were deleted and any errors that occurred.

#### func (*APIClient) GetAllSOCKSIPsAndPorts

```go
func (api *APIClient) GetAllSOCKSIPsAndPorts() ([]string, error)
```
GetAllSOCKSIPsAndPorts will return a slice of IP:Port formatted proxy strings

#### func (*APIClient) GetAuthIPs

```go
func (api *APIClient) GetAuthIPs() ([]AuthIP, error)
```
GetAuthIPs gets all authentication IPs active on your account.

#### func (*APIClient) GetPackageSOCKS

```go
func (api *APIClient) GetPackageSOCKS(packageid int) ([]string, error)
```
GetPackageSOCKS returns a specified packages SOCKS5 proxies in host:port format.

#### func (*APIClient) GetProxyPackages

```go
func (api *APIClient) GetProxyPackages() ([]UserPackage, error)
```
GetProxyPackages gets current proxy packages from your account.
