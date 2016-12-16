package pubip

import "time"

// Version indicates the version of this library.
const Version = "1.0.0"

// MaxTries is the maximum amount of tries to attempt to one service.
const MaxTries = 3

// APIURIs is the URIs of the services.
var APIURIs = []string{
	"https://api.ipify.org",
	"http://myexternalip.com/raw",
	"http://ipinfo.io/ip",
	"http://ipecho.net/plain",
	"http://icanhazip.com",
	"http://ifconfig.me/ip",
	"http://ident.me",
	"http://checkip.amazonaws.com",
	"http://bot.whatismyipaddress.com",
	"http://whatismyip.akamai.com",
	"http://wgetip.com",
	"http://ip.appspot.com",
	"http://ip.tyk.nu",
	"https://shtuff.it/myip/short",
}

// Timeout sets the time limit of collecting results from different services.
var Timeout = time.Second
