# ProxyGonanza
[![GoDoc](https://godoc.org/git.tcp.direct/kayos/proxygonanza?status.svg)](https://godoc.org/git.tcp.direct/kayos/proxygonanza)
[![Go Report Card](https://goreportcard.com/badge/github.com/yunginnanet/proxygonanza)](https://goreportcard.com/report/github.com/yunginnanet/proxygonanza)

`import "git.tcp.direct/kayos/proxygonanza"`

## Documentation

1. [Getting Started](#getting-started)
   1. [type APIClient](#type-apiclient)
   1. [func  NewAPIClient](#func--newapiclient)
   1. [func  NewCustomClient](#func--newcustomclient)
2. [Functions](#Functions)
   1. [func (*APIClient) AddAuthIP](#func-apiclient-addauthip)
   1. [func (*APIClient) AddCurrentIPtoAllPackages](#func-apiclient-addcurrentiptoallpackages)
   1. [func (*APIClient) DeleteAllAuthIPs](#func-apiclient-deleteallauthips)
   1. [func (*APIClient) DeleteAuthIPByID](#func-apiclient-deleteauthipbyid)
   1. [func (*APIClient) DeleteAuthIPByIP](#func-apiclient-deleteauthipbyip)
   1. [func (*APIClient) GetAllSOCKSIPsAndPorts](#func-apiclient-getallsocksipsandports)
   1. [func (*APIClient) GetAuthIPs](#func-apiclient-getauthips)
   1. [func (*APIClient) GetPackageSOCKS](#func-apiclient-getpackagesocks)
   1. [func (*APIClient) GetProxyPackages](#func-apiclient-getproxypackages)
3. [Additional Details](https://godoc.org/git.tcp.direct/kayos/proxygonanza)
---

```go
const (
	APIBaseURL = "https://proxybonanza.com/api/v1/"
)
```

### Getting Started

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

---

### Functions

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
account. It returns the amount of successful packages that it was applied to.

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
