package pubip

import "time"

// The version of this library.
const Version = "1.0.0"

// The maximum amount of tries to attempt when making API calls.
const MaxTries = 3

// This is the ipify service base URI.  This is where all API requests go.
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

var Timeout = time.Second
