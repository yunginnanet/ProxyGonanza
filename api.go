package proxygonanza

import (
	"net/http"
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
	Debug         bool
	c             *http.Client
}

// PackageResponse represents an API response from proxybonanza.com containing proxy package information.
type PackageResponse struct {
	Success     bool          `json:"success"`
	Message     string        `json:"message"`
	PackageData []UserPackage `json:"data"`
	Pages       pagination    `json:"pagination"`
}

// AuthIPResponse represents an API response from proxybonanza.com.
type AuthIPResponse struct {
	Success    bool       `json:"success"`
	Message    string     `json:"message"`
	AuthIPData []AuthIP   `json:"data"`
	Pages      pagination `json:"pagination"`
}

// AddAuthIPResponse represents an API response from proxybonanza.com.
type AddAuthIPResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		ID int `json:"id"`
	} `json:"data"`
}

// DelAuthIPResponse represents an API response from proxybonanza.com.
type DelAuthIPResponse struct {
	Success bool `json:"success"`
}

// PackageDetails represents an API response from proxybonanza.com containing proxy package information.
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

type pagination struct {
	PageCount   int         `json:"page_count"`
	CurrentPage int         `json:"current_page"`
	HasNextPage bool        `json:"has_next_page"`
	HasPrevPage bool        `json:"has_prev_page"`
	Count       int         `json:"count"`
	Limit       interface{} `json:"limit"`
}

// UserPackage represents a proxy package purchased from proxybonanza.com.
type UserPackage struct {
	ID         int         `json:"id"`
	CustomName interface{} `json:"custom_name"`
	Login      string      `json:"login"`
	Password   string      `json:"password"`
	Expires    time.Time   `json:"expires"`
	Bandwidth  int64       `json:"bandwidth"`

	// FIXME:
	// See https://github.com/yunginnanet/ProxyGonanza/issues/1
	// LastIPChange                   time.Time      `json:"last_ip_change"`

	LowBanwidthNotificationPercent int            `json:"low_banwidth_notification_percent"`
	Package                        PackageDetails `json:"package"`
	BandwidthGb                    float64        `json:"bandwidth_gb"`
	AdditionalBandwidthGb          float64        `json:"additional_bandwidth_gb"`
	BandwidthPercentLeftHuman      string         `json:"bandwidth_percent_left_human"`
	ExpirationDateHuman            string         `json:"expiration_date_human"`
	Name                           string         `json:"name"`
}

// AuthIP is an IP address authorized to use the proxies in the related package.
type AuthIP struct {
	UserpackageID int         `json:"userpackage_id"`
	ID            interface{} `json:"id"`
	IP            string      `json:"ip"`
}

// PackageStatistics represents the statistics for the related proxy package.
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
