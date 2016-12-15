// Package ipify provides a single function for retrieving your computer's
// public IP address from the ipify service: http://www.ipify.org
package pubip

import (
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jpillora/backoff"
)

// GetIp queries the ipify service (http://www.ipify.org) to retrieve this
// machine's public IP address.  Returns your public IP address as a string, and
// any error encountered.  By default, this function will run using exponential
// backoff -- if this function fails for any reason, the request will be retried
// up to 3 times.
//
// Usage:
//
//		package main
//
//		import (
//			"fmt"
//			"github.com/rdegges/go-ipify"
//		)
//
//		func main() {
//			ip, err := ipify.GetIp()
//			if err != nil {
//				fmt.Println("Couldn't get my IP address:", err)
//			} else {
//				fmt.Println("My IP address is:", ip)
//			}
//		}
func GetIp(dest string) (string, error) {
	b := &backoff.Backoff{
		Jitter: true,
	}
	client := &http.Client{}

	req, err := http.NewRequest("GET", dest, nil)
	if err != nil {
		return "", err
	}

	//req.Header.Add("User-Agent", USER_AGENT)

	for tries := 0; tries < MaxTries; tries++ {
		resp, err := client.Do(req)
		if err != nil {
			d := b.Duration()
			time.Sleep(d)
			continue
		}

		defer resp.Body.Close()

		out, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}

		if resp.StatusCode != 200 {
			return "", errors.New(dest + " status code" + strconv.Itoa(resp.StatusCode) + ", body: " + string(out))
		}

		tout := strings.TrimSpace(string(out))
		ip := net.ParseIP(tout)
		if ip == nil {
			return "", errors.New("IP address not valid: " + tout)
		}
		return ip.String(), nil
	}

	return "", errors.New("Failed to reach " + dest)
}

func isValid(rs []string) bool {
	if len(rs) < 2 {
		return false
	}
	for _, s := range rs {
		if s != rs[0] {
			return false
		}
	}
	return true
}

func worker(d string, r chan<- string, e chan<- error) {
	ip, err := GetIp(d)
	if err != nil {
		e <- err
		return
	}
	r <- ip
}

func Get() (string, error) {
	var results []string
	var errs []string
	resultCh := make(chan string, len(APIURIs))
	errCh := make(chan error)

	for _, d := range APIURIs {
		go worker(d, resultCh, errCh)
	}
	for {
		select {
		case err := <-errCh:
			errs = append(errs, err.Error())
		case r := <-resultCh:
			results = append(results, r)
			if isValid(results) {
				return results[0], nil
			}
		case <-time.After(Timeout):
			return "", errors.New(strings.Join(errs, "\n"))
		}
	}
}
