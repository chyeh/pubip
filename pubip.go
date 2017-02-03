// Package pubip gets your public IP address from several services
package pubip

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/jpillora/backoff"
)

// GetIPBy queries an API to retrieve a `net.IP` of this machine's public IP
// address.
//
// Usage:
//
//		package main
//
//		import (
//			"fmt"
//			"github.com/chyeh/pubip"
//		)
//
//		func main() {
//			ip, err := pubip.GetIPBy("https://api.ipify.org")
//			if err != nil {
//				fmt.Println("Couldn't get my IP address:", err)
//			} else {
//				fmt.Println("My IP address is:", ip)
//			}
//		}
func GetIPBy(dest string) (net.IP, error) {
	b := &backoff.Backoff{
		Jitter: true,
	}
	client := &http.Client{}

	req, err := http.NewRequest("GET", dest, nil)
	if err != nil {
		return nil, err
	}

	for tries := 0; tries < MaxTries; tries++ {
		resp, err := client.Do(req)
		if err != nil {
			d := b.Duration()
			time.Sleep(d)
			continue
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode != 200 {
			return nil, errors.New(dest + " status code " + strconv.Itoa(resp.StatusCode) + ", body: " + string(body))
		}

		tb := strings.TrimSpace(string(body))
		ip := net.ParseIP(tb)
		if ip == nil {
			return nil, errors.New("IP address not valid: " + tb)
		}
		return ip, nil
	}

	return nil, errors.New("Failed to reach " + dest)
}

// GetIPStrBy queries an API to retrieve a `string` of this machine's public IP
// address.
//
// Usage:
//
//		package main
//
//		import (
//			"fmt"
//			"github.com/chyeh/pubip"
//		)
//
//		func main() {
//			ip, err := pubip.GetIPBy("https://api.ipify.org")
//			if err != nil {
//				fmt.Println("Couldn't get my IP address:", err)
//			} else {
//				fmt.Println("My IP address is:", ip)
//			}
//		}
func GetIPStrBy(dest string) (string, error) {
	ip, err := GetIPBy(dest)
	return ip.String(), err
}

func detailErr(err error, errs []error) error {
	errStrs := []string{err.Error()}
	for _, e := range errs {
		errStrs = append(errStrs, e.Error())
	}
	j := strings.Join(errStrs, "\n")
	return errors.New(j)
}

func validate(rs []net.IP) (net.IP, error) {
	if rs == nil {
		return nil, fmt.Errorf("Failed to get any result from %d APIs", len(APIURIs))
	}
	if len(rs) < 3 {
		return nil, fmt.Errorf("Less than %d results from %d APIs", 3, len(APIURIs))
	}
	first := rs[0]
	for i := 1; i < len(rs); i++ {
		if !reflect.DeepEqual(first, rs[i]) { //first != rs[i] {
			return nil, fmt.Errorf("Results are not identical: %s", rs)
		}
	}
	return first, nil
}

func worker(d string, r chan<- net.IP, e chan<- error) {
	ip, err := GetIPBy(d)
	if err != nil {
		e <- err
		return
	}
	r <- ip
}

// Get queries several APIs to retrieve a `net.IP` of this machine's public IP
// address.
//
// Usage:
//
//		package main
//
//		import (
//			"fmt"
//			"github.com/chyeh/pubip"
//		)
//
//		func main() {
//			ip, err := pubip.Get()
//			if err != nil {
//				fmt.Println("Couldn't get my IP address:", err)
//			} else {
//				fmt.Println("My IP address is:", ip)
//			}
//		}
func Get() (net.IP, error) {
	var results []net.IP
	resultCh := make(chan net.IP, len(APIURIs))
	var errs []error
	errCh := make(chan error, len(APIURIs))

	for _, d := range APIURIs {
		go worker(d, resultCh, errCh)
	}
	for {
		select {
		case err := <-errCh:
			errs = append(errs, err)
		case r := <-resultCh:
			results = append(results, r)
		case <-time.After(Timeout):
			r, err := validate(results)
			if err != nil {
				return nil, detailErr(err, errs)
			}
			return r, nil
		}
	}
}

// GetStr queries several APIs to retrieve a `string` of this machine's public
// IP address.
//
// Usage:
//
//		package main
//
//		import (
//			"fmt"
//			"github.com/chyeh/pubip"
//		)
//
//		func main() {
//			ip, err := pubip.Get()
//			if err != nil {
//				fmt.Println("Couldn't get my IP address:", err)
//			} else {
//				fmt.Println("My IP address is:", ip)
//			}
//		}
func GetStr() (string, error) {
	ip, err := Get()
	return ip.String(), err
}
