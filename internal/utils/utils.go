package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
)

const urlTemplate = "https://adventofcode.com/2021/day/{}/input"

// Loads puzzle input based on day, must set environment variable AOC_SESSION_COOKIE 
func LoadInput(day string) ([]byte, error) {
	secret := os.Getenv("AOC_SESSION_COOKIE")

	if len(secret) == 0 {
		return nil, fmt.Errorf("missing AOC_SESSION_COOKIE")
	}

	var cookies []*http.Cookie

	cookie := &http.Cookie{
		Name: "session",
		Value: secret,
		Path: "/",
		Domain: ".adventofcode.com",
	}
	cookies = append(cookies, cookie)

	jar, _ := cookiejar.New(nil)
	
	inputUrl := strings.Replace(urlTemplate, "{}", "1", -1)
	cookieUrl, _ := url.Parse(inputUrl)

	jar.SetCookies(cookieUrl, cookies)

	client := &http.Client{
		Jar: jar,
	}

	req, _ := http.NewRequest("GET", inputUrl, nil)
	resp, err := client.Do(req)

	if err != nil {
        return nil, fmt.Errorf("errored when sending request to the server")
    }

    defer resp.Body.Close()
    resp_body, _ := ioutil.ReadAll(resp.Body)

    fmt.Println(resp.Status)
    fmt.Println(string(resp_body))

	return []byte(resp_body), nil
}