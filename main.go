package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var tr = &http.Transport{
	TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	DisableCompression: true,
}

var client = &http.Client{Transport: tr,
	Timeout: time.Duration(10 * time.Second),
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		// panic(err)
	}
}

func savePage(filename string, resp *http.Response) {
	defer resp.Body.Close()
	out, err := os.Create(fmt.Sprintf("./get/%s.html", filename))
	if err != nil {
		// panic?
	}
	defer out.Close()
	io.Copy(out, resp.Body)
}

var ascii_lowercase = "abcdefghijklmnopqrstuvwxyz"
var lowercaseArray = strings.Split(ascii_lowercase, "")

func page(urlname string) {

	for {
		url := fmt.Sprintf("http://www.lazo.twmail.org/%s.html", urlname)
		resp, err := client.Get(url)
		if err != nil {
			continue
		}
		defer resp.Body.Close()
		if resp.Body != nil {
			check(err)
			if resp.StatusCode != 404 {
				fmt.Println(urlname, resp.StatusCode)
				savePage(urlname, resp)
				break
			} else {
				break
			}
		}

	}
}

func intToString(a int) string {

	b := []int{}
	for i := 1; i < a; i = i * 26 {
		b = append([]int{i}, b...) // prepend trick
	}

	o := []byte{}

	// make the original number
	for _, b := range b {
		o = append(o, 96+byte(a/b))
		a = a % b
	}
	return string(o)

}

func start(urlname string) {
	// for i := 0; i < 101; i++ {
	fullName := fmt.Sprintf("%s", urlname)
	// fmt.Println(fullName)
	page(fullName)
	// }

}

var wg sync.WaitGroup

func main() {
	var going int = 0
	for i := 0; i < 45679; i++ {
		numToURL := ""
		if i < 24 {
			numToURL = lowercaseArray[i]
		} else {
			numToURL = intToString(i)
		}
		fmt.Println(i, numToURL)
		wg.Add(1)
		go func(numToURL string) {
			defer wg.Done()
			start(numToURL)
		}(numToURL)

		going++
		if going == 500 {
			wg.Wait()
			going = 0
		}
	}
}
func main2() {
	var going int = 0
	stringsList := []string{"a", "bx", "new", "xd", "xf", "xg", "xm", "xo", "xp", "xx"}
	for i, numToURL := range stringsList {
		fmt.Println(i, numToURL)
		wg.Add(1)
		go func(numToURL string) {
			defer wg.Done()
			start(numToURL)
		}(numToURL)
		going++
	}
	wg.Wait()
}
