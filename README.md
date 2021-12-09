# proxygonanza
--
    import "."


## Usage

```go
const (
	APIBaseURL = "https://proxybonanza.com/api/v1/"
)
```

#### func  GetMyIP

```go
func GetMyIP() net.IP
```

#### type APIClient

```go
type APIClient struct {
	Key           string
	KnownPackages map[int]PackageDetails
}
```

APIClient is a client for ProxyBonanza.com.

#### func  NewApiClient

```go
func NewApiClient(key string) *APIClient
```
NewApiClient instantiates a proxybonanza.com API client with the given key.

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

#### func (*APIClient) GetAuthIPs

```go
func (api *APIClient) GetAuthIPs() ([]AuthIP, error)
```
GetAuthIPs gets all authentication IPs active on your account.

#### func (*APIClient) GetProxyPackages

```go
func (api *APIClient) GetProxyPackages() ([]UserPackage, error)
```
GetProxyPackages gets current proxy packages from your account.

#### type AddAuthIPResponse

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

#### type AuthIP

```go
type AuthIP struct {
	UserpackageID int         `json:"userpackage_id"`
	ID            interface{} `json:"id"`
	IP            string      `json:"ip"`
}
```

AuthIP is an IP address authorized to use the proxies in the related package.

#### type AuthIPResponse

```go
type AuthIPResponse struct {
	Success    bool       `json:"success"`
	Message    string     `json:"message"`
	AuthIPData []AuthIP   `json:"data"`
	Pages      pagination `json:"pagination"`
}
```

AuthIPResponse represents an API response from proxybonanza.com.

#### type DelAuthIPResponse

```go
type DelAuthIPResponse struct {
	Success bool `json:"success"`
}
```

DelAuthIPResponse represents an API response from proxybonanza.com.

#### type Package

```go
type Package struct {
	ID           int
	AuthIPs      []AuthIP
	AllTimeStats PackageStatistics
	HourlyStats  map[time.Time]PackageStatistics
}
```

Package contains what we know about a particular proxybonanza package.

#### type PackageDetails

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

#### type PackageResponse

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

#### type PackageStatistics

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

#### type UserPackage

```go
type UserPackage struct {
	ID                             int            `json:"id"`
	CustomName                     interface{}    `json:"custom_name"`
	Login                          string         `json:"login"`
	Password                       string         `json:"password"`
	Expires                        time.Time      `json:"expires"`
	Bandwidth                      int64          `json:"bandwidth"`
	LastIPChange                   time.Time      `json:"last_ip_change"`
	LowBanwidthNotificationPercent int            `json:"low_banwidth_notification_percent"`
	Package                        PackageDetails `json:"package"`
	BandwidthGb                    float64        `json:"bandwidth_gb"`
	AdditionalBandwidthGb          int            `json:"additional_bandwidth_gb"`
	BandwidthPercentLeftHuman      string         `json:"bandwidth_percent_left_human"`
	ExpirationDateHuman            string         `json:"expiration_date_human"`
	Name                           string         `json:"name"`
}
```

UserPackage represents a proxy package purchased from proxybonanza.com.
