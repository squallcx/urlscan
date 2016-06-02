package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
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
		o = append(o, 97+byte(a/b))
		a = a % b
	}
	return string(o)

}

func start(urlname string) {
	// for i := 0; i < 101; i++ {
	fullName := fmt.Sprintf("%s%02d", urlname, 1)
	fullName2 := fmt.Sprintf("%s", urlname)
	// fmt.Println(fullName)
	page(fullName)
	page(fullName2)
	// }

}

var wg sync.WaitGroup

func writeToFile(text string) {
	fileHandle, _ := os.Create("output.txt")
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()

	fmt.Fprintln(writer, text)
	writer.Flush()
}

func readRecord() int {
	dat, err := ioutil.ReadFile("output.txt")
	check(err)
	f := string(dat)
	f = strings.Replace(f, "\n", "", 2)
	record, err := strconv.Atoi(f)
	check(err)
	return record
}
func main() {
	var going int = 0
	startnum := readRecord()
	fmt.Print(startnum)
	for i := startnum; i < 208827064576; i++ {
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
		if going == 100 {
			wg.Wait()
			iStr := strconv.Itoa(i)
			writeToFile(iStr)
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
