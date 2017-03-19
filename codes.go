// useful codes
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

// simple http client

func test() {
	res, _ :=
		http.Get("http://goinpracticebook.com")
	b, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("Ts", b)
}

// check weather a timeout error occours
func hasTimeOut(err error) bool {
	switch err := err.(type) {
	case *url.Error:
		// a url.Error may be caused by underlying net error that can checked for a timeout
		if err, ok := err.Err.(net.Error); ok && err.Timeout() {
			return true
		}

	//looks for timeouts detected by the net package
	case net.Error:
		if err.Timeout() {
			return true
		}

	//looks for timeouts detected by the net package
	case *net.OpError:
		if err.Timeout() {
			return true
		}
	}

	// some errors, without a custom type or variable to check against, can indicate a timeout
	errTxt := "use of closed network connection"
	if err != nil && strings.Contains(err.Error(), errTxt) {
		return true
	}
	return false
}

// hasTimeOut use case
func testHasTimeOut() {
	_, err := http.Get("http://example.com/test.zip")
	if err != nil && hasTimeOut(err) {
		fmt.Println("A timeout error occured")
		return
	}
}

//timeout and resume with httpp
func main() {
	file, err := os.Create("file.zip")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	location := "https://example.com/file.zip"
	err = download(location, file, 100)
	if err != nil {
		fmt.Println(err)
		return
	}
	fi, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Got it with %v bytes downloaded", fi.Size())

	// hash check for file integrity improve trust in final download
}

func download(location string, file *os.File, retries int64) error {
	req, err := http.NewRequest("GET", location, nil)
	if err != nil {
		return err
	}
	fi, err := file.Stat()
	if err != nil {
		return err
	}
	current := fi.Size()
	if current > 0 {
		start := strconv.FormatInt(current, 10)
		req.Header.Set("Range", "bytes"+start+"-")
	}
	cc := &http.Client{Timeout: 5 * time.Minute}
	res, err := cc.Do(req)
	if err != nil && hasTimeOut(err) {
		if retries > 0 {
			return download(location, file, retries-1)
		}
		return err
	} else if err != nil {
		return err
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		errFmt := "Unsuccess HTTP request. Status: %s"
		return fmt.Errorf(errFmt, res.Status)
	}

	if res.Header.Get("Accept-Ranges") != "bytes" {
		retries = 0
	}

	_, err = io.Copy(file, res.Body)
	if err != nil && hasTimeOut(err) {
		if retries > 0 {
			return download(location, file, retries-1)
		}
		return err
	} else if err != nil {
		return err
	}

	return nil
}

// float precision set
// truncate float64 into specified precision
func SetFloatPrecision(num float64, precision int) float64 {
	fmts := fmt.Sprintf("%%.%vf", precision)
	sn := fmt.Sprintf(fmts, num)
	n, _ := strconv.ParseFloat(sn, 64)

	return n
}

//struct field init with func expression ?
type test struct {
	x string
}

func tests(a string) string {
	return "def"
}

func tmain() {
	a := test{x: tests("abc")}
	fmt.Println(a)
}
